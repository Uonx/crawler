package worker

import (
	"errors"
	"fmt"
	"learn/crawler/anjuke/parser"
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"log"
)

//{"ProfileParser",name}
type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(request engine.Request) Request {
	name, args := request.Paser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeParseResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request,error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{},err
	}
	return engine.Request{
		Url:   r.Url,
		Paser: parser,
	},nil
}

func DeserializeParseResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items:r.Items,
	}

	for _,req:=range r.Requests{
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialize request: %v",err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser,error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),nil
	case config.ParseZuFang:
		return engine.NewFuncParser(parser.ParseZuFang,config.ParseZuFang),nil
	case config.ParseFangYuan:
		return engine.NewFuncParser(parser.ParseFangYuan,config.ParseFangYuan),nil
	case config.ParseFangYuanDetail:
		if title ,ok:=p.Args.(string);ok{
			return parser.NewProfileParser(title),nil
		}else {
			return nil,fmt.Errorf("invalid arg: %v",p.Args)
		}
	default:
		return nil,errors.New("unknown parser name"+p.Name)
	}
}

