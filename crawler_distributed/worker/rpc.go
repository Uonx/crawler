package worker

import "learn/crawler/engine"

type CrawlerService struct {}

func (CrawlerService) Process(request Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(request)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(engineResult)
	return nil
}
