package crawler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/slogger/pkg/slogger"
)

func getEbayNlPage(watcher *api.QueueResponse) (Response, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"),
	)

	c.OnRequest(func(r *colly.Request) {
		slogger.Debug(fmt.Sprintf("Visiting %s", r.URL.String()))
	})

	response := Response{}
	response.URL = watcher.URL

	c.OnHTML("span#prcIsum", func(e *colly.HTMLElement) {
		response.Price = transformEbayNlPrice(e.Attr("content"))
	})

	c.OnHTML("h1#itemTitle", func(e *colly.HTMLElement) {
		response.Name = strings.ReplaceAll(e.Text, "Details over", "")
	})

	err := c.Visit(watcher.URL)
	if err != nil {
		return response, err
	}

	return response, nil
}

func transformEbayNlPrice(value string) float32 {
	if value == "" {
		return 0
	}

	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}

	return float32(floatValue)
}
