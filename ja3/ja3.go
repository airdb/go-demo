package ja3

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CUCyber/ja3transport"
)

type Browser struct {
	Ja3       string
	UserAgent string
}

func Req(browser Browser) {
	tr, _ := ja3transport.NewTransport(browser.Ja3)
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", "https://ja3er.com/json", nil)
	req.Header.Set("User-Agent", browser.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Proto)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
