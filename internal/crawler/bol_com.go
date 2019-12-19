package crawler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/slogger/pkg/slogger"
)

func getBolComPage(watcher *api.QueueResponse) (Response, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"),
	)

	c.OnRequest(func(r *colly.Request) {
		slogger.Debug(fmt.Sprintf("Visiting %s", r.URL.String()))
	})

	response := Response{}
	response.URL = watcher.URL

	c.OnHTML("span.promo-price", func(e *colly.HTMLElement) {
		response.Price = transformValueBolCom(e.Text)
	})

	c.OnHTML("h1.page-heading", func(e *colly.HTMLElement) {
		response.Name = e.ChildText("span.h-boxedright--xs")
	})

	err := c.Visit(watcher.URL)
	if err != nil {
		return response, err
	}

	return response, nil
}

func transformValueBolCom(value string) float32 {
	if value == "" {
		return 0
	}

	value = strings.Replace(value, "\n", ".", 1)
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\n", "")
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "-", "")

	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}

	return float32(floatValue)
}
