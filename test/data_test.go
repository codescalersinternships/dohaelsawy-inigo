package test

import (
	"reflect"
	"testing"

	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
)

func TestLoadFile(t *testing.T) {

	t.Run("this test should pass on maps equivalentes", func(t *testing.T) {
		ini := inipkg.NewIni()


		res := &inipkg.IniFile{
			IniMap: map[string]map[string]string{
				"example": {
					"key":  "value",
					"key2": "value",
				},
			},
		}


		iniresult, _ := ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")
		
		if !reflect.DeepEqual(iniresult, res) {
			t.Errorf("i expected %v , found %v", res, iniresult)
		}
	})

	t.Run("this test should pass on not returning errors", func(t *testing.T) {
		ini := inipkg.NewIni()

		_, err := ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")
		
		if err != nil {
			t.Errorf("error is : %e" , err)
		}
	})
}



func TestLoadFromString(t *testing.T) {
	input := `[      example ]
    key = value
    key2 = value`

	t.Run("this test should pass on maps equivalentes", func(t *testing.T) {

		ini := inipkg.NewIni()


		res := &inipkg.IniFile{
			IniMap: map[string]map[string]string{
				"example": {
					"key":  "value",
					"key2": "value",
				},
			},
		}


		iniresult, _ := ini.LoadFromString(input)
		
		if !reflect.DeepEqual(iniresult, res) {
			t.Errorf("i expected %v , found %v", res, iniresult)
		}
	})

	t.Run("this test should pass on not returning errors", func(t *testing.T) {
		ini := inipkg.NewIni()

		_, err := ini.LoadFromString(input)
		
		if err != nil {
			t.Errorf("error is : %s" , err)
		}
	})
}
