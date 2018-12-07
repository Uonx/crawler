package parser

import (
	"learn/crawler/fetcher"
	"testing"
)

func TestParseZuFang(t *testing.T) {
	contents, err := fetcher.Fetche("https://beijing.anjuke.com/")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", contents)

	result := ParseZuFang(contents)

	resultSize := 1

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}
}
