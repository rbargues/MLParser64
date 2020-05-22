package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"time"
)
// func readJSON() map[string]string {
// 	jsonFile, err := os.Open("levels.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer jsonFile.Close()
// 	byteVal, _ := ioutil.ReadAll(jsonFile)

// 	var result map[string]string
// 	json.Unmarshal([]byte(byteVal), &result)

// 	return result
// }

func getTime() map[string]map[string]string {
	jsonFile, err := os.Open("times.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteVal, _ := ioutil.ReadAll(jsonFile)

	var result map[string]map[string]string
	json.Unmarshal([]byte(byteVal), &result)

	return result
}
// func storeJSON(jsonVal map[string]string) {
// 	jsonString, _ := json.Marshal(jsonVal)
// 	jsonFile, err := os.Create("levels.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer jsonFile.Close()
// 	jsonFile.Write(jsonString)
// }
func storeTimeJson(timeJSON map[string]map[string]string) {
	jsonString, _ := json.Marshal(timeJSON)
	jsonFile, err := os.Create("times.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonString)
}

func main() {
	times := getTime()
	bobGold := times["bob-omb"]["gold-split"]

	parsed, _ := time.Parse(time.RFC3339Nano, bobGold)
	fmt.Printf("%v, %T\n", time.Since(parsed), time.Since(parsed))
	fmt.Printf("%v, %T\n", bobGold, bobGold)

	times["bob-omb"]["gold-split"] = time.Now().Format(time.RFC3339Nano)
	storeTimeJson(times)
}