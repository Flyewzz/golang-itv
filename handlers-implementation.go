package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	if strings.ToLower(method) != "get" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := Get(url)
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

func Get(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return NewResponse(1, resp.Status,
		HeadersToString(&resp.Header), string(bodyText), resp.ContentLength), nil
	// defer resp.Body.Close()
	// var body []byte
	// for true {

	// }

}

func HeadersToString(reqHeader *http.Header) string {
	var result []string
	for name, headers := range *reqHeader {
		// name = strings.ToLower(name)
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v\n", name, h))
		}
	}
	return strings.Join(result, "\n")
}

func SearchFunc(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// text := r.URL.Query().Get("data")
	// users, err := SearchUser(GetDbAdapter(), text)
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("404 Not Found"))
	// 	return
	// }
	// sentData, _ := json.Marshal(users)
	// w.Header().Set("Content-type", "application/json")
	// w.Write(sentData)
}
