package tests

import (
	"os"
	"path/filepath"
)

func CreateJsonFile(fileName string, content string) (string, error) {
	filePath, err := getBasePathFromFileName(fileName)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return "", err
	}

	return filePath, nil
}

func RemoveJsonFile(fileName string) error {
	filePath, err := getBasePathFromFileName(fileName)
	if err != nil {
		return err
	}

	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func getBasePathFromFileName(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(cwd, ".tmp", fileName), nil
}

func createTmpDirIfNotExists() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	tmpDirPath := filepath.Join(cwd, ".tmp")
	if _, err := os.Stat(tmpDirPath); os.IsNotExist(err) {
		if err := os.Mkdir(tmpDirPath, 0755); err != nil {
			panic(err)
		}
	}
}
