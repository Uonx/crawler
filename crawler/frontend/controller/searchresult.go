package controller

import (
	"context"
	"learn/crawler/engine"
	"learn/crawler/frontend/model"
	"learn/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v3"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	//fmt.Fprintf(w, "q = %s, from = %d", q, from)

	page, err := h.GetSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) GetSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("crawler").Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).From(from).DoC(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = int(resp.TotalHits())
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.Query = q
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	s := re.ReplaceAllString(q, "PayLoad.$1:")
	return s
}
