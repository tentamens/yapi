package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

var url = "https://flathub.org/api/v2/"

func SearchRequest(seachText string) string {
	jsonData := []byte(`{"query": "discord"}`)
	var resp, err = http.Post(url+"search", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln(err)
		return "error"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)

		return "othererr"
	}
	return string(body)
}
