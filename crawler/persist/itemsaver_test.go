package persist

import (
	"context"
	"encoding/json"
	"learn/crawler/engine"
	"learn/crawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v3"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://bj.zu.anjuke.com/fangyuan/1244278867",
		Type: "anjuke",
		Id:   "1244278867",
		PayLoad: model.Profile{
			Title:  "松北区观江国际B区拎包入住精装价格值得拥有",
			Price:  2300,
			Type:   "面议",
			Door:   "2室1厅1卫",
			Area:   106,
			Photos: []string{"https://pic1.ajkimg.com/display/hj/fe9bb5890596bca2e2904243e30d2385/600x450c.jpg?t=1"},
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = Save(client, expected, "crawler")
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("crawler").
		Type(expected.Type).
		Id(expected.Id).
		DoC(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.PayLoad)

	actual.PayLoad = actualProfile

	t.Logf("%s", actual)
	t.Logf("%s", expected)
	if actual != expected {
		t.Errorf("Got %v; expected %v", actual, expected)
	}
}
