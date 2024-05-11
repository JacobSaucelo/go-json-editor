package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileType struct {
	Name        string `json:"Name"`
	UserName    string `json:"UserName"`
	Status      int    `json:"Status"`
	CreatedDate string `json:"Created Date"`
	UpdatedDate string `json:"Updated Date"`
}

type FileTypeUpdated struct {
	Name        string `json:"Name"`
	UserName    string `json:"UserName"`
	Status      int    `json:"Status"`
	CreatedDate string `json:"CreatedDate"`
	UpdatedDate string `json:"UpdatedDate"`
}

func main() {
	jsonFile, err := os.Open("Data.Applications.json")
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

	fmt.Println("data", jsonData[2])

	var updatedJsonData []FileTypeUpdated

	for i := range jsonData {
		updatedJsonData = append(updatedJsonData, FileTypeUpdated{
			Name:        jsonData[i].Name,
			UserName:    jsonData[i].UserName,
			Status:      jsonData[i].Status,
			CreatedDate: jsonData[i].CreatedDate,
			UpdatedDate: jsonData[i].UpdatedDate,
		})
	}

	updatedJSON, err := json.Marshal(updatedJsonData)
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
