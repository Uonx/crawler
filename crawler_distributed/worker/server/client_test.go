package main

import (
	"fmt"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/rpcsupport"
	"learn/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlerService(t *testing.T) {
	const host string = ":9000"

	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error(err)
	}

	req := worker.Request{
		Url: "https://bj.zu.anjuke.com/fangyuan/1237268666",
		Parser: worker.SerializedParser{
			Name: config.ParseFangYuanDetail,
			Args: "双周保洁 物业齐全 智能门锁双桥",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println(result)
	}

}
