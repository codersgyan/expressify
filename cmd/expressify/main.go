package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codersgyan/expressify/internal/cli_model"
)

func main() {

	// cwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("unable to get current working directory: %w", err)
	// 	os.Exit(1)
	// }
	// srcPath := cwd + "/.templates/jsbase"
	// dstPath := cwd + "/.expressify/auth-service"

	// cpErr := structure.CopyDir(srcPath, dstPath)
	// if cpErr != nil {
	// 	fmt.Printf("Error copying directory: %s\n", cpErr)
	// } else {
	// 	fmt.Println("Directory copied successfully.")
	// }
	// return n

	p := tea.NewProgram(cli_model.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
