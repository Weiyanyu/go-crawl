package engine

import (
	"go-crawl/fetcher"
	"log"
)

func Run(Seeds ...Request) {
	var requestQ []Request
	for _, request := range Seeds {
		requestQ = append(requestQ, request)
	}

	for len(requestQ) > 0 {
		request := requestQ[0]
		requestQ = requestQ[1:]

		log.Println("fetching url : ", request.Url)

		doc, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("fetch url error. url %s : %v", request.Url, err)
			continue
		}

		parseResult := request.ParserFunc(doc)

		requestQ = append(requestQ, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item : %v\n", item)
		}
	}
}
