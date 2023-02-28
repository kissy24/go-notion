package config

import (
	"log"
	"testing"
)

func TestParseToml(t *testing.T) {
	s := Config{}.ParseToml()
	log.Println(s.ApiKey)
	log.Println(s.DatabaseId)
}
