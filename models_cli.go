package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type myViewport struct {
	viewport viewport.Model
}

func newViewport(in string) (*myViewport, error) {
	const width = 120

	vp := viewport.New(width, 40)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	in += "\n## Press Ctrl+C to quit or Esc to back to input"
	str, err := renderer.Render(in)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &myViewport{
		viewport: vp,
	}, nil
}

func (vp myViewport) Init() tea.Cmd {
	return nil
}

func (vp myViewport) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return vp, nil
}

func (vp myViewport) View() string {
	return vp.viewport.View()
}

type model struct {
	textInput    textarea.Model
	showResponse bool
	content      string
	err          error
	viewport     viewport.Model
}

func initialModel() model {
	return model{
		textInput:    newTextarea(),
		err:          nil,
		showResponse: false,
	}
}

func newTextarea() textarea.Model {
	ti := textarea.New()
	ti.Placeholder = "Your question"
	ti.ShowLineNumbers = false
	ti.SetHeight(10)
	ti.SetWidth(50)
	ti.Focus()

	return ti
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tea.ClearScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlP:
			fmt.Println()
			fmt.Println("Sending request to GPT3, wait a moment ...")
			content, err := chatgpt.GetResponse(m.textInput.Value())
			if err != nil {
				panic(err)
			}

			modelViewport, err := newViewport(content)
			if err != nil {
				return m, tea.Quit
			}
			return model{
					showResponse: true,
					textInput:    m.textInput,
					viewport:     modelViewport.viewport,
				},
				tea.ClearScreen
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return model{
				textInput:    newTextarea(),
				showResponse: false,
				err:          nil,
			}, tea.ClearScreen
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput.Focus()
	m.textInput, cmd = m.textInput.Update(msg)
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.showResponse {
		return m.viewport.View()
	}

	tea.ClearScreen()
	return fmt.Sprintf(
		"Ask to ChatGPT\n\n%s\n\n%s",
		m.textInput.View(),
		"(press Ctrl+P to send question, Ctrl+C to quit)",
	) + "\n"
}
