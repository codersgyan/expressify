package cli_model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codersgyan/expressify/internal/languages"
	"github.com/codersgyan/expressify/internal/package_managers"
	"github.com/codersgyan/expressify/internal/selector"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

type AppState int

const (
	StateWelcome AppState = iota
	StateProjectName
	StateLanguage
	StatePackageManager
	// StateConfig
	// StateLogger
	// StateAuth
	// StateMainWork
	// StateDone
)

type CliModel struct {
	CurrentState           AppState
	ProjectNameInput       textinput.Model
	WelcomeMessage         string
	LanguageList           list.Model
	SelectedLanguage       string
	PackageManagerList     list.Model
	SelectedPackageManager string
	Error                  error
}

func (m CliModel) Init() tea.Cmd {
	return nil
}

type (
	errMsg error
)

func (m CliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.LanguageList.SetWidth(msg.Width)
		m.ProjectNameInput.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// here, only changing the state
			if m.CurrentState == StateWelcome {
				m.CurrentState = StateProjectName
				m.ProjectNameInput.Focus()
				return m, nil
			}
			if m.CurrentState == StateProjectName {
				m.CurrentState = StateLanguage
				return m, nil
			}

			if m.CurrentState == StateLanguage {
				i, ok := m.LanguageList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedLanguage = string(i)
				}
				m.CurrentState = StatePackageManager
				return m, nil
			}

			if m.CurrentState == StatePackageManager {
				i, ok := m.PackageManagerList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedPackageManager = string(i)
				}
				// todo: transition next state.
				// m.CurrentState = StatePackageManager
				return m, nil
			}

		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	if m.CurrentState == StateProjectName {
		m.ProjectNameInput, cmd = m.ProjectNameInput.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateLanguage {
		m.LanguageList, cmd = m.LanguageList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StatePackageManager {
		m.PackageManagerList, cmd = m.PackageManagerList.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m CliModel) View() string {
	var s string

	switch m.CurrentState {
	case StateWelcome:
		s = m.WelcomeMessage
	case StateProjectName:
		s = "Enter your project name:\n" + m.ProjectNameInput.View()
	case StateLanguage:
		if m.SelectedLanguage != "" {
			var str string
			if m.SelectedLanguage == "JavaScript" {
				str = "🎉 Awesome choice! JavaScript brings flexibility and dynamism to your project. Let's get coding! 🚀"
			} else if m.SelectedLanguage == "TypeScript" {
				str = "👍 Great pick! TypeScript adds type safety and robustness to your application. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.LanguageList.View()
	case StatePackageManager:
		if m.SelectedPackageManager != "" {
			var str string
			if m.SelectedPackageManager == "NPM" {
				str = "🎉 Awesome choice! NPM is the world's most popular package manager. Let's get coding! 🚀"
			} else if m.SelectedPackageManager == "PNPM" {
				str = "👍 Great pick! PNPM is a fast, disk space efficient package manager. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.PackageManagerList.View()
	}
	return s
}

func InitialModel() CliModel {
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "What would you like to name your project?"
	return CliModel{
		CurrentState:     StateWelcome,
		ProjectNameInput: projectNameInput,
		WelcomeMessage: `
🌟🚀 Welcome to Expressify! 🚀🌟
We're thrilled to have you on board for a seamless and efficient Express.js application setup.
Get ready to supercharge your development process with our intuitive CLI tool.

Let's create something amazing together! 🎉👨‍💻👩‍💻

Press Enter to begin... (or ESC to quit)
`,
		LanguageList:       languages.NewLanguageSelector().List,
		PackageManagerList: package_managers.NewPManagerSelector().List,
	}
}
