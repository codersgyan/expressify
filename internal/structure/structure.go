package structure

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codersgyan/expressify/internal/languages"
)

func CreateBaseFileStructure(projectName string, language string) error {
	cwd, err := os.Getwd()
	fmt.Println("Current Working Directory:", cwd)

	if err != nil {
		return fmt.Errorf("unable to get current working directory: %w", err)
	}

	projectPath := filepath.Join(cwd, ".expressify", projectName)
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		return fmt.Errorf("folder \"%s\" already exists or cannot check existence", projectPath)
	}
	mkdirProjectDirErr := os.Mkdir(projectPath, 0755) // 0755 is commonly used permissions
	if mkdirProjectDirErr != nil {
		return fmt.Errorf("unable to create project folder: %w", err)
	}

	var languageWisePath string
	if language == string(languages.JavaScript) {
		languageWisePath = "jsbase"
	} else {
		languageWisePath = "tsbase"
	}

	srcPath := filepath.Join(cwd, ".templates", languageWisePath)
	dstPath := filepath.Join(projectPath)

	cpErr := CopyDir(srcPath, dstPath)
	if cpErr != nil {
		fmt.Printf("Error copying directory: %s\n", cpErr)
	} else {
		fmt.Println("Directory copied successfully.")
	}
	return nil
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// copyDir recursively copies a directory from src to dst.
func CopyDir(src, dst string) error {
	// Get properties of source dir
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create the destination directory
	err = os.MkdirAll(dst, info.Mode())
	if err != nil {
		return err
	}

	// Read directory contents
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursive call for directories
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// Copy files
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
