package lib

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	FormatEntityNamePascalCase = iota
	FormatEntityNameLowerCase
)

func MkdirAll(path string) error {
	if strings.Contains(path, "/") || strings.Contains(path, "\\") {
		err := os.MkdirAll(filepath.Dir(path), 0777)
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

// Returns the entity name in pascal case or lower case.
// Takes a file path in snake case.
// Convert snake case to pascal case or lower case.
// Convert lower case to pascal case.
func GetEntityName(path string, format int) string {
	fileName := filepath.Base(path)
	entityName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	entityName = strings.Replace(entityName, "_controller", "", 1)

	switch format {
	case FormatEntityNamePascalCase:
		if strings.Contains(entityName, "_") { // snake case to pascal case
			words := strings.Split(entityName, "_")
			for i, word := range words {
				words[i] = strings.ToUpper(string(word[0])) + word[1:]
			}
			entityName = strings.Join(words, "")
		} else { // lower case to pascal case
			entityName = strings.ToUpper(string(entityName[0])) + entityName[1:]
		}
	case FormatEntityNameLowerCase:
		if strings.Contains(entityName, "_") { // snake case to lower case
			entityName = strings.Replace(entityName, "_", "", -1)
		}
	}

	return entityName
}

func RemoveFilesAlongWithDir(dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return
	}

	for _, file := range files {
		_ = os.Remove(file)
	}

	_ = os.Remove(dir)
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func RemoveLastSlash(s string) string {
	if strings.Contains(s, "/") || strings.Contains(s, "\\") {
		s = strings.TrimSuffix(s, "/")
		s = strings.TrimSuffix(s, "\\")
	}

	return s
}
