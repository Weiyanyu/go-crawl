package fetcher

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Fetch(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error status code : %d", resp.StatusCode)
	}
	return goquery.NewDocumentFromReader(resp.Body)
}
