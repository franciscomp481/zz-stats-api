package webscrapper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func SearchPlayer(playerName string, index int) (string, error) {
	baseURL := "https://www.zerozero.pt/jqc_search_solr.php"

	// Construct the URL with the search query
	searchURL := fmt.Sprintf("%s?queryString=%s", baseURL, playerName)

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
	selection := doc.Find("a").Eq(index)
	firstURL, exists := selection.Attr("href")
	if !exists {
		return "", fmt.Errorf("no results found at index %d", index)
	}

	// Prepend the base URL to the relative URL
	fullURL := fmt.Sprintf("https://www.zerozero.pt%s", firstURL)

	return fullURL, nil
}
