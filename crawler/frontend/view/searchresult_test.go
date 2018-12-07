package view

import (
	"learn/crawler/engine"
	"learn/crawler/frontend/model"
	common "learn/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://bj.zu.anjuke.com/fangyuan/1244278867",
		Type: "anjuke",
		Id:   "1244278867",
		PayLoad: common.Profile{
			Title:  "松北区观江国际B区拎包入住精装价格值得拥有",
			Price:  2300,
			Type:   "面议",
			Door:   "2室1厅1卫",
			Area:   106,
			Photos: []string{"https://pic1.ajkimg.com/display/hj/fe9bb5890596bca2e2904243e30d2385/600x450c.jpg?t=1"},
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
