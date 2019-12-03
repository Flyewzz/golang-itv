package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// ----------------------------------------------------------------------------
// *****************************          *************************************
// ***************************** Handlers *************************************
// *****************************          *************************************
// ----------------------------------------------------------------------------

func MainFunc(w http.ResponseWriter, r *http.Request) {
	method, url := r.URL.Query().Get("method"),
		r.URL.Query().Get("url")
	fmt.Printf("method: %s, url: %s\n", method, url)
	timeout := 5 * time.Second
	resp, err := SendRequest(&http.Client{
		Timeout: timeout,
	}, method, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		log.Println(err)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		log.Println(err)
		return
	}
	w.Write(data)
}

func HeadersToString(reqHeader *http.Header) string {
	var result []string
	for name, headers := range *reqHeader {
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v\n", name, h))
		}
	}
	return strings.Join(result, "\n")
}
