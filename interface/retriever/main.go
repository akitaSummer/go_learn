package main

import (
	"fmt"
	"learn_go/interface/retriever/mock"
	"learn_go/interface/retriever/real"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "akita",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println(r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "this is a fake baidu.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))

	s := &mock.Retriever{
		Contents: "this is a fake baidu.com",
	}
	fmt.Println(session(s))
	fmt.Println(s)
}
