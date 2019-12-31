package parser

import (
	"go-crawl/engine"

	"github.com/PuerkitoBio/goquery"
)

const cityListSelection string = ".city-list>dd>a"
const htmlHrefAttr string = "href"

func ParseCityList(doc *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	doc.Find(cityListSelection).Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr(htmlHrefAttr)
		if !ok {
			return
		}
		result.Requests = append(result.Requests, engine.Request{Url: url, ParserFunc: ParseCityAndProfile})
	})
	return result
}
