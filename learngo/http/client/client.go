package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//resp, err := http.Get("https://shmiloveu.fun")
	req, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0.1; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect to", req)
			return nil
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response)
	fmt.Println(req.UserAgent())
}
