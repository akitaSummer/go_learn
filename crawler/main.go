package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)

	e := determineEncoding(resp.Body)

	reader := transform.NewReader(resp.Body, e.NewDecoder())

	//all, err := ioutil.ReadAll(resp.Body)
	all, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}