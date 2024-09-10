package test

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
)

type testCases struct {
	input string
	expect map[string]map[string]string
	err error
}

const filePath = "/home/doha/doha/codescalers/week2/ini/test/testdata/data.ini"


func ReturnedExpectedMap() map[string]map[string]string{
	return map[string]map[string]string {
		"package name": {
			"name" : "ini parser",
			"file path" : "/pkg/ini.go",
		},
		"package version" : {
			"version" : "v1.0.0",
		},
	}
}

func TestLoadFromFile(t *testing.T){
	t.Run("tests should pass", func(t *testing.T) {

		ini := inipkg.NewIni()
		err := ini.LoadFromFile(filePath)
		
		assert.Equal(t ,ReturnedExpectedMap() ,ini.IniMap)	
		assert.Equal(t ,nil ,err)
	})
}



func TestLoadFromString(t *testing.T){
	tests := []testCases{
		{
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
			input: `  [ package name 
			name = ini parser
			file path =  /pkg/ini.go

			[package version]
			version = 
			;comment should not be part of ini map

			`,
			expect: make(map[string]map[string]string, 0),

			err: errors.New("the file is not following ini rules"),
		},
		{
			input: `  [ package name 
			 = ini parser
			file path =  /pkg/ini.go

			 version]
			version = 
			;comment should not be part of ini map

			`,

			expect: make(map[string]map[string]string, 0),

			err: errors.New("the file is not following ini rules"),
		},

	}
	t.Run("tests should pass", func(t *testing.T) {
		for _, test := range tests {
			ini := inipkg.NewIni()
			err := ini.LoadFromString(test.input)
			
			assert.Equal(t ,test.expect ,ini.IniMap)	
			assert.Equal(t ,test.err ,err)
		}
	})
}



