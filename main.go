package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileType struct {
	Name        string `json:Name`
	UserName    string `json:UserName`
	Status      int    `json:Status`
	CreatedDate string `json:Created Date`
	UpdatedDate string `json:Updated Date`
}

func main() {
	jsonFile, err := os.Open("file.json")
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonData []FileType

	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return
	}

	for i := range jsonData {
		jsonData[i].CreatedDate = jsonData[i].CreatedDate
		jsonData[i].UpdatedDate = jsonData[i].UpdatedDate
	}

	updatedJSON, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile("updated.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing json:", err)
		return
	}

	fmt.Println("JSON file modified successfully")
}
