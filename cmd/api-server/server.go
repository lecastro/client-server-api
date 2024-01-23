package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lecastro/client-server-api/helpers"
)

type dollarPrice struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	http.HandleFunc("/", DollarPrice)
	http.ListenAndServe(":8080", nil)
}

func DollarPrice(w http.ResponseWriter, r *http.Request) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(url)

	if err != nil {
		helpers.HandleError(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		helpers.HandleError(err)
	}

	var data map[string]dollarPrice

	err = json.Unmarshal(body, &data)

	if err != nil {
		helpers.HandleError(err)
	}

	respJson, err := json.Marshal(data["USDBRL"])

	if err != nil {
		w.Write([]byte("Key 'USDBRL' not found JSON."))
	}

	_, err = w.Write(respJson)

	if err != nil {
		fmt.Println("Error writing answer:", err)
	}
}
