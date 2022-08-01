package main

import (
	"encoding/json"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
)

var filename = "emails.json"

func saveToFile(emailToSave string) (emailExist bool, err error) {
	data, err := getDataFromJSON()
	if err != nil {
		return false, err
	}

	// Check if email exist in json database
	for _, email := range data {
		if email == emailToSave {
			return true, err
		}
	}

	// Check if email valid
	if _, err := mail.ParseAddress(emailToSave); err != nil {
		return false, err
	}

	data = append(data, emailToSave)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return false, err
	}
	return false, err
}

func getDataFromJSON() (data []string, err error) {
	err = checkFile(filename)
	if err != nil {
		return data, err
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	if len(file) == 0 {
		return []string{}, nil
	}

	err = json.Unmarshal(file, &data)
	return data, err
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

func sendEmails(price float64) (err error) {
	data, err := getDataFromJSON()
	if err != nil {
		return err
	}
	return sendCurrentPrice(price, data)
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'G', -1, 64)
}

func sendCurrentPrice(price float64, emails []string) error {
	from := "redkopavli@gmail.com"
	password := "kgshbnijjpcwmaji"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(float64ToString(price))
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, emails, message)
	if err != nil {
		return err
	}
	return nil
}
