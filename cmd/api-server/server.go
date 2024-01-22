package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", DollarPrice)
	http.ListenAndServe(":8080", nil)
}

func DollarPrice(w http.ResponseWriter, r *http.Request) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(url)

	if err != nil {
		w.Write([]byte("Error"))
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)

	// if err != nil {

	// }

	w.Write([]byte("ola"))
}
