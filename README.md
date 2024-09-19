# INI package 
### Description
ini is a simple Go package for manipulating ini files.
### Features :
- Load from multiple data sources(file, string)
- Set new sections, keys with their values
- Get value by section name and key
- Get all section names in ini file
- Retrive a your ini file in structure way
- Save your edits to ini file to new or exist file
- convet your ini file to single string

## Installation 
 - 1. Get package
```
   go get https://github.com/codescalersinternships/dohaelsawy-inigo
```
- 2. Import package as inipkg
```golang
  import inipkg "github.com/dohaelsawy/codescalers/ini/pkg" // you can change inipkg name
```
- 3. You need to make a new ini paser
 ```golang
   ini := inipkg.New()
```
## Then you can do the following
- 1 - Load your data from file or string
  ```golang
  err := ini.LoadFromFile(filePath)
  err := ini.LoadFromString(someInput)
- 2 - Get your section name in ini file
  ```golang
  result := ini.GetSectionNames()
   ```
- 3 - Get your ini content
  ```golang
  result := ini.GetSections()
  ```
- 4 - Set new value for section or key or chage their value
  ```golang
  ini.Set(section, key, value)
  ```
- 5 - Get value by section name and key
  ```golang
  value, err := ini.Get(section, key)
  ```
- 6 - Save your new structure to a file
  ```golang
  err = ini.SaveToFile(path)
  ```
- 7 - Convret your ini to string
  ```golang
  result := ini.ToString()
  ```
  
       


