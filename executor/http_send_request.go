package executor

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Flyewzz/golang-itv/features"
	"github.com/Flyewzz/golang-itv/models"
)

func (ex HttpExecutor) Execute(client *http.Client, method, url string, id int) (*models.Response, error) {
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
	return models.NewResponse(id, resp.Status,
		features.HeadersToString(&resp.Header), string(bodyText), resp.ContentLength), nil

}