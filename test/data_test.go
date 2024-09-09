package test

import (
	"reflect"
	"testing"

	inipkg "github.com/dohaelsawy/codescalers/ini/pkg"
)

func TestLoadFile(t *testing.T) {

	t.Run("this test should pass on maps equivalentes", func(t *testing.T) {
		ini := inipkg.NewIni()

		expect := map[string]map[string]string{
			"example": {
				"key":  "value",
				"key2": "value",
			},
			"example2": {
				"key":  "value",
				"key2": "value",
			},
		}
		ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")

		if !reflect.DeepEqual(ini.IniMap, expect) {
			t.Errorf("i expected %v , found %v", expect, ini.IniMap)
		}
	})

	t.Run("this test should pass on not returning errors", func(t *testing.T) {
		ini := inipkg.NewIni()

		err := ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")

		if err != nil {
			t.Errorf("error is : %e", err)
		}
	})
}

func TestLoadFromString(t *testing.T) {
	input := `[      example ]
key = value
key2 = value

[example2]
key = value
key2 = value`

	t.Run("this test should pass on maps equivalentes", func(t *testing.T) {

		ini := inipkg.NewIni()

		expect := map[string]map[string]string{
			"example": {
				"key":  "value",
				"key2": "value",
			},
			"example2": {
				"key":  "value",
				"key2": "value",
			},
		}

		ini.LoadFromString(input)

		if !reflect.DeepEqual(ini.IniMap, expect) {
			t.Errorf("i expected %v , found %v", expect, ini.IniMap)
		}
	})

	t.Run("this test should pass on not returning errors", func(t *testing.T) {
		ini := inipkg.NewIni()

		err := ini.LoadFromString(input)

		if err != nil {
			t.Errorf("error is : %s", err)
		}
	})
}

func TestGetSectionNames(t *testing.T) {
	t.Run("return all section names inside map", func(t *testing.T) {

		ini := inipkg.NewIni()

		expect := []string{
			"example",
			"example2",
		}

		ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")
		ans := ini.GetSectionNames()

		if !reflect.DeepEqual(expect, ans) {
			t.Errorf("i expected %v , found %v", expect, ans)
		}
	})
}

func TestGetSections(t *testing.T) {

	t.Run("should return all the map", func(t *testing.T) {
		ini := inipkg.NewIni()

		expect := map[string]map[string]string{
			"example": {
				"key":  "value",
				"key2": "value",
			},
			"example2": {
				"key":  "value",
				"key2": "value",
			},
		}

		ini.LoadFromFile("/home/doha/doha/codescalers/week2/ini/sample.ini")
		ans := ini.GetSections()

		if !reflect.DeepEqual(expect, ans) {
			t.Errorf("i expected %v , found %v", expect, ans)
		}
	})

}
