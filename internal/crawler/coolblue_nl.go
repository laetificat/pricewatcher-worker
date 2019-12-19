package crawler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/slogger/pkg/slogger"
)

func getCoolblueNlPage(watcher *api.QueueResponse) (Response, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"),
	)

	c.OnRequest(func(r *colly.Request) {
		slogger.Debug(fmt.Sprintf("Visiting %s", r.URL.String()))
	})

	response := Response{}
	response.URL = watcher.URL

	c.OnHTML(".sales-price__current", func(e *colly.HTMLElement) {
		response.Price = transformCoolblueNlPrice(e.Text)
	})

	c.OnHTML("span.js-product-name", func(e *colly.HTMLElement) {
		response.Name = e.Text
	})

	err := c.Visit(watcher.URL)
	if err != nil {
		return response, err
	}

	return response, nil
}

func transformCoolblueNlPrice(value string) float32 {
	value = strings.ReplaceAll(value, ",-", "")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")

	if value == "" {
		return 0
	}

	valFloat, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}

	return float32(valFloat)
}
