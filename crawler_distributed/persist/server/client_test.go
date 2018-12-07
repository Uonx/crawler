package main

import (
	"learn/crawler/engine"
	"learn/crawler/model"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemServer(t *testing.T) {
	const host string = ":1234"
	//go serverRpc(host, "test1")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
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
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
