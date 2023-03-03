package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type model struct {
	textInput    textinput.Model
	showResponse bool
	content      string
	err          error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Your question"
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

			fmt.Println()
			fmt.Println("Sending request to GPT3, wait a moment ...")
			content, err := chatgpt.GetResponse(m.textInput.Value())
			if err != nil {
				panic(err)
			}

			return model{
					showResponse: true,
					textInput:    m.textInput,
					content:      content,
				},
				tea.ClearScreen

		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return model{
				textInput:    textinput.New(),
				showResponse: false,
				err:          nil,
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
	if m.showResponse {
		// remove  \n from the start of m.content

		in := `## Response 	
## ` + m.content +

			`
## Press Ctrl+C to quit or Esc to back to input`
		str, err := getGlamourResponse(in)
		if err != nil {
			panic(err)
		}
		return str
	}

	return fmt.Sprintf(
		"Ask to ChatGPT?\n\n%s\n\n%s",
		m.textInput.View(),
		"(press enter to see send question, ctrl/c to quit)",
	) + "\n"
}

func getGlamourResponse(content string) (string, error) {
	const width = 100
	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)

	if err != nil {
		return "", err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return "", err
	}

	return str, nil
}
