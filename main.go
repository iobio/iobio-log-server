package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caddyserver/certmagic"
)

type IobioReq struct {
	Type        string `json:"type"`
	NumAttempts int    `json:"numAttempts"`
	Endpoint    string `json:"endpoint"`
}

func main() {

	http.HandleFunc("/eGJvfRfF300fGpxnB52LmFpD9IIJPzYb", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Invalid HTTP method")
			return
		}

		var iobioReq IobioReq

		err := json.NewDecoder(r.Body).Decode(&iobioReq)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			return
		}

		fmt.Println(iobioReq)
	})

	err := certmagic.HTTPS([]string{"logs.anderspitman.net"}, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
