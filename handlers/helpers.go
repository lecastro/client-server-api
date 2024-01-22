package handlers

import (
	"fmt"
	"net"
	"net/url"
)

func HandleError(err error) {
	switch e := err.(type) {
	case net.Error:
		fmt.Println("Error de rede:", e)
	case *url.Error:
		fmt.Println("Error de URL:", e)
	default:
		fmt.Println("Error generic:", e)
	}
}
