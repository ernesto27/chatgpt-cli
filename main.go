package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

var chatgpt *chatGPT

func main() {
	chatgpt = New()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
