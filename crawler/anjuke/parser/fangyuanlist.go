package parser

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"regexp"
)

var (
	fangyuanRe     = regexp.MustCompile(`_soj="Filter_[0-9]*[^"]+"[^href]+href="(https://[a-z]+.zu.anjuke.com/fangyuan/[0-9]+)"[^>]+>([^<]+)</a>`)
	fangyuanPageRe = regexp.MustCompile(`href="(https://[a-z]+.zu.anjuke.com/fangyuan/p[0-9]*/)" class="aNxt"`)
)

func ParseFangYuan(contents []byte, _ string) engine.ParseResult {
	matches := fangyuanRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//log.Printf("cityListCount: %d", len(matches))
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			Paser: NewProfileParser(string(m[2])),
		})
	}
	matches = fangyuanPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:   string(m[1]),
			Paser: engine.NewFuncParser(ParseFangYuan, config.ParseFangYuan),
		})
	}
	return result
}

type ProfileParser struct {
	Title string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseFangYuanDetail(contents, p.Title, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseFangYuanDetail, p.Title
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		Title:name,
	}
}
