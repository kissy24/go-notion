package api

import (
	"io"
	"log"
	"net/http"
)

// Get the response of what you requested
func GetResponse(req *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
