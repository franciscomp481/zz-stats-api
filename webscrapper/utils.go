package webscrapper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

// Search for the player by name and return the URL

func GetPage(playerURL string) (*goquery.Document, error) {

	// Fetch the HTML content of the player's profile page
	resp, err := http.Get(playerURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body with the correct encoding
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	return doc, nil
}

func EncodeName(name string) string {
	encodedName := strings.ReplaceAll(name, " ", "+")
	encodedName = strings.ReplaceAll(encodedName, "ç", "%E7")
	encodedName = strings.ReplaceAll(encodedName, "á", "%E1")
	encodedName = strings.ReplaceAll(encodedName, "à", "%E0")
	encodedName = strings.ReplaceAll(encodedName, "ã", "%E3")
	return encodedName
}
