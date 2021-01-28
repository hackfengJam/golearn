package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Fetcher struct {
	SpiderUrl string
	Data      string
	Host      string
	path      string
	Client    *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{}
}

func (fetcher *Fetcher) init(spiderUrl string) {
	fetcher.SpiderUrl = spiderUrl
	_url, err := url.Parse(fetcher.SpiderUrl)
	if err != nil {
		panic(err)
	}
	fetcher.Host = _url.Host
	fetcher.path = _url.Path
	fetcher.Data = ""

	queryString := _url.RawQuery
	if queryString != "" {
		fetcher.path = strings.Join([]string{fetcher.path, queryString}, "?")
	}

	if fetcher.path == "" {
		fetcher.path = "/"
	}
}

func (fetcher *Fetcher) Run(spiderUrl string) ([]byte, error) {
	fetcher.init(spiderUrl)
	return fetcher.getUrl()
}

func (fetcher *Fetcher) getUrl() ([]byte, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		fetcher.SpiderUrl, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	fetcher.Client = client
	fmt.Println("send: ", fetcher.SpiderUrl)
	resp, err := fetcher.Client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	fetcher := NewFetcher()
	uids := []string{"NWQ8NsMsUo3HwFCJrMueaA", "9z9YyZkccMtfekvZpKLQSH", "uDKxEF2mLwDuiSMmcrcxFf", "nNSHZXk35wjPeaQ9BC4Fi3"}
	spiderUrl := "http://testing.ecams-backend.cloudcare.cn:8087/ecams/detector_data?cloud_account_unique_id=%s"
	results := make(chan []byte)

	for _, uid := range uids {
		go func(uid string) {
			v, err := fetcher.Run(fmt.Sprintf(spiderUrl, uid))
			if err != nil {
				panic(err)
			}
			results <- v
		}(uid)
	}
	for v := range results {
		for _, i := range v {
			fmt.Print(string(i))
		}
		fmt.Println()
	}

}
