package features

import "strings"

func CheckMethodValid(name string) bool {
	allowedMethods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
	}
	upperName := strings.ToUpper(name)
	for _, method := range allowedMethods {
		if upperName == method {
			return true
		}
	}
	return false
}
