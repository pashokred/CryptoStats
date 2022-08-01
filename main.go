package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	gMarket = "btcuah"
)

func main() {

	http.HandleFunc("/api/rate", func(w http.ResponseWriter, r *http.Request) {
		price, err := GetLatestPrice(gMarket)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		resp := make(map[string]string)
		priceStr := float64ToString(price)
		resp["mktprice"] = priceStr
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		r.ParseForm()
		var data []string
		for _, value := range r.Form {
			data = append(data, value[0])
		}

		if emailExist, err := saveToFile(data[0]); err != nil {
			log.Fatal(err)
		} else if emailExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})

	http.HandleFunc("/api/sendEmails", func(rw http.ResponseWriter, r *http.Request) {
		price, err := GetLatestPrice(gMarket)
		if err != nil {
			log.Fatal(err)
		}
		if err = sendEmails(price); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8000", nil)
}
