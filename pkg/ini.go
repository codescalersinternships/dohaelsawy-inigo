package pkg

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	// used when the given ini file/string not following ini rules
	ErrInStructure = errors.New("the file is not following ini rules")

	// used when trying to retrive specified section that is not exist
	ErrNoSection = errors.New("there is no section with this name")

	// used when trying to retrive specified key that is not exist
	ErrNoKey = errors.New("there is no key with this name")
)

// representation of ini file structure
type IniFile struct {
	IniMap map[string]map[string]string
}



// initialize the ini structure before parsing
func NewIni() *IniFile {
	return &IniFile{IniMap: make(map[string]map[string]string)}
}




func (ini *IniFile) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	// after reading data file we pass it to LoadFromString
	// to parse the content of the file directly
	return ini.LoadFromString(string(data))
}



/* Rules of ini file:

1- assume there're no global keys, every keys need to be part of a section
2- assume the key value separator is just '='
3- keys and values should have spaces trimmed
4- comments are only valid at the beginning of the line

*/

func (ini *IniFile) LoadFromString(data string) error {

	lines := strings.Split(data, "\n")

	// to store last section found
	var lastSection string

	// to store key and value after spliting on '='
	var keyValue []string

	var key, value string

	for _, line := range lines {

		//removes any spaces, newlines, and taps from beginning and ending only
		iniLine := strings.Trim(line, " \n\t")

		// checkes for emptyline or comments
		if len(iniLine) == 0 || iniLine[0] == ';' {
			continue
		}

		// find a valid section
		if iniLine[0] == '[' && iniLine[len(iniLine)-1] == ']' {

			// removed [] from section
			section := strings.Trim(iniLine, "[ ]")

			// if the section ended up empty then return error and reset the structure
			if section == "" {
				ini.IniMap = make(map[string]map[string]string)
				return ErrInStructure
			}

			// section valid
			ini.IniMap[section] = make(map[string]string, 0)

			// sortes last section for rest its key and values
			lastSection = section
			continue
		}

		// if we got a key for not exist section 
		if _, ok := ini.IniMap[lastSection]; !ok {

			// reset the structure and return an error
			ini.IniMap = map[string]map[string]string{}
			return ErrInStructure
		}

		// valid key, seperate key and value
		keyValue = strings.Split(iniLine, "=")

		// remove any spaces, newlines, and taps
		key = strings.Trim(keyValue[0], " \n\t")
		value = strings.Trim(keyValue[1], " \n\t")

		ini.IniMap[lastSection][key] = value
	}
	return nil
}

// iterate over ini structure and retrive all sections that found
func (ini *IniFile) GetSectionNames() []string {
	var ans []string

	for section := range ini.IniMap {
		ans = append(ans, section)
	}

	return ans
}

// retrive ini structure with all its sections, keys and values
func (ini *IniFile) GetSections() map[string]map[string]string {
	return ini.IniMap
}



func (ini *IniFile) Get(section_name, key string) (string, error) {

	// check if given section_name is exist in ini structure and return error if not
	if _, ok := ini.IniMap[section_name]; !ok {
		return "", ErrNoSection
	}

	// check if given key is exist in ini structure and return error if not
	if _, ok := ini.IniMap[section_name][key]; !ok {
		return "", ErrNoKey
	}

	// return the value of given section_name and key
	return ini.IniMap[section_name][key], nil
}

/*
	SET function logic :

if the section_name is not exist in the ini strucutre then
it initialize a new section_name and assign it to given key and value
*/
func (ini *IniFile) Set(section_name, key, value string) {

	if _, ok := ini.IniMap[section_name]; !ok {
		ini.IniMap[section_name] = make(map[string]string, 0)
	}

	ini.IniMap[section_name][key] = value
}



func (ini *IniFile) SaveToFile(filename string) error {
	// creates a file with given path
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	// make the content of file using ToString() function
	initext := ini.ToString()

	// writes the content into file
	_, err = file.WriteString(initext)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

/* ToString function logic :

because the range function doesn't have a specified order for iterating over maps

i get all section names by using GetSectionNames()
and sort the section names

and iterate over the sored sections
-- NOTE: the problem is still exist because the keys is still appear in random order
but atleast i eliminate the occurence of this problem
*/

func (ini *IniFile) ToString() string {
	var iniText string

	// gets all section names
	sections := ini.GetSectionNames()

	// sort section names
	sort.Strings(sections)

	// this loop won't have a random order over section names
	for _, section := range sections {

		iniText += fmt.Sprintf("[%s]\n", section)

		// this loop will have a random order over keys
		for key, value := range ini.IniMap[section] {
			iniText += fmt.Sprintf("%s = %s\n", key, value)
		}
	}
	return iniText
}
