package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) (*goquery.Document, error) {
	<-rateLimiter
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
