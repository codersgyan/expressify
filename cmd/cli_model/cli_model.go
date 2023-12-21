package cli_model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codersgyan/expressify/cmd/languages"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

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

type CliModel struct {
	CurrentState AppState
	Project
	UIElems
	PredefinedMsgs
	List   list.Model
	Choice string
}

func (m CliModel) Init() tea.Cmd {
	return nil
}

func (m CliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
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
				m.CurrentState = StateLanguage                      // Transition to next state (e.g., StateLanguage)
				return m, nil
			} else if m.CurrentState == StateLanguage {
				i, ok := m.List.SelectedItem().(languages.Item)
				if ok {
					m.Choice = string(i)
				}
			} else {
				return m, tea.Quit
			}
		}

		if m.CurrentState == StateLanguage {
			m.List, cmd = m.List.Update(msg)
			return m, cmd
		}
	}

	if m.CurrentState == StateProjectName {
		m.UIElems.ProjectNameInput, cmd = m.UIElems.ProjectNameInput.Update(msg)
	}

	return m, cmd
}

func (m CliModel) View() string {
	var s string

	switch m.CurrentState {
	case StateWelcome:
		s = m.PredefinedMsgs.WelcomeMessage + "\nPress Enter to continue..."
	case StateProjectName:
		s = "Enter your project name:\n" + m.UIElems.ProjectNameInput.View()
	case StateLanguage:
		if m.Choice != "" {
			return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.Choice))
		}
		return "\n" + m.List.View()
	}
	return s
}

func InitialModel() CliModel {
	l := languages.NewLanguageSelector()
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "What would you like to name your project?"
	// projectNameInput.Focus()
	return CliModel{
		Project: Project{},
		UIElems: UIElems{
			ProjectNameInput: projectNameInput,
		},
		PredefinedMsgs: PredefinedMsgs{
			WelcomeMessage: "Hey artisan! Welcome to expressify. \nLet's build together an awesome project!",
		},
		List: l.List,
	}
}
