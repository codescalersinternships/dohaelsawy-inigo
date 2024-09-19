package iniparser

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	// used when the given ini file/string not following ini rules
	ErrSyntax = errors.New("the file is not following ini rules")

	ErrNoGlobalKey = errors.New("global keys ara not allowed")

	// used when trying to retrive specified section that is not exist
	ErrNoSection = errors.New("there is no section with this name")

	// used when trying to retrive specified key that is not exist
	ErrNoKey = errors.New("there is no key with this name")
)

// representation of ini file structure
type Parser struct {
	iniMap map[string]map[string]string
}

// initialize the ini parser structure
func NewIni() *Parser {
	return &Parser{iniMap: make(map[string]map[string]string)}
}

// LoadFromFile loads INI files
func (ini *Parser) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return ini.LoadFromString(string(data))
}

/* Rules of ini file:

1- assume there're no global keys, every keys need to be part of a section
2- assume the key value separator is just '='
3- keys and values should have spaces trimmed
4- comments are only valid at the beginning of the line

*/

// LoadFromString parser ini text
func (ini *Parser) LoadFromString(data string) error {

	lines := strings.Split(data, "\n")

	var lastSection string

	for _, line := range lines {

		iniLine := strings.TrimSpace(line)

		if len(iniLine) == 0 || iniLine[0] == ';' || iniLine[0] == '#' {
			continue
		}

		// TODO : ASK FOR BETTER WAY

		if (!strings.HasPrefix(iniLine, "[") && strings.HasSuffix(iniLine, "]")) || (strings.HasPrefix(iniLine, "[") && !strings.HasSuffix(iniLine, "]")) {
			return ErrSyntax
		}

		if strings.HasPrefix(iniLine, "[") && strings.HasSuffix(iniLine, "]") {

			section := iniLine[1 : len(iniLine)-1]

			ini.iniMap[section] = make(map[string]string, 0)

			lastSection = section
			continue
		}

		if _, ok := ini.iniMap[lastSection]; !ok || lastSection == "" {

			ini.iniMap = map[string]map[string]string{}
			return ErrNoGlobalKey
		}

		key, value, ok := strings.Cut(iniLine, "=")

		if !ok {
			return ErrSyntax
		}

		ini.iniMap[lastSection][key] = value
	}
	return nil
}

// iterate over ini structure and retrive all sections that found
func (ini *Parser) GetSectionNames() []string {
	var sections []string

	for section := range ini.iniMap {
		sections = append(sections, section)
	}

	return sections
}

// retrive ini structure with all its sections, keys and values
func (ini *Parser) GetSections() map[string]map[string]string {
	return ini.iniMap
}

// Get gets value, true if exist
func (ini *Parser) Get(sectionName, key string) (string, bool) {

	value, exists := ini.iniMap[sectionName][key]

	return value, exists
}

// Set sets value for a key in a section
func (ini *Parser) Set(sectionName, key, value string) {

	if _, ok := ini.iniMap[sectionName]; !ok {
		ini.iniMap[sectionName] = make(map[string]string, 0)
	}

	ini.iniMap[sectionName][key] = value
}

func (ini *Parser) SaveToFile(filename string) error {

	initext := ini.String()

	err := os.WriteFile("file.ini", []byte(initext), 0644)

	if err != nil {
		return err
	}

	return nil
}

// String converts INI data to string ordered alphabetically
func (ini *Parser) String() string {
	var iniText strings.Builder

	sections := ini.GetSectionNames()

	sort.Strings(sections)

	for _, section := range sections {

		iniText.WriteString(fmt.Sprintf("[%s]\n", section))

		var keys []string

		for key := range ini.iniMap[section] {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			iniText.WriteString(fmt.Sprintf("%s=%s\n", key, ini.iniMap[section][key]))
		}
	}
	return iniText.String()
}
