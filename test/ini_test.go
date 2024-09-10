package test

import (
	"errors"
	"reflect"
	"testing"
	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
	"github.com/stretchr/testify/assert"
)

const filePath = "/home/doha/doha/codescalers/week2/ini/test/testdata/data.ini"

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
	t.Run("tests should pass", func(t *testing.T) {

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
			description: "test for correct input by INI Rules",
			input: `  [ package name ]
					name = ini parser
				file path =  /pkg/ini.go

				[package version]
				version = v1.0.0

				;comment should not be part of ini map
			`,
			expect: ReturnedExpectedMap(),
			err:    nil,
		},
		{
			description: "test for input missing ] between key and value",
			input: `  [ package name 
			name = ini parser
			file path =  /pkg/ini.go

			[package version]
			version = v1.0.0

			;comment should not be part of ini map

			`,
			expect: make(map[string]map[string]string, 0),
			err:    errors.New("the file is not following ini rules"),
		},
		{
			description: "test for no key and value for section",
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
			description: "test for input empty",
			input: `
			`,
			expect: make(map[string]map[string]string, 0),
			err:    nil,
		},
		{
			description: "test for input with empty section and key = value",
			input: ` []
			key = value
			`,
			expect: make(map[string]map[string]string, 0),
			err:    errors.New("the file is not following ini rules"),
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
	t.Run("tests should pass for sending input by strings", func(t *testing.T) {

		tests := []struct {
			description string
			input       string
			expect      []string
		}{
			{
				description: "test for retrive all section names correctly",
				input: `[ package name ]
					name = ini parser
				file path =  /pkg/ini.go

				[package version]`,
				expect: []string{"package name", "package version"},
			},
			{
				description: "test for retrive empty section names",
				input: `[]
				[]`,
				expect: []string(nil),
			},
		}

		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				ini := inipkg.NewIni()
				ini.LoadFromString(test.input)

				result := ini.GetSectionNames()

				if !reflect.DeepEqual(result, test.expect) {
					t.Fail()
				}
			})
		}

	})
}

func TestGetSections(t *testing.T) {
	t.Run("tests should pass for sending input by strings", func(t *testing.T) {

		tests := []struct {
			description string
			input       string
			expect      map[string]map[string]string
		}{
			{
				description: "test for retrive ini map correctly",
				input: `[ package name ]
					name = ini parser
				file path =  /pkg/ini.go
				[package version]
				version = v1.0.0`,
				expect: ReturnedExpectedMap(),
			},
			{
				description: "test for retrive empty section",
				input: `[]
				key = value
				`,
				expect: make(map[string]map[string]string),
			},
			{
				description: "test for retrive iniMap while there is no key and value for section",
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
				ini.LoadFromString(test.input)

				result := ini.GetSections()
				assert.Equal(t, test.expect, result)
			})
		}

	})
}

func TestGet(t *testing.T) {
	t.Run("tests should pass for sending input by strings", func(t *testing.T) {
		input := `
				[ package name ]
			name = ini parser
		file path =  /pkg/ini.go

		[package version]
		version = v1.0.0

		;comment should not be part of ini map

		`

		tests := []struct {
			description string
			section     string
			key         string
			expect      string
			err         error
		}{
			{
				description: "test for retrive value correctly",
				section:     "package name",
				key:         "name",
				expect:      "ini parser",
				err:         nil,
			},
			{
				description: "test for retrive key not exist",
				section:     "package name",
				key:         "wrongname",
				expect:      "",
				err:         errors.New("there is no key with this name"),
			},
			{
				description: "test for retrive secion not exist",
				section:     "package",
				key:         "wrongname",
				expect:      "",
				err:         errors.New("there is no section with this name"),
			},
		}

		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				ini := inipkg.NewIni()
				ini.LoadFromString(input)

				value, err := ini.Get(test.section, test.key)
				assert.Equal(t, test.expect, value)
				assert.Equal(t, test.err, err)
			})
		}

	})

}

func TestSet(t *testing.T) {
	t.Run("tests should pass for sending input by strings", func(t *testing.T) {

		input := `
				[ package name ]
			name = ini parser
		file path =  /pkg/ini.go

		[package version]
		version = v1.0.0

		;comment should not be part of ini map

		`
		tests := []struct {
			description string
			section     string
			key         string
			value       string
		}{
			{
				description: "test for set new value correctly",
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
		}

		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				ini := inipkg.NewIni()
				ini.LoadFromString(input)

				ini.Set(test.section, test.key, test.value)

				expect, _ := ini.Get(test.section, test.key)
				assert.Equal(t, expect, test.value)
			})
		}

	})

}

func TestSaveToFile(t *testing.T) {
	t.Run("tests should pass for sending input by strings and saving file with correct format", func(t *testing.T) {
		input := `
				[ package name ]
			name = ini parser
		file path =  /pkg/ini.go

		[package version]
		version = v1.0.0
		comping = v2.0.0


		;comment should not be part of ini map

		`

		ini := inipkg.NewIni()
		ini.LoadFromString(input)

		path := "/home/doha/doha/codescalers/week2/ini/file.ini"

		ini.SaveToFile(path)

		savedResult := ini.GetSections()

		ini.LoadFromFile(path)
		expect := ini.IniMap

		assert.Equal(t, expect ,savedResult)

	})
}
