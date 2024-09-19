package iniparser

import (
	"reflect"
	"sort"
	"testing"
)

const filePath = "./testdata/data.ini"

const generalInput = `  
[package name]
name=ini parser
file path=/pkg/ini.go = hello

[package version]
version=v1.0.0

;comment should not be part of ini map
`

func ReturnedExpectedMap() map[string]map[string]string {

	return map[string]map[string]string{
		"package name": {
			"file path": "/pkg/ini.go = hello",
			"name":      "ini parser",
		},
		"package version": {
			"version": "v1.0.0",
		},
	}
}

func assertIsEqual(t *testing.T, type1, type2 any) {
	t.Helper() 

	if !reflect.DeepEqual(type1, type2) {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error is %v", err)
	}
}

func TestLoadFromFile(t *testing.T) {
	t.Run("parse input by file", func(t *testing.T) {

		ini := NewIni()
		err := ini.LoadFromFile(filePath)

		assertIsEqual(t, ReturnedExpectedMap(), ini.iniMap)
		assertIsEqual(t, nil, err)
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
			file path =  /pkg/ini.go = hello
			[package version]
			version = v1.0.0
			;comment should not be part of ini map
			`,
			expect: make(map[string]map[string]string, 0),
			err:    ErrSyntax,
		},

		{
			description: "section with no key and value",
			input: `[package name]
					name=ini parser
					file path=/pkg/ini.go = hello

					[package version]
				`,
			expect: map[string]map[string]string{
				"package name": {
					"name":      "ini parser",
					"file path": "/pkg/ini.go = hello",
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
			err:    ErrNoGlobalKey,
		},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {

			ini := NewIni()
			err := ini.LoadFromString(test.input)

			assertIsEqual(t, test.expect, ini.iniMap)
			assertIsEqual(t, test.err, err)
		})
	}
}

func TestGetSectionNames(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expect      []string
		err         error
	}{
		{
			description: "return all section names correctly",
			input:       generalInput,
			expect:      []string{"package name", "package version"},
			err:         nil,
		},
		{
			description: "return empty section names",
			input: `[]
			[]`,
			expect: []string{""},
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			ini := NewIni()
			err := ini.LoadFromString(test.input)

			result := ini.GetSectionNames()

			sort.Strings(result)
			sort.Strings(test.expect)

			assertIsEqual(t, test.expect, result)
			assertNoError(t, err)
		})
	}
}

func TestGetSections(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expect      map[string]map[string]string
		err         error
	}{
		{
			description: "return ini map correctly",
			input:       generalInput,
			expect:      ReturnedExpectedMap(),
			err:         nil,
		},
		{
			description: "empty section with key and value",
			input: `[]
			key = value
			`,
			expect: make(map[string]map[string]string),
			err:    ErrNoGlobalKey,
		},
		{
			description: "section with no key and value",
			input: `[package name]
				name=ini parser
			file path=/pkg/ini.go = hello

			[package version]
			`,
			expect: map[string]map[string]string{
				"package name": {
					"name":      "ini parser",
					"file path": "/pkg/ini.go = hello",
				},
				"package version": make(map[string]string),
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ini := NewIni()
			err := ini.LoadFromString(test.input)

			result := ini.GetSections()
			assertIsEqual(t, test.expect, result)
			assertIsEqual(t, test.err, err)
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
		ok          bool
	}{
		{
			description: "return value correctly",
			section:     "package name",
			key:         "name",
			expect:      "ini parser",
			ok:          true,
		},
		{
			description: "return key not exist",
			section:     "package name",
			key:         "wrongname",
			expect:      "",
			ok:          false,
		},
		{
			description: "return secion not exist",
			section:     "package",
			key:         "wrongname",
			expect:      "",
			ok:          false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ini := NewIni()
			err := ini.LoadFromString(input)

			value, ok := ini.Get(test.section, test.key)
			assertIsEqual(t, test.expect, value)
			assertIsEqual(t, test.ok, ok)
			assertNoError(t, err)
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
			ini := NewIni()
			err := ini.LoadFromString(input)

			ini.Set(test.section, test.key, test.value)

			expect, _ := ini.Get(test.section, test.key)
			assertIsEqual(t, expect, test.value)
			assertNoError(t, err)
		})
	}

}

func TestSaveToFile(t *testing.T) {
	t.Run("parse input by strings and saving file with valid format", func(t *testing.T) {

		ini := NewIni()
		err := ini.LoadFromString(generalInput)

		assertNoError(t, err)

		path := "./file.ini"

		err = ini.SaveToFile(path)
		assertNoError(t, err)

		savedResult := ini.GetSections()

		errfile := ini.LoadFromFile(path)
		assertNoError(t, errfile)

		expect := ini.iniMap

		assertIsEqual(t, expect, savedResult)

	})
}

func TestString(t *testing.T) {
	t.Run("print ini map", func(t *testing.T) {

		ini := NewIni()

		errfile := ini.LoadFromFile(filePath)
		assertNoError(t, errfile)

		result := ini.String()

		err := ini.LoadFromString(result)
		assertNoError(t, err)

		expect := ini.String()

		assertIsEqual(t,expect,result)
	})
}
