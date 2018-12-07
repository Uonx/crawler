package parser

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"regexp"
)

const zufangRe = `<a class="a_navnew" hidefocus="true" href="(https://[a-z]+.zu.anjuke.com/)" _soj="navigation">([^<]+)</a>`

func ParseZuFang(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(zufangRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			Paser: engine.NewFuncParser( ParseFangYuan,config.ParseFangYuan),
		})
	}

	return result
}
