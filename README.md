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

  ## Installation :
 - 1. Clone project
```
https://github.com/codescalersinternships/dohaelsawy-inigo.git
```
- 2. Import package as inipkg
```
inipkg "github.com/dohaelsawy/codescalers/ini/pkg" // you can change inipkg name
```
- 3. You need to make a new ini paser
 ```
 ini := inipkg.New()
```
## Then you can do the following
- 1 - Load your data from file or string
  ```
  err := ini.LoadFromFile(filePath)
  err := ini.LoadFromString(someInput)
- 2 - Get your section name in ini file
  ```
  result := ini.GetSectionNames()
   ```
- 3 - Get your ini content
  ```
  result := ini.GetSections()
  ```
- 4 - Set new value for section or key or chage their value
  ```
  ini.Set(section, key, value)
  ```
- 5 - Get value by section name and key
  ```
  value, err := ini.Get(section, key)
  ```
- 6 - Save your new structure to a file
  ```
  err = ini.SaveToFile(path)
  ```
- 7 - Convret your ini to string
  ```
  result := ini.ToString()
  ```
  
       


