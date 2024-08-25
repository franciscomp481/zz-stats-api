package webscrapper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/franciscomp481/zerozero-stats-api/model"
)

func SearchClub(filters model.ClubFilters) (string, error) {
	baseURL := "https://www.zerozero.pt/equipas"

	encondedClubName := EncodeName(filters.ClubName)

	searchURL := fmt.Sprintf("%s?search_txt=%s&order=popular", baseURL, encondedClubName)

	// Send the GET Resquest
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

	selection := doc.Find(".zz-search-item").Eq(filters.Index).Find("a")
	firstURL, exists := selection.Attr("href")
	if !exists {
		return "", fmt.Errorf("no results found at index %d", filters.Index)
	}

	fullURL := fmt.Sprintf("https://www.zerozero.pt%s", firstURL)

	return fullURL, nil
}
