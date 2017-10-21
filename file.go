package main

import (
	"fmt"
	"net/http"
)

// Obtains statuscode, content-type and content lenght from a specific url
func fileInfo(url string) string {
	response, error := http.Get(url)
	if error != nil {
		return fmt.Sprintf("%s,%s,,\n", url, error.Error())
	}
	headers := response.Header
	statusCode := response.StatusCode
	contentType := headers["Content-Type"][0]
	contentLength := response.ContentLength
	return fmt.Sprintf("%s,%d,%s,%d\n", url, statusCode, contentType, contentLength)
}
