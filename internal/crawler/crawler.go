package crawler

import (
	"fmt"

	"github.com/laetificat/pricewatcher-worker/internal/api"
)

// Response is the API response model.
type Response struct {
	Name  string
	Price float32
	URL   string
}

/*
GetPage calls the scraper for the correct domain.
*/
func GetPage(watcher *api.QueueResponse) (Response, error) {
	switch domain := watcher.Domain; domain {
	case "bol.com":
		return getBolComPage(watcher)
	case "ebay.nl":
		return getEbayNlPage(watcher)
	case "coolblue.nl":
		return getCoolblueNlPage(watcher)
	}

	return Response{}, fmt.Errorf("shop type '%s' not found", watcher.Domain)
}
