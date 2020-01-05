package worker

import (
	"errors"
	"go-crawl/distributed/config"
	"go-crawl/engine"
	"go-crawl/zhenai/parser"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(request engine.Request) Request {
	name, args := request.Parser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(result engine.ParseResult) ParseResult {
	res := ParseResult{
		Items: result.Items,
	}

	for _, req := range result.Requests {
		res.Requests = append(res.Requests, SerializeRequest(req))
	}

	return res
}

func DeserializeRequest(req Request) (engine.Request, error) {

	parser, err := deserializeParser(req.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    req.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(res ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: res.Items,
	}

	for _, req := range res.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCityAndProfile:
		return engine.NewFuncParser(parser.ParseCityAndProfile, config.ParseCityAndProfile), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknow parser name")
	}
}
