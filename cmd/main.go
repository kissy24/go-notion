package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-notion/config"
	"go-notion/pkg/api"
)

// Retrieve page information from Notion.
func retrievePageSample(notion api.Notion) {
	req := notion.RetrievePage()
	b := api.GetResponse(req)
	output(b)
}

func createPageSample(notion api.Notion) {
	f := "../files/create_page.json"
	json := config.ReadJson(f)
	req := notion.CreatePage(json)
	b := api.GetResponse(req)
	output(b)
}

func updatePageSample(notion api.Notion) {
	f := "../files/update_page.json"
	json := config.ReadJson(f)
	req := notion.UpdatePage(json)
	b := api.GetResponse(req)
	output(b)
}

func searchSample(notion api.Notion) {
	f := "../files/search.json"
	json := config.ReadJson(f)
	req := notion.Search(json)
	b := api.GetResponse(req)
	output(b)

}

func output(body []byte) {
	var out bytes.Buffer
	json.Indent(&out, body, "", "  ")
	fmt.Println(&out)
}

func main() {
	secret := config.Config{}.ParseToml()
	n := api.Notion{ApiKey: secret.ApiKey, DatabaseId: secret.DatabaseId}
	retrievePageSample(n)
	createPageSample(n)
	updatePageSample(n)
	searchSample(n)
}
