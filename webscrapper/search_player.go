package webscrapper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franciscomp481/zerozero-stats-api/model"
)

func SearchPlayer(filters model.PlayerFilters) (string, error) {
	baseURL := "https://www.zerozero.pt/jogadores"

	// Replace spaces with '+' and handle special characters manually
	encodedPlayerName := strings.ReplaceAll(filters.PlayerName, " ", "+")
	encodedPlayerName = strings.ReplaceAll(encodedPlayerName, "ç", "%E7")
	searchURL := ""
	nationality := strings.ToLower(filters.Nationality)

	if nationality == "" {
		searchURL = fmt.Sprintf("%s?search_txt=%s&order=popular", baseURL, encodedPlayerName)
	} else {
		searchURL = fmt.Sprintf("%s/%s?search_txt=%s&order=popular", baseURL, nationality, encodedPlayerName)
	}
	// Construct the URL with the search query

	// Send the GET request
	resp, err := http.Get(searchURL)
	if err != nil {
		return "", fmt.Errorf("error fetching search results: %v", err)
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing search results: %v", err)
	}

	// Extract the URL from the specified index in the search results
	selection := doc.Find(".zz-search-item").Eq(filters.Index).Find("a")
	firstURL, exists := selection.Attr("href")
	if !exists {
		return "", fmt.Errorf("no results found at index %d", filters.Index)
	}

	// Prepend the base URL to the relative URL
	fullURL := fmt.Sprintf("https://www.zerozero.pt%s", firstURL)

	return fullURL, nil
}
