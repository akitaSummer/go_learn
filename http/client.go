package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	requset, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	requset.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Mobile Safari/537.36")
	//resp, err := http.DefaultClient.Do(requset)
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request,
		) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(requset)
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
