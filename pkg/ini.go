package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInStructure = errors.New("the file is not following ini rules")
	ErrNoSection = errors.New("there is no section with this name")
	ErrNoKey = errors.New("there is no key with this name")
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

	lines := strings.Split(data, "\n\t")

	var lastSection string
	var keyValue []string

	for _, line := range lines {

		iniLine := strings.Trim(line, " \n\t")

		if len(iniLine) == 0 || iniLine[0] == ';' {
			continue
		}

		if iniLine[0] == '[' && iniLine[len(iniLine)-1] == ']' {

			section := strings.Trim(iniLine, "[ ]")
			ini.IniMap[section] = make(map[string]string,0)
			lastSection = section
			continue
		}

		if _, ok := ini.IniMap[lastSection]; !ok {
			return ErrInStructure
		}

		keyValue = strings.Split(iniLine, "=")
		
		if len(keyValue) != 2 {
			return ErrInStructure
		}

		for i , kv := range keyValue{
			keyValue[i] = strings.Trim(kv," \n")
		}

		ini.IniMap[lastSection][keyValue[0]] = keyValue[1]
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


func (ini *IniFile) Get(section_name, key string) (string, error){
	if _, ok := ini.IniMap[section_name]; !ok {
		return "" , ErrNoSection
	}
	if _, ok := ini.IniMap[section_name][key]; !ok {
		return "" , ErrNoKey
	}
	return ini.IniMap[section_name][key] , nil
}



func (ini *IniFile) Set(section_name, key, value string) {
	ini.IniMap[section_name][key] = value
}

func (ini *IniFile) SaveToFile() error {
	file , err := os.Create("/home/doha/doha/codescalers/week2/ini/file.ini")
	if err != nil {
		return err
	}
	for section , keys := range ini.IniMap {
		file.WriteString("[")
		file.WriteString(section)
		file.WriteString("]")
		file.WriteString("\n")
		for key,value := range keys {
			file.WriteString(key)
			file.WriteString(" = ")
			file.WriteString(value)
			file.WriteString("\n")
		}
	}

	defer file.Close()
	return nil
}


func (ini *IniFile) ToString() string{
	var iniText string
	for section , keys := range ini.IniMap {
		iniText += fmt.Sprintf("[%s]\n",section)
		for key,value := range keys {
			iniText += fmt.Sprintf("%s = %s\n" ,key, value)
		}
	}
	return iniText
}
