package client

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got Saver:got item "+
				"#%d: %v", itemCount, item)
			itemCount++
			result := ""
			err = client.Call(config.ItemSaverRpc,
				item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
