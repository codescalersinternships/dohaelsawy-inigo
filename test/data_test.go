package test

import (
	"reflect"
	"testing"

	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
)

// var ini = inipkg.IniFile{
// 	IniMap: map[string]map[string]string{
// 		"example": {
// 			"key": " value",
// 		},
// 	},
// }

func TestLoadFile(t *testing.T) {

	t.Run("this test should pass", func(t *testing.T) {
		ini := inipkg.NewIni()
		// ans := &map[string]map[string]string{
		// 	"eample": {
		// 		"key":  "value",
		// 		"key2": "value",
		// 	},
		// }

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
}
