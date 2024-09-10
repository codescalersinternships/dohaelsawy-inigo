package test

import (
	"errors"
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
		input  string
		expect map[string]map[string]string
		err    error
	}{
		{
			description: "correct input by INI Rules",

			input: `  [ package name ]
					name = ini parser
				file path =  /pkg/ini.go

				[package version]
				version = v1.0.0

				;comment should not be part of ini map
			`,
			expect: ReturnedExpectedMap(),

			err: nil,
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

			err: errors.New("the file is not following ini rules"),
		},
		{
			description: "input missing the key and its value ",

			input: `  [ package name ]
			name = ini parser
			file path =  /pkg/ini.go

			[package version]
			 = 
			;comment should not be part of ini map

			`,
			expect: make(map[string]map[string]string, 0),

			err: errors.New("the file is not following ini rules"),
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
	t.Run("tests should pass for sending input by file", func(t *testing.T) {
		expect := []string{"package name", "package version"}

		ini := inipkg.NewIni()
		ini.LoadFromFile(filePath)

		result := ini.GetSectionNames()
		t.Logf("%v",ini.IniMap)
		assert.Equal(t, expect, result)
	})
}



func TestGetSections(t *testing.T){
	ini := inipkg.NewIni()
	ini.LoadFromFile(filePath)

	result := ini.GetSections()
	t.Errorf("%v   %v",result,ReturnedExpectedMap())
}


func TestGet(t *testing.T){
	ini := inipkg.NewIni()
	ini.LoadFromFile(filePath)

	result , _ := ini.Get("package name","name")
	t.Errorf("%v   ",result)
}
