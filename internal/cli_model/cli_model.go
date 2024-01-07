package cli_model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codersgyan/expressify/internal/coding_styles"
	"github.com/codersgyan/expressify/internal/configs"
	"github.com/codersgyan/expressify/internal/databases"
	"github.com/codersgyan/expressify/internal/languages"
	"github.com/codersgyan/expressify/internal/loggers"
	"github.com/codersgyan/expressify/internal/orms"
	"github.com/codersgyan/expressify/internal/package_managers"
	"github.com/codersgyan/expressify/internal/selector"
	"github.com/codersgyan/expressify/internal/structure"
	"github.com/codersgyan/expressify/internal/test_frameworks"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

type AppState int

const (
	StateWelcome AppState = iota
	StateProjectName
	StateLanguage
	StatePackageManager
	StateTestFramework
	StateLoggerLibrary
	StateDatabase
	StateORM
	StateConfig
	StateCodingStyle
	StateFolderStructure
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
	TestFrameworkList      list.Model
	SelectedTestFramework  string
	LoggerLibraryList      list.Model
	SelectedLoggerLibrary  string
	DatabaseList           list.Model
	SelectedDatabase       string
	ORMList                list.Model
	SelectedORM            string
	ConfigList             list.Model
	SelectedConfig         string
	CodingStyleList        list.Model
	SelectedCodingStyle    string
	FolderStructureCreated bool
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
				m.CurrentState = StateTestFramework
				return m, nil
			}

			if m.CurrentState == StateTestFramework {
				i, ok := m.TestFrameworkList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedTestFramework = string(i)
				}
				m.CurrentState = StateLoggerLibrary
				return m, nil
			}

			if m.CurrentState == StateLoggerLibrary {
				i, ok := m.LoggerLibraryList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedLoggerLibrary = string(i)
				}
				m.CurrentState = StateDatabase
				return m, nil
			}

			if m.CurrentState == StateDatabase {
				i, ok := m.DatabaseList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedDatabase = string(i)
				}
				m.CurrentState = StateORM
				return m, nil
			}

			if m.CurrentState == StateORM {
				i, ok := m.ORMList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedORM = string(i)
				}
				m.CurrentState = StateConfig
				return m, nil
			}

			if m.CurrentState == StateConfig {
				i, ok := m.ConfigList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedConfig = string(i)
				}
				m.CurrentState = StateCodingStyle
				return m, nil
			}

			if m.CurrentState == StateCodingStyle {
				i, ok := m.CodingStyleList.SelectedItem().(selector.Item)
				if ok {
					m.SelectedCodingStyle = string(i)
				}
				m.CurrentState = StateFolderStructure
				return m, nil
			}

			if m.CurrentState == StateFolderStructure {
				// Create folder structure
				err := structure.CreateBaseFileStructure(m.ProjectNameInput.Value(), m.SelectedLanguage)
				if err != nil {
					fmt.Printf("error creating folder structure: %v", err)
					return m, tea.Quit
				}
				// todo: transition to next state
				// m.CurrentState = StateFolderStructure
				return m, nil
			}

		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	case errMsg:
		m.Error = msg
		return m, tea.Quit
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

	if m.CurrentState == StateTestFramework {
		m.TestFrameworkList, cmd = m.TestFrameworkList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateLoggerLibrary {
		m.LoggerLibraryList, cmd = m.LoggerLibraryList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateDatabase {
		m.DatabaseList, cmd = m.DatabaseList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateORM {
		m.ORMList, cmd = m.ORMList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateConfig {
		m.ConfigList, cmd = m.ConfigList.Update(msg)
		return m, cmd
	}

	if m.CurrentState == StateCodingStyle {
		m.CodingStyleList, cmd = m.CodingStyleList.Update(msg)
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
		s = "\n📗 Enter your project name:\n\n" + m.ProjectNameInput.View()
	case StateLanguage:
		if m.SelectedLanguage != "" {
			var str string
			if m.SelectedLanguage == string(languages.JavaScript) {
				str = "🎉 Awesome choice! JavaScript brings flexibility and dynamism to your project. Let's get coding! 🚀"
			} else if m.SelectedLanguage == string(languages.TypeScript) {
				str = "👍 Great pick! TypeScript adds type safety and robustness to your application. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.LanguageList.View()
	case StatePackageManager:
		if m.SelectedPackageManager != "" {
			var str string
			if m.SelectedPackageManager == string(package_managers.NPM) {
				str = "🎉 Awesome choice! NPM is the world's most popular package manager. Let's get coding! 🚀"
			} else if m.SelectedPackageManager == string(package_managers.PNPM) {
				str = "👍 Great pick! PNPM is a fast, disk space efficient package manager. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.PackageManagerList.View()
	case StateTestFramework:
		if m.SelectedTestFramework != "" {
			var str string
			if m.SelectedTestFramework == string(test_frameworks.SuperTestWithJest) {
				str = "🎉 Awesome choice! Supertest is best framework 🚀"
			} else if m.SelectedTestFramework == string(test_frameworks.MochaWithChaiHTTP) {
				str = "👍 Great pick! Mocha with Chai is powerful framework"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.TestFrameworkList.View()

	case StateLoggerLibrary:
		if m.SelectedLoggerLibrary != "" {
			var str string
			if m.SelectedLoggerLibrary == string(loggers.Winston) {
				str = "🎉 Awesome choice! Winston is best logger out there 🚀"
			} else if m.SelectedLoggerLibrary == string(loggers.Bunyan) {
				str = "👍 Great pick! Bunyan is powerful logger"
			} else if m.SelectedLoggerLibrary == string(loggers.Pino) {
				str = "👍 Great pick! Pino is powerful logger"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.LoggerLibraryList.View()

	case StateDatabase:
		if m.SelectedDatabase != "" {
			var str string
			if m.SelectedDatabase == string(databases.MongoDB) {
				str = "🎉 Awesome choice! MongoDB is powerful NoSQL database 🚀"
			} else if m.SelectedDatabase == string(databases.PostgreSQL) {
				str = "👍 Great pick! PostgreSQL is powerful relational database"
			} else if m.SelectedDatabase == string(databases.MySQL) {
				str = "👍 Great pick! MySQL is powerful database"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.DatabaseList.View()

	case StateORM:
		if m.SelectedORM != "" {
			var str string
			if m.SelectedORM == string(orms.Mongoose) {
				str = "🎉 Awesome choice! 🚀"
			} else if m.SelectedORM == string(orms.Sequelize) {
				str = "👍 Great pick!"
			} else if m.SelectedORM == string(orms.TypeORM) {
				str = "👍 Great pick!"
			} else if m.SelectedORM == string(orms.Prisma) {
				str = "👍 Great pick!"
			} else if m.SelectedORM == string(orms.None) {
				str = "🔥 Loving hardcore, yeah 💪"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.ORMList.View()

	case StateConfig:
		if m.SelectedConfig != "" {
			var str string
			if m.SelectedConfig == string(configs.ENV) {
				str = "🎉 Awesome choice! 🚀"
			} else if m.SelectedConfig == string(configs.JSON) {
				str = "👍 Great pick!"
			} else if m.SelectedConfig == string(configs.YAML) {
				str = "👍 Great pick!"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.ConfigList.View()

	case StateCodingStyle:
		if m.SelectedCodingStyle != "" {
			var str string
			if m.SelectedCodingStyle == string(coding_styles.Functional) {
				str = "🎉 Awesome choice! 🚀"
			} else if m.SelectedCodingStyle == string(coding_styles.ObjectOriented) {
				str = "👍 Great pick!"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.CodingStyleList.View()
	}
	return s
}

func InitialModel() CliModel {
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "my-expressify-app"
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
		TestFrameworkList:  test_frameworks.NewTestFrameworkSelector().List,
		LoggerLibraryList:  loggers.NewLoggerSelector().List,
		DatabaseList:       databases.NewDatabaseSelector().List,
		ORMList:            orms.NewORMSelector().List,
		ConfigList:         configs.NewConfigSelector().List,
		CodingStyleList:    coding_styles.NewCodingStyleSelector().List,
	}
}
