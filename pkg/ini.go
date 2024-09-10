package pkg

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ErrInStructure = errors.New("the file is not following ini rules")
	ErrNoSection   = errors.New("there is no section with this name")
	ErrNoKey       = errors.New("there is no key with this name")
)

type IniFile struct {
	IniMap map[string]map[string]string
}

func NewIni() *IniFile {
	return &IniFile{IniMap: make(map[string]map[string]string)}
}

func (ini *IniFile) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}
	return ini.LoadFromString(string(data))
}

// RUles :

// assume there're no global keys, every keys need to be part of a section
// assume the key value separator is just =
// keys and values should have spaces trimmed
// comments are only valid at the beginning of the line

func (ini *IniFile) LoadFromString(data string) error {

	lines := strings.Split(data, "\n")

	var lastSection string
	var keyValue []string
	var key, value string

	for _, line := range lines {

		iniLine := strings.Trim(line, " \n\t")

		if len(iniLine) == 0 || iniLine[0] == ';' {
			continue
		}

		if iniLine[0] == '[' && iniLine[len(iniLine)-1] == ']' {

			section := strings.Trim(iniLine, "[ ]")


			if section == "" {
				ini.IniMap = make(map[string]map[string]string)
				return ErrInStructure
			}


			ini.IniMap[section] = make(map[string]string, 0)
			lastSection = section
			continue
		}

		if _, ok := ini.IniMap[lastSection]; !ok {

			ini.IniMap = map[string]map[string]string{}
			return ErrInStructure
		}

		keyValue = strings.Split(iniLine, "=")

		key = strings.Trim(keyValue[0], " \n\t")
		value = strings.Trim(keyValue[1], " \n\t")

		ini.IniMap[lastSection][key] = value

	}
	return nil
}

func (ini *IniFile) GetSectionNames() []string {
	var ans []string
	for section := range ini.IniMap {
		ans = append(ans, section)
	}
	return ans
}

func (ini *IniFile) GetSections() map[string]map[string]string {
	return ini.IniMap
}

func (ini *IniFile) Get(section_name, key string) (string, error) {
	if _, ok := ini.IniMap[section_name]; !ok {
		return "", ErrNoSection
	}
	if _, ok := ini.IniMap[section_name][key]; !ok {
		return "", ErrNoKey
	}
	return ini.IniMap[section_name][key], nil
}

func (ini *IniFile) Set(section_name, key, value string){
	if _, ok := ini.IniMap[section_name]; !ok {
		ini.IniMap[section_name] = make(map[string]string, 0)
	} 
	ini.IniMap[section_name][key] = value
}

func (ini *IniFile) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	initext := ini.ToString()
	_, err =  file.WriteString(initext)

	if err != nil {
		return err
	} 
	defer file.Close()
	return nil
}

func (ini *IniFile) ToString() string {
	var iniText string

	sections := ini.GetSectionNames()
	sort.Strings(sections)

	for _, section := range sections {
		iniText += fmt.Sprintf("[%s]\n", section)
		for key, value := range ini.IniMap[section] {
			iniText += fmt.Sprintf("%s = %s\n", key, value)
		}
	}
	return iniText
}
