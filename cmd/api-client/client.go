package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lecastro/client-server-api/internal/helpers"
)

type CurrencyData struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		helpers.HandleError(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		helpers.HandleError(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.HandleError(err)
	}

	fmt.Println("Corpo da resposta:", string(body))

	persistDollarPriceFile(body)
}

func persistDollarPriceFile(body []byte) {
	var data CurrencyData

	_ = json.Unmarshal(body, &data)

	f, err := os.Create("./cotacao.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := fmt.Sprintf("Dollar: %s\n", data.Bid)

	_, err = f.WriteString(str)

	if err != nil {
		panic(err)
	}
}
