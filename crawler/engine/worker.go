package engine

import (
	"learn/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {

	log.Printf("feting url : %s", r.Url)
	body, err := fetcher.Fetche(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.Paser.Parse(body, r.Url), nil
}
