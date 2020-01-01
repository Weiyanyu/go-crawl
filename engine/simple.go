package engine

import (
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
		parseResult, err := Worker(request)
		if err != nil {
			continue
		}
		requestQ = append(requestQ, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item : %v\n", item)
		}
	}
}
