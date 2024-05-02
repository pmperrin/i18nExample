package simpleI18n

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	propertiesHandler map[string]langFile
	mu                sync.RWMutex
)

type langFile struct {
	properties map[string]string
}

type SimpleI18n struct {
	Path     string
	Filename string
}

func (i SimpleI18n) Init() error {

	langList, err := i.findLanguageFile()
	if err != nil {
		return err
	}
	if len(langList) == 0 {
		return errors.New("No language file found")
	}

	mu.Lock()

	mu.Unlock()
}

// findLanguageFile reads the directory at the specified path and finds the
// corresponding language file for the given filename.
//
// It returns a slice of strings containing the names of the found language files.
// If an error occurs while reading the directory, it prints an error message and
// returns nil.
//
// Parameters:
//   - i (SimpleI18n): The instance of the SimpleI18n struct containing the path and filename.
//
// Returns:
//   - []string: A slice of strings containing the lang of the found properties files.
//   - error: An error if an error occurs while reading the directory.
func (i SimpleI18n) findLanguageFile() ([]string, error) {
	files, err := os.ReadDir(i.Path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, err
	}

	langList := []string{}
	for _, file := range files {
		fmt.Println(file.Name())

		if strings.Contains(file.Name(), ".properties") && strings.Contains(file.Name(), "_") {
			strArray := strings.Split(file.Name(), "_")
			if strArray[0] == i.Filename {
				langName := strings.Split(strArray[1], ".properties")[0]
				langList = append(langList, langName)
			}
		}
	}
	return langList, nil
}
