package clirunner

import (
	"fmt"
	"os"
)

var EmptyJson = fmt.Errorf("Empty Json")

type ValidateCmd struct {
	Path string `arg:"" name:"path" help:"Path to validate" type:"path"`
}

func (v *ValidateCmd) Run() error {
	fmt.Println(v.Path)
	content, err := v.GetContentsIfPathExists()
	if err != nil {
		return err
	}

	return v.Validate(content)
}

func (v ValidateCmd) GetContentsIfPathExists() (string, error) {
	path := v.Path
	if !fileExists(path) {
		return "", fmt.Errorf("File doesnt exists")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (v ValidateCmd) Validate(content string) error {
	if len(content) == 0 {
		return EmptyJson
	}

	// TODO: add implementation
	return nil
}
