package engine

import "github.com/PuerkitoBio/goquery"

type Request struct {
	Url        string
	ParserFunc func(*goquery.Document) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(*goquery.Document) ParseResult {
	return ParseResult{}
}
