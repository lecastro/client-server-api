package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/lecastro/client-server-api/internal/helpers"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		helpers.HandleError(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		helpers.HandleError(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
