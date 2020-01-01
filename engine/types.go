package engine

import "github.com/PuerkitoBio/goquery"

type Request struct {
	Url        string
	ParserFunc func(*goquery.Document) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParser(*goquery.Document) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChannel() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
