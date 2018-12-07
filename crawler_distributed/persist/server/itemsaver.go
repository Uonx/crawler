package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/persist"
	"learn/crawler_distributed/rpcsupport"
	"log"
)

var port =flag.Int("port",0,"the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
	}
	err := serverRpc(fmt.Sprintf(":%d",*port), config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serverRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}