package pkg

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrFileNotExist = errors.New("file does not exist")
	ErrInStructure  = errors.New("the file is not following ini rules")
	succesLoading   = "loaded successfully"
)

type IniFile struct {
	IniMap map[string]map[string]string
}

func NewIni() *IniFile {
	return &IniFile{IniMap: make(map[string]map[string]string)}
}

func (ini IniFile) LoadFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", ErrFileNotExist
	}
	err = ini.LoadFromString(string(data))
	return succesLoading, err
}

// RUles :

// assume there're no global keys, every keys need to be part of a section
// assume the key value separator is just =
// keys and values should have spaces trimmed
// comments are only valid at the beginning of the line

func (ini IniFile) LoadFromString(data string) error {
	lines := strings.Split(data, "\n")
	var lastSection string = ErrInStructure.Error()
	for _, line := range lines {
		parts := strings.Fields(line)

		if len(parts) == 0 || parts[0] == ";" {
			continue
		}

		if parts[0] == "[" && parts[len(parts)-1] == "]" {
			ini.IniMap[parts[1]] = map[string]string{
				"key": "value",
			}
			lastSection = parts[1]
		} else if _, ok := ini.IniMap[lastSection]; !ok {
			return ErrInStructure
		} else {
			ini.IniMap[lastSection][parts[0]] = parts[len(parts)-1]
		}
	}
	return nil
}
