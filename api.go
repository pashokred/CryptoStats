package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	gClient  *http.Client
	gBaseURL string = "https://api.kuna.io"
)

func GetLatestPrice(market string) (p float64, err error) {
	gClient = &http.Client{Timeout: time.Duration(1) * time.Second}
	r, err := gClient.Get(fmt.Sprintf("%s/v3/tickers?symbols=%s", gBaseURL, market))
	defer r.Body.Close()
	if err != nil {
		return p, err
	}
	return decodeValue(r)
}

func decodeValue(r *http.Response) (p float64, err error) {
	m, err := jsonGetSlice(r)
	if err != nil {
		return p, err
	}
	last, err := jsonGetValue(m[7])
	if err != nil {
		return p, err
	}
	return last, nil
}

func jsonGetSlice(r *http.Response) ([]interface{}, error) {
	if r == nil {
		return nil, errors.New(
			"expected dict but NIL found")
	}
	var j []interface{}
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		return nil, fmt.Errorf("json decoder error")
	}

	var v []interface{}
	switch j[0].(type) {
	case []interface{}:
		v = j[0].([]interface{})
		return v, nil
	}
	return nil, fmt.Errorf(
		"expected dict but %#v (%T) found", v, v)
}

func jsonGetValue(v interface{}) (float64, error) {
	if v == nil {
		return 0, errors.New("expected float but nil found")
	}
	switch v.(type) {
	case float64:
		return v.(float64), nil
	case string:
		return strconv.ParseFloat(v.(string), 64)
	}
	return 0, fmt.Errorf(
		"expected float but %#v (%T) found", v, v)
}
