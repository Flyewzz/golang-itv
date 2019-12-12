package features

import (
	"fmt"
	"net/http"
	"strings"
)

func HeadersToString(reqHeader *http.Header) string {
	var result []string
	for name, headers := range *reqHeader {
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v\n", name, h))
		}
	}
	return strings.Join(result, "\n")
}
