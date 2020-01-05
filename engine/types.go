package engine

import "github.com/PuerkitoBio/goquery"

type ParserFunc func(*goquery.Document) ParseResult

type Parser interface {
	Parse(*goquery.Document) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

//NilParser
type NilParser struct {
}

func (NilParser) Parse(d *goquery.Document) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "Nil", nil
}

//Function Parser
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(d *goquery.Document) ParseResult {
	return f.parser(d)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {

	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
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
