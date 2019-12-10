package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1:18081/index", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("X-Token", "Y2Y1YzBlYWQtNjAwYy00YWExLWIxN2EtMjE5YjhmMDU3YWIwLTQ0MjQwLTMxNjY=")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	//resp, err := http.Get("http://www.baidu.com")
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
