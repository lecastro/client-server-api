package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lecastro/client-server-api/internal/helpers"
)

type dollarPrice struct {
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

	io.Copy(os.Stdout, res.Body)

	body, err := json.Marshal(res.Body)

	if err != nil {
		helpers.HandleError(err)
	}

	persistDollarPriceFile(body)
}

func persistDollarPriceFile(value []byte) {
	var data map[string]dollarPrice

	_ = json.Unmarshal(value, &data)

	log.Println(data)
	f, err := os.Create("./cotacao.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := fmt.Sprintf("Dólar: {%0.2f}\n", data)

	_, err = f.WriteString(str)

	if err != nil {
		panic(err)
	}
}
