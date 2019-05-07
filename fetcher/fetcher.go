package main

import (
	"bufio"
	"fmt"
	"io"
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

func (fetcher *Fetcher) Run(spiderUrl string) <-chan byte {
	fetcher.init(spiderUrl)
	data := fetcher.getUrl()
	return data
}

func (fetcher *Fetcher) getUrl() <-chan byte {
	request, err := http.NewRequest(
		http.MethodGet,
		fetcher.SpiderUrl, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	fetcher.Client = client
	resp, err := fetcher.Client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()
	return NetworkSource(resp)
}
func NetworkSource(resp *http.Response) <-chan byte {
	out := make(chan byte)
	go func() {
		r := ReaderSource(bufio.NewReader(resp.Body), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan byte {
	out := make(chan byte, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				for _, v := range buffer {
					out <- v
				}
			}
			if err != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func main() {
	fetcher := NewFetcher()
	uids := []string{"NWQ8NsMsUo3HwFCJrMueaA", "9z9YyZkccMtfekvZpKLQSH", "uDKxEF2mLwDuiSMmcrcxFf", "nNSHZXk35wjPeaQ9BC4Fi3"}
	spiderUrl := "http://testing.ecams-backend.cloudcare.cn/ecams/detector_data?cloud_account_unique_id=%s"
	results := []<-chan byte{}
	for _, uid := range uids {
		data := fetcher.Run(fmt.Sprintf(spiderUrl, uid))
		results = append(results, data)
	}

	for {
		select {
		case v := <-results[0]:
			fmt.Println("0:=", v)
		case v := <-results[1]:
			fmt.Println("1:=", v)
		case v := <-results[2]:
			fmt.Println("2:=", v)
		case v := <-results[3]:
			fmt.Println("3:=", v)
		}
	}
}
