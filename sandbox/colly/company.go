package colly

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func SpiderCompanyInfo() {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	/*
		err := c.SetProxy("http://116.108.202.236:41113")
		if err != nil {
			log.Println(err)
			panic("set proxy fail")
		}
	*/

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Set HTML callback
	// Won't be called if error occurs
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	// Set error handler
	c.OnResponse(func(r *colly.Response) {
		log.Println("status code:", r.StatusCode)
		log.Println(string(r.Body))
		// http request IP tag service
	})

	log.Println("start...")
	c.Visit("https://www.sgpbusiness.com/company/Marnix-Reinsurance-Brokers-Pte-Ltd")
}
