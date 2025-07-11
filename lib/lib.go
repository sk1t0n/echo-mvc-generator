package lib

import (
	"os"
	"path/filepath"
	"strings"
)

func MkdirAll(path string) error {
	if strings.Contains(path, "/") || strings.Contains(path, "\\") {
		err := os.MkdirAll(filepath.Dir(path), os.FileMode(os.O_CREATE))
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetEntityName(path string) string {
	fileName := filepath.Base(path)
	entityName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	if strings.Contains(entityName, "_") {
		words := strings.Split(entityName, "_")
		for i, word := range words {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
		entityName = strings.Join(words, "")
	} else {
		entityName = strings.ToUpper(string(entityName[0])) + entityName[1:]
	}

	return entityName
}
