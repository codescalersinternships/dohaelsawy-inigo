package test

import (
	"errors"
	"reflect"
	"sort"
	"testing"

	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
	"github.com/stretchr/testify/assert"
)

const filePath = "./testdata/data.ini"

const generalInput = `  
	[ package name ]
	name = ini parser
	   file path =  /pkg/ini.go

	[package version]
	version = v1.0.0

	;comment should not be part of ini map
`

var (
	ErrInStructure = errors.New("the file is not following ini rules")
	ErrNoSection   = errors.New("there is no section with this name")
	ErrNoKey       = errors.New("there is no key with this name")
)

func ReturnedExpectedMap() map[string]map[string]string {

	return map[string]map[string]string{
		"package name": {
			"file path": "/pkg/ini.go",
			"name":      "ini parser",
		},
		"package version": {
			"version": "v1.0.0",
		},
	}
}

func TestLoadFromFile(t *testing.T) {
	t.Run("parse input by file", func(t *testing.T) {

		ini := inipkg.NewIni()
		err := ini.LoadFromFile(filePath)

		assert.Equal(t, ReturnedExpectedMap(), ini.IniMap)
		assert.Equal(t, nil, err)
	})
}

func TestLoadFromString(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expect      map[string]map[string]string
		err         error
	}{
		{
			description: "input by INI Rules",
			input:       generalInput,
			expect:      ReturnedExpectedMap(),
			err:         nil,
		},

		{
			description: "input missing ] between key and value",
			input: `  [ package name 
			name = ini parser
			file path =  /pkg/ini.go
			[package version]
			version = v1.0.0
			;comment should not be part of ini map
			`,
			expect: make(map[string]map[string]string, 0),
			err:    ErrInStructure,
		},

		{
			description: "section with no key and value",
			input: `[ package name ]
					name = ini parser
				file path =  /pkg/ini.go

				[package version]
				`,
			expect: map[string]map[string]string{
				"package name": {
					"name":      "ini parser",
					"file path": "/pkg/ini.go",
				},
				"package version": make(map[string]string),
			},

			err: nil,
		},

		{
			description: "input empty",
			input:       ``,
			expect:      make(map[string]map[string]string, 0),
			err:         nil,
		},

		{
			description: "empty section and key = value",
			input: ` []
			key = value
			`,
			expect: make(map[string]map[string]string, 0),
			err:    ErrInStructure,
		},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {

			ini := inipkg.NewIni()
			err := ini.LoadFromString(test.input)

			assert.Equal(t, test.expect, ini.IniMap)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestGetSectionNames(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expect      []string
	}{
		{
			description: "return all section names correctly",
			input:       generalInput,
			expect:      []string{"package name", "package version"},
		},
		{
			description: "return empty section names",
			input: `[]
			[]`,
			expect: []string(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			ini := inipkg.NewIni()
			_ = ini.LoadFromString(test.input)

			result := ini.GetSectionNames()

			sort.Strings(result)
			sort.Strings(test.expect)

			assert.Equal(t, test.expect, result)
		})
	}
}

func TestGetSections(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expect      map[string]map[string]string
	}{
		{
			description: "return ini map correctly",
			input:       generalInput,
			expect:      ReturnedExpectedMap(),
		},
		{
			description: "empty section with key and value",
			input: `[]
			key = value
			`,
			expect: make(map[string]map[string]string),
		},
		{
			description: "section with no key and value",
			input: `[ package name ]
				name = ini parser
			file path =  /pkg/ini.go

			[package version]
			`,
			expect: map[string]map[string]string{
				"package name": {
					"name":      "ini parser",
					"file path": "/pkg/ini.go",
				},
				"package version": make(map[string]string),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ini := inipkg.NewIni()
			_ = ini.LoadFromString(test.input)

			result := ini.GetSections()
			assert.Equal(t, test.expect, result)
		})
	}

}

func TestGet(t *testing.T) {
	input := generalInput

	tests := []struct {
		description string
		section     string
		key         string
		expect      string
		err         error
	}{
		{
			description: "return value correctly",
			section:     "package name",
			key:         "name",
			expect:      "ini parser",
			err:         nil,
		},
		{
			description: "return key not exist",
			section:     "package name",
			key:         "wrongname",
			expect:      "",
			err:         ErrNoKey,
		},
		{
			description: "return secion not exist",
			section:     "package",
			key:         "wrongname",
			expect:      "",
			err:         ErrNoSection,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ini := inipkg.NewIni()
			_ = ini.LoadFromString(input)

			value, err := ini.Get(test.section, test.key)
			assert.Equal(t, test.expect, value)
			assert.Equal(t, test.err, err)
		})
	}

}

func TestSet(t *testing.T) {
	input := generalInput

	tests := []struct {
		description string
		section     string
		key         string
		value       string
	}{
		{
			description: "new value correctly",
			section:     "package name",
			key:         "name",
			value:       "parser",
		},
		{
			description: "test for set new key and value correctly",
			section:     "package version",
			key:         "new version",
			value:       "v2.0.0",
		},
		{
			description: "test for set new section and value correctly",
			section:     "package",
			key:         "new version",
			value:       "v3.0.0",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ini := inipkg.NewIni()
			_ = ini.LoadFromString(input)

			ini.Set(test.section, test.key, test.value)

			expect, _ := ini.Get(test.section, test.key)
			assert.Equal(t, expect, test.value)
		})
	}

}

func TestSaveToFile(t *testing.T) {
	t.Run("parse input by strings and saving file with valid format", func(t *testing.T) {
		input := generalInput

		ini := inipkg.NewIni()
		_ = ini.LoadFromString(input)

		path := "./file.ini"

		_ = ini.SaveToFile(path)

		savedResult := ini.GetSections()

		_ = ini.LoadFromFile(path)
		expect := ini.IniMap

		assert.Equal(t, expect, savedResult)

	})
}

func TestToString(t *testing.T) {
	t.Run("converting iniMap to string", func(t *testing.T) {

		ini := inipkg.NewIni()

		_ = ini.LoadFromFile(filePath)
		result := ini.ToString()

		_ = ini.LoadFromString(result)
		expect := ini.ToString()

		if !reflect.DeepEqual(expect, result) {
			t.Errorf("i expect %v , i got %v", expect, result)
		}
	})
}
