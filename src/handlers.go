package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
)

var url = "https://flathub.org/api/v2/"

type AppHit struct {
	Name              string   `json:"name"`
	Keywords          []string `json:"keywords"`
	Summary           string   `json:"summary"`
	Description       string   `json:"description"`
	ID                string   `json:"id"`
	UpdatedAt         int64    `json:"updated_at"`
	Arches            []string `json:"arches"`
	Type              string   `json:"type"`
	Translations      struct{} `json:"translations"` // Adjust based on actual content if necessary
	ProjectLicense    string   `json:"project_license"`
	IsFreeLicense     bool     `json:"is_free_license"`
	AppID             string   `json:"app_id"`
	Icon              string   `json:"icon"`
	Categories        []string `json:"categories"`
	MainCategories    []string `json:"main_categories"`
	SubCategories     []string `json:"sub_categories"`
	DeveloperName     string   `json:"developer_name"`
	Runtime           string   `json:"runtime"`
	Trending          float64  `json:"trending"`
	InstallsLastMonth int      `json:"installs_last_month"`
	AddedAt           int64    `json:"added_at"`
}

type ServerResponse struct {
	Hits []AppHit `json:"hits"`
}

func SearchRequest(seachText string) ServerResponse {
	searchData := fmt.Sprintf(`{"query": "%s"}`, seachText)
	jsonData := []byte(searchData)
	var resp, err = http.Post(url+"search", "application/json", bytes.NewBuffer(jsonData))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result ServerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
