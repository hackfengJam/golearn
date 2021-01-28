package engine

import (
	"golearn/project/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) OutPut(items []interface{}) {
	// TODO to mq
	for _, item := range items {
		log.Printf("Got item %s", item)
	}
}
func (e SimpleEngine) Run(seeds ...Request) {
	var requests [] Request

	// 种子 -> requests
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// Processor
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		// worker: requests -> parseResult{Request, Item}
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// request <- parseResult.Requests...
		requests = append(requests, parseResult.Requests...)

		// OutPut
		e.OutPut(parseResult.Items)
	}
}

func worker(
	r Request) (ParseResult, error) {

	log.Printf("Fetching %s", r.Url)

	// 抓取页面
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
