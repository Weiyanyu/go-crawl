package engine

import (
	"go-crawl/fetcher"
	"log"
)

func Worker(request Request) (ParseResult, error) {
	log.Println("fetching url : ", request.Url)

	//fetch
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("fetch url error. url %s : %v", request.Url, err)
		return ParseResult{}, err
	}

	//parse
	return request.Parser.Parse((doc)), nil
}
