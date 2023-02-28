package main

import (
	"bytes"
	"encoding/json"
	"go-notion/config"
	"go-notion/pkg/api"
)

func retrievePageSample() {
	secret := config.Config{}.ParseToml()
	n := api.Notion{ApiKey: secret.ApiKey, DatabaseId: secret.DatabaseId}
	req := n.RetrievePage()
	body := api.GetResponse(req)
	var out bytes.Buffer
	json.Indent(&out, body, "", "  ")
	config.WriteToJson(out)
}

func main() {
	retrievePageSample()
}
