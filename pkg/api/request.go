package api

import (
	"log"
	"net/http"
	"strings"
)

const PAGES = "https://api.notion.com/v1/pages/"
const SEARCH = "https://api.notion.com/v1/search"

type Notion struct {
	ApiKey     string
	DatabaseId string
}

func (n Notion) RetrievePage() *http.Request {
	req, err := http.NewRequest("GET", PAGES+n.DatabaseId, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+n.ApiKey)
	req.Header.Add("Notion-Version", "2022-06-28")
	return req
}

func (n Notion) CreatePage(body *strings.Reader) *http.Request {
	req, err := http.NewRequest("POST", PAGES+n.DatabaseId, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+n.ApiKey)
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")
	return req
}

func (n Notion) UpdatePage(body *strings.Reader) *http.Request {
	req, err := http.NewRequest("PATCH", PAGES+n.DatabaseId, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+n.ApiKey)
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")
	return req
}

func (n Notion) Search(body *strings.Reader) *http.Request {
	req, err := http.NewRequest("POST", SEARCH+n.DatabaseId, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+n.ApiKey)
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")
	return req
}
