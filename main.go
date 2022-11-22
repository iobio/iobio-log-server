package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
)

type IobioReq struct {
	RequestId   string `json:"requestId"`
	Type        string `json:"type"`
	NumAttempts int    `json:"numAttempts"`
	Endpoint    string `json:"endpoint"`
}

func main() {

	port := flag.Int("port", 9001, "Port")
	flag.Parse()

	http.HandleFunc("/eGJvfRfF300fGpxnB52LmFpD9IIJPzYb", func(w http.ResponseWriter, r *http.Request) {

                headers := w.Header()

                headers["Access-Control-Allow-Origin"] = []string{"*"}
                headers["Access-Control-Allow-Methods"] = []string{"POST, OPTIONS"}

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

		timestamp := time.Now().Format(time.RFC3339)
		remoteIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%d", timestamp, remoteIp, iobioReq.RequestId, iobioReq.Type, iobioReq.Endpoint, iobioReq.NumAttempts))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
