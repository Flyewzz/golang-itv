package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRequest(client *http.Client, method, url string) (*Response, error) {
	var resp *http.Response
	method = strings.ToUpper(method)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return NewResponse(1, resp.Status,
		HeadersToString(&resp.Header), string(bodyText), resp.ContentLength), nil

}
