package parser

import (
	"learn/crawler/fetcher"
	"testing"
)

func TestParseFangYuan(t *testing.T) {
	contents, err := fetcher.Fetche("https://le.zu.anjuke.com/") //ioutil.ReadFile("fangyuanlist_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", contents)

	result := ParseFangYuan(contents)

	resultSize := 60

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}
}
