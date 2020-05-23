package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

func readJSON(filename string) map[string]string {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteVal, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteVal), &result)

	return result
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

