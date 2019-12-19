package cmd

import (
	"fmt"

	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/slogger/pkg/slogger"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available queues",
		Run: func(cmd *cobra.Command, args []string) {
			queues, err := api.GetAvailableQueues()
			if err != nil {
				slogger.Fatal(err.Error())
				return
			}

			if len(queues) == 0 {
				fmt.Println("There are no queues available.")
				return
			}

			fmt.Println("Available queues:")
			for _, v := range queues {
				fmt.Println(v)
			}
		},
	}
)

func registerListCmd() {
	rootCmd.AddCommand(listCmd)
}
