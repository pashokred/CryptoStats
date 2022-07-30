package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MailStruct struct {
	Emails string `json:"emails"`
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveToFile(dataToSave string) error {
	filename := "emails.json"
	err := checkFile(filename)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data := []MailStruct{}

	// Make loop and check if exist element

	json.Unmarshal(file, &data)

	newStruct := &MailStruct{
		Emails: dataToSave,
	}

	data = append(data, *newStruct)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
