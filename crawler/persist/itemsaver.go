package persist

import (
	"context"
	"errors"
	"learn/crawler/engine"
	"log"

	"gopkg.in/olivere/elastic.v3"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

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
			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		DoC(context.Background())

	if err != nil {
		panic(err)
	}

	return nil
}
