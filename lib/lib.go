package lib

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	FormatEntityNamePascalCase = iota
	FormatEntityNameSnakeCase
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

func GetEntityName(path string, format int) string {
	fileName := filepath.Base(path)
	entityName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	switch format {
	case FormatEntityNamePascalCase:
		if strings.Contains(entityName, "_") {
			words := strings.Split(entityName, "_")
			for i, word := range words {
				words[i] = strings.ToUpper(string(word[0])) + word[1:]
			}
			entityName = strings.Join(words, "")
		} else {
			entityName = strings.ToUpper(string(entityName[0])) + entityName[1:]
		}
		entityName = strings.Replace(entityName, "Controller", "", 1)
	case FormatEntityNameSnakeCase:
		if strings.Contains(entityName, "_") {
			entityName = strings.Replace(entityName, "_controller", "", 1)
		} else {
			entityName = strings.Replace(entityName, "Controller", "", 1)

			var buf bytes.Buffer
			for i, letter := range entityName {
				if i == 0 && unicode.IsUpper(letter) {
					buf.WriteRune(unicode.ToLower(letter))
				} else if unicode.IsUpper(letter) {
					buf.WriteRune('_')
					buf.WriteRune(unicode.ToLower(letter))
				} else {
					buf.WriteRune(letter)
				}
			}
			entityName = buf.String()
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
