package pkg

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrInStructure  = errors.New("the file is not following ini rules")
)

type IniFile struct {
	IniMap map[string]map[string]string
}

func NewIni() *IniFile {
	return &IniFile{IniMap: make(map[string]map[string]string)}
}

var emptyINitFIle = &IniFile{}

func (ini *IniFile) LoadFromFile(path string) (*IniFile, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return emptyINitFIle, err
	}
	return ini.LoadFromString(string(data))
}

// RUles :

// assume there're no global keys, every keys need to be part of a section
// assume the key value separator is just =
// keys and values should have spaces trimmed
// comments are only valid at the beginning of the line

func (ini *IniFile) LoadFromString(data string) (*IniFile, error) {

	lines := strings.Split(data, "\n")

	var lastSection string
	var keyValue []string

	for _, line := range lines {
	
		iniLine := strings.ReplaceAll(line, " ", "")
		iniLine = strings.Trim(iniLine,"\n")
	
		if len(iniLine) == 0 || iniLine[0] == ';' {
			continue
		}
	
		if iniLine[0] == '[' && iniLine[len(iniLine)-1] == ']' {

			section := strings.Trim(iniLine,"[]")
			ini.IniMap[section] = make(map[string]string)
			lastSection = section
			continue
		}

		if _, ok := ini.IniMap[lastSection]; !ok {
			return emptyINitFIle , ErrInStructure
		}

		keyValue = strings.Split(iniLine,"=")
		
		if len(keyValue) != 2 {
			return emptyINitFIle , ErrInStructure
		}

		ini.IniMap[lastSection][keyValue[0]] = keyValue[1]
	}
	return ini, nil
}



func (ini *IniFile) GetSectionNames(mapIni *IniFile) []string{
	var ans []string 
	for section := range mapIni.IniMap {	
		ans = append(ans, section)
	} 
	return ans
}


func (ini *IniFile) GetSections(mapIni *IniFile) map[string]map[string]string {
	return mapIni.IniMap
}