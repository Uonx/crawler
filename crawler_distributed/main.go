package main

import (
	"flag"
	"learn/crawler/anjuke/parser"
	"learn/crawler/engine"
	"learn/crawler/scheduler"
	"learn/crawler_distributed/config"
	itemSaver "learn/crawler_distributed/persist/client"
	"learn/crawler_distributed/rpcsupport"
	worker "learn/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

const url = "https://www.anjuke.com/sy-city.html"

var(
	ItemSaverHost = flag.String("itemsaver_host","","itemsaver host")

	workerHosts = flag.String("worker_hosts","","worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	if *ItemSaverHost == "" {
		log.Println("must specify a port")
	}
	if *workerHosts == "" {
		log.Println("must specify ports")
	}
	itemChan, err := itemSaver.ItemSaver(*ItemSaverHost)

	if err != nil {
		panic(err)
	}

	pool := CreateClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{}, //&scheduler.SimpleScheduler{},//Scheduler接口实现类
		WorkerCount:      100,                          //创建work的数量
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:   url,                                                              //url地址
		Paser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), //每个请求需要执行的方法
	})

}

func CreateClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)

		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", host)
		} else {
			log.Printf("error connecting to %s: %v", host, err)
		}
	}
	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
