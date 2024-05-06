package simpleI18n

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	mu sync.RWMutex
)

const (
	VarStartDelimiter = "{{"
	VarEndDelimiter   = "}}"
)

type SimpleI18n struct {
	Path        string
	Filename    string
	DefaultLang string
	langMap     map[string]LocalFile
}

type LocalFile struct {
	contante map[string]string
}

func InitI18n(path string, filename string, defaultLang string) (SimpleI18n, error) {
	fmt.Println("Loading")
	i18n := SimpleI18n{
		Path:        path,
		Filename:    filename,
		DefaultLang: defaultLang,
	}
	langList, err := i18n.findLanguageFile()
	if err != nil {
		return i18n, err
	}
	if len(langList) == 0 {
		return i18n, errors.New("no language file found")
	}

	mu.Lock()
	langMap, mapErr := i18n.loadLangMap(langList)
	mu.Unlock()
	i18n.langMap = langMap
	return i18n, mapErr
}

func (i SimpleI18n) GetLang(lang string) (LocalFile, error) {
	if i.langMap == nil {
		return LocalFile{}, errors.New("no language file found during initialization")
	}
	if _, ok := i.langMap[lang]; !ok {
		if _, ook := i.langMap[i.DefaultLang]; !ook {
			return LocalFile{}, errors.New("no default language file found")
		}
		return i.langMap[i.DefaultLang], nil
	}
	return i.langMap[lang], nil
}

// GetText returns the value of the specified key from the localized content.
//
// Parameters:
//   - key (string): The key for which the corresponding value is required.
//
// Returns:
//   - (string): The value associated with the specified key in the localized content.
func (t LocalFile) GetText(key string) string {
	return t.contante[key]
}

func (t LocalFile) GetTextWithParam(key string, param map[string]string) string {
	result := t.contante[key]
	for key, value := range param {
		result = strings.ReplaceAll(result, VarStartDelimiter+key+VarEndDelimiter, value)
	}
	return result
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
		return nil, err
	}

	langList := []string{}
	for _, file := range files {

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

// loadLangMap loads the language files and returns a map of LocalFile instances.
//
// It takes a slice of strings containing the names of the found language files.
// It returns a map of LocalFile instances, where the keys are the language codes and the values are the corresponding LocalFile instances.
//
// If an error occurs while reading the directory or loading the language file, it returns nil and the error.
//
// Parameters:
//   - langList ([]string): A slice of strings containing the names of the found language files.
//
// Returns:
//   - (map[string]LocalFile, error): A map of LocalFile instances and an error if any.
func (i SimpleI18n) loadLangMap(langList []string) (map[string]LocalFile, error) {
	tmpMap := make(map[string]LocalFile)
	for _, lang := range langList {
		fileContante, err := i.loadPropertiesContante(lang)
		if err != nil {
			return nil, err
		}

		tmpMap[lang] = LocalFile{
			contante: fileContante,
		}
	}
	return tmpMap, nil
}

// loadPropertiesContante reads a language file and returns a map of key-value pairs.
//
// It takes a language code as a parameter and returns a map of strings where the keys are the keys from the properties file and the values are the corresponding values.
//
// The function opens the specified language file, reads its content line by line, and splits each line into a key-value pair. It then populates a map with these key-value pairs and returns it.
//
// If an error occurs while opening the file or reading its content, the function returns nil and the error.
//
// Parameters:
//   - lang (string): The language code for which the language file is to be read.
//
// Returns:
//   - (map[string]string, error): A map of strings containing the key-value pairs from the language file, and an error if any.
func (i SimpleI18n) loadPropertiesContante(lang string) (map[string]string, error) {
	props := make(map[string]string)
	filePath := i.Path + "/" + i.Filename + "_" + lang + ".properties"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			props[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return props, nil
}
