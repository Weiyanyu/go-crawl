package engine

import (
	"go-crawl/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (*SimpleEngine) Run(Seeds ...Request) {
	var requestQ []Request
	for _, request := range Seeds {
		requestQ = append(requestQ, request)
	}

	for len(requestQ) > 0 {
		request := requestQ[0]
		requestQ = requestQ[1:]
		parseResult, err := worker(request)
		if err != nil {
			continue
		}
		requestQ = append(requestQ, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item : %v\n", item)
		}
	}
}

func worker(request Request) (ParseResult, error) {
	log.Println("fetching url : ", request.Url)

	//fetch
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("fetch url error. url %s : %v", request.Url, err)
		return ParseResult{}, err
	}

	//parse
	return request.ParserFunc(doc), nil
}
