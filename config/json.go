package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadJson() *strings.Reader {
	body, err := ioutil.ReadFile("../config/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	b := strings.NewReader(string(body))
	return b
}

func WriteToJson(b bytes.Buffer) {
	file, err := os.Create("../files/notion.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSON ファイルに書き込み
	encoder := json.NewEncoder(file)
	fmt.Println(&b)
	if err := encoder.Encode(&b); err != nil {
		log.Fatal(err)
	}
}
