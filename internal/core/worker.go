package core

import (
	"fmt"
	"time"

	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/pricewatcher-worker/internal/crawler"
	"github.com/laetificat/slogger/pkg/slogger"
	"github.com/spf13/viper"
)

/*
StartWorker starts an infinite loop which gets a job from the queue and runs a crawler, will sleep if queue is empty.
*/
func StartWorker() error {
	for {
		slogger.Debug("Checking for jobs...")
		watcher, err := api.GetJobFromQueue()
		if err != nil {
			return err
		}

		if watcher.ID == 0 {
			slogger.Debug(fmt.Sprintf("Queue is empty, going to sleep for %s minutes", viper.GetDuration("worker.sleep")))
			time.Sleep(viper.GetDuration("worker.sleep") * time.Minute)
			continue
		}

		slogger.Debug("Job found!")
		res, err := crawler.GetPage(watcher)
		if err != nil {
			return err
		}

		req := api.UpdateRequest{
			ID:   watcher.ID,
			Name: res.Name,
			Price: api.Price{
				Value:     res.Price,
				Timestamp: time.Now(),
			},
		}

		slogger.Debug("Updating price...")
		err = api.UpdatePrice(&req)
		if err != nil {
			return err
		}

		time.Sleep(viper.GetDuration("worker.timeout") * time.Second)
	}
}
