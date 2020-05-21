package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"fmt"
	"strings"
)

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
		files = append(files, path)
		}
		return nil
	})
	return files, err
}
func storeJSON(jsonVal map[string]string) {
	jsonString, _ := json.Marshal(jsonVal)
	jsonFile, err := os.Create("levels.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonString)
}
func main() {
	files, _ := FilePathWalkDir("./")
	fmt.Printf("%v\n", files)

	jsonFile := make(map[string]string)
	for x := range files {
		filename := files[x]
		if filepath.Ext(filename) != ".png" {
			continue
		}
		fmt.Printf("%v\n", strings.Replace(filename, ".png", "", -1))
		dct := obtainDCT(filename)
		phashVal := phash(dct)
		jsonFile[strings.Replace(filename,".png","", -1)] = phashVal
		// fmt.Printf("%v", phashVal)
	}
	fmt.Printf("%v", jsonFile)
	storeJSON(jsonFile)
}