package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	gMarket string = "btcuah"
)

type emailJSON struct {
	emails []string
}

func main() {

	http.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		price, err := GetLatestPrice(gMarket)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		resp := make(map[string]string)
		priceStr := strconv.FormatFloat(price, 'E', -1, 64)
		resp["mktprice"] = priceStr
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		saveToFile(priceStr)
		w.Write(jsonResp)
	})

	http.HandleFunc("/subscribe", func(rw http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	})

	http.HandleFunc("/sendEmails", func(rw http.ResponseWriter, r *http.Request) {
		// Get data from json, send to all emails using net/http
	})

	http.ListenAndServe(":8000", nil)
}
