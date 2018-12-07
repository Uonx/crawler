package main

import (
	"learn/crawler/anjuke/parser"
	"learn/crawler/engine"
	"learn/crawler/persist"
	"learn/crawler/scheduler"
	"learn/crawler_distributed/config"
)

const url = "https://www.anjuke.com/sy-city.html"

func main() {
	itemChan, err := persist.ItemSaver("crawler")

	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //&scheduler.SimpleScheduler{},//Scheduler接口实现类
		WorkerCount: 100,                          //创建work的数量
		ItemChan:    itemChan,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url:       url,                  //url地址
		Paser: engine.NewFuncParser(parser.ParseCityList,config.ParseCityList) , //每个请求需要执行的方法
	})

}
