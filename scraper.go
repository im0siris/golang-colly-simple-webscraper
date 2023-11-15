package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector()

	var items []item

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		item := item{
			Name:   h.ChildText("h2"),
			Price:  h.ChildText(".price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	//Mehrere Seiten scrapen

	c.OnHTML("a.page-numbers", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://scrapeme.live/shop/")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}
