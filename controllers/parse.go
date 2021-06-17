package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type PageInfo struct {
	Url       string
	Title     string
	Tag       string
	Views     float64
	Bookmarks float64
	Comments  float64
}

func ParseHabr() {
	var items []PageInfo

	c := colly.NewCollector(
	// colly.MaxDepth(2),
	// colly.Async(),
	)

	pageC := c.Clone()

	visited := []string{
		"https://habr.com/ru/",
	}

	c.OnHTML("a.toggle-menu__item-link_pagination", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("href"))

		for _, item := range visited {
			if item == url {
				return
			}
		}

		visited = append(visited, url)
		c.Visit(url)

	})

	c.OnHTML("h2.post__title a[href]", func(e *colly.HTMLElement) {
		pageC.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	pageC.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("Parse page ", e.Request.URL.String())

		e.ForEach(".post__tags .inline-list__item_tag", func(i int, h *colly.HTMLElement) {
			items = append(items, PageInfo{
				Url:       e.Request.URL.String(),
				Tag:       strings.TrimSpace(h.Text),
				Title:     e.ChildText("h1"),
				Views:     ParseFloat(e.ChildText(".post-stats__views-count")),
				Bookmarks: ParseFloat(e.ChildText(".bookmark__counter")),
				Comments:  ParseFloat(e.ChildText(".post-stats__comments-count")),
			})
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://habr.com/ru/")

	fmt.Println("save items: ", len(items))
	save(items)
}

func save(data []PageInfo) {
	file, _ := os.Create("result.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"Title",
		"Url",
		"Views",
		"Bookmarks",
		"Comments",
	})

	for _, value := range data {
		writer.Write([]string{
			fmt.Sprintf("%v", value.Title),
			fmt.Sprintf("%v", value.Url),
			fmt.Sprintf("%v", value.Views),
			fmt.Sprintf("%v", value.Bookmarks),
			fmt.Sprintf("%v", value.Comments),
		})
	}
}

func ParseFloat(vl string) float64 {

	if strings.Contains(vl, "k") {
		vl = convertString(vl)
	}

	s, err := strconv.ParseFloat(vl, 64)
	if err != nil {
		return 0
	}
	return s
}

func convertString(s string) string {
	r := regexp.MustCompile(`([0-9])`)
	res := make([]string, 0)
	for _, item := range r.FindAllStringSubmatch(s, -1) {
		res = append(res, item[0])
	}
	res = append(res, "000")

	return strings.Join(res, "")
}
