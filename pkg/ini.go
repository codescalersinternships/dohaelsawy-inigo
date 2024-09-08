package pkg

import (
	"errors"
	"os"
)

var (
	ErrFileNotExist = errors.New("File does not exist")
)

type iniFile struct {
	initmap map[string]map[string]string
}

func (ini iniFile) LoadFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", ErrFileNotExist
	}

	return ini.LoadFromString(string(data)), nil
}

func (ini iniFile) LoadFromString(data string) string {

}
