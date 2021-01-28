package main

import (
	"fmt"
	"golearn/basic_study/retriever/mock"
	ireal "golearn/basic_study/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)

}
func post(poster Poster) string {
	return poster.Post(url,
		map[string]string{
			"name":   "hackfun",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another facked hackfun.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fack hackfun.com"}
	inspect(&retriever)

	r = &ireal.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	// mockRetriever := r.(mock.Retriever)
	// fmt.Println(mockRetriever.Contents)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	// realRetriever := r.(*real.Retriever)
	// fmt.Println(realRetriever.TimeOut)

	// fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *ireal.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
