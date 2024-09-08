package test

import (
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
		str, err := ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")
		if err != nil {
			t.Errorf("err %e", err)
		}
		println(str)
	})
}
