package parser

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"regexp"
)

const cityListRe string = `<a href="(https://[0-9a-zA-Z]+.anjuke.com)">([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:   string(m[1]),
			Paser: engine.NewFuncParser(ParseZuFang, config.ParseZuFang),
		})
	}

	return result
}
