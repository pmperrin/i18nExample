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
	Path     string
	Filename string
	langMap  map[string]LocalFile
}

type LocalFile struct {
	contante map[string]string
}

func InitI18n(path string, filename string) (SimpleI18n, error) {
	i18n := SimpleI18n{
		Path:     path,
		Filename: filename,
	}
	langList, err := i18n.findLanguageFile()
	if err != nil {
		return i18n, err
	}
	if len(langList) == 0 {
		return i18n, errors.New("No language file found")
	}

	mu.Lock()
	langMap, mapErr := i18n.loadLangMap(langList)
	mu.Unlock()
	i18n.langMap = langMap
	fmt.Println(i18n.langMap)
	return i18n, mapErr
}

// GetLangInstance returns the LocalFile instance for the specified language.
//
// Parameters:
//   - lang (string): The language code for which the LocalFile instance is required.
//
// Returns:
//   - (LocalFile): The LocalFile instance for the specified language.
func (i SimpleI18n) GetLang(lang string) LocalFile {
	fmt.Println(i.langMap)
	return i.langMap[lang]
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

func (i SimpleI18n) loadLangMap(langList []string) (map[string]LocalFile, error) {
	tmpMap := make(map[string]LocalFile)
	for _, lang := range langList {
		fileContante, err := i.loadLangFile(lang)
		if err != nil {
			return nil, err
		}

		tmpMap[lang] = LocalFile{
			contante: fileContante,
		}
		fmt.Println(lang + ": " + tmpMap[lang].contante["welcome.message"])
	}
	fmt.Println(tmpMap)
	return tmpMap, nil
}

func (i SimpleI18n) loadLangFile(lang string) (map[string]string, error) {
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
