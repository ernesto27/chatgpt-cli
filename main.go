package main

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	glamour   bool
	content   string
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 60

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c := gogpt.NewClient("")

			resp, err := c.CreateChatCompletion(
				context.Background(),
				gogpt.ChatCompletionRequest{
					Model: gogpt.GPT3Dot5Turbo,
					Messages: []gogpt.ChatCompletionMessage{
						{
							Role:    "user",
							Content: m.textInput.Value(),
						},
					},
				},
			)

			if err != nil {
				panic(err)
			}
			return model{
					glamour:   true,
					textInput: m.textInput,
					content:   resp.Choices[0].Message.Content,
				},
				tea.ClearScreen
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return model{
				textInput: textinput.New(),
				glamour:   false,
				err:       nil,
			}, tea.ClearScreen
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput.Focus()
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.glamour {

		in := `## Response 
		
## ` + m.content +
			`
## Press Ctrl+C or Esc to exit`

		const width = 100

		vp := viewport.New(width, 100)
		vp.Style = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			PaddingRight(2)

		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width),
		)

		if err != nil {
			panic(err)
		}

		str, err := renderer.Render(in)
		if err != nil {
			panic(err)
		}

		return str
	}

	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		m.textInput.View(),
		"(press enter to see glamour model, esc to quit)",
	) + "\n"
}
