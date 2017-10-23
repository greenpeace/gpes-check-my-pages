package main

import (
	"fmt"
	"net/http"
)

// getHTTPinfoAsCsvline Obtains statuscode, content-type, content lenght and final URL from a specific http get request
func getHTTPinfoAsCsvline(url string) string {
	response, error := http.Get(url)
	if error != nil {
		return fmt.Sprintf("%s,%s,,\n", url, error.Error())
	}
	headers := response.Header
	statusCode := response.StatusCode
	contentType := headers["Content-Type"][0]
	contentLength := response.ContentLength
	finalURL := response.Request.URL.String()
	return fmt.Sprintf("%s,%d,%s,%d,%s\n", url, statusCode, contentType, contentLength, finalURL)
}
