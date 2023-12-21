package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

type AppState int

const (
	StateWelcome AppState = iota
	StateProjectName
	StateLanguage
	// StateConfig
	// StateLogger
	// StateAuth
	// StateMainWork
	// StateDone
)

type Project struct {
	Name     string
	Language string
}

type UIElems struct {
	ProjectNameInput textinput.Model
}

type PredefinedMsgs struct {
	WelcomeMessage string
}

type model struct {
	CurrentState AppState
	Project
	UIElems
	PredefinedMsgs
}

func initialModel() model {
	// project name input
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "What would you like to name your project?"
	// projectNameInput.Focus()

	return model{
		Project: Project{},
		UIElems: UIElems{
			ProjectNameInput: projectNameInput,
		},
		PredefinedMsgs: PredefinedMsgs{
			WelcomeMessage: "Hey artisan! Welcome to expresify. \nLet's build together an awesome project!",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.CurrentState == StateWelcome {
				m.CurrentState = StateProjectName  // Transition to project name input
				m.UIElems.ProjectNameInput.Focus() // Focus on the input
				return m, nil
			} else if m.CurrentState == StateProjectName {
				m.Project.Name = m.UIElems.ProjectNameInput.Value() // Store the project name
				// Transition to next state (e.g., StateLanguage)
				// m.CurrentState = StateLanguage
				return m, nil
			}
		}
	}

	if m.CurrentState == StateProjectName {
		m.UIElems.ProjectNameInput, cmd = m.UIElems.ProjectNameInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var s string

	switch m.CurrentState {
	case StateWelcome:
		s = m.PredefinedMsgs.WelcomeMessage + "\nPress Enter to continue..."
	case StateProjectName:
		s = "Enter your project name:\n" + m.UIElems.ProjectNameInput.View()
	}
	return s
}

// func handleProjectName(q string) tea.Cmd {
// 	return func() tea.Msg {
// 		return QueryResponse{Result: q}
// 	}
// }

// type QueryResponse struct {
// 	Result string
// }
