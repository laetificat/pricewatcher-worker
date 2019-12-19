package cmd

import (
	"fmt"

	"github.com/laetificat/pricewatcher-worker/internal/api"
	"github.com/laetificat/slogger/pkg/slogger"

	"github.com/laetificat/pricewatcher-worker/internal/core"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "pricewatcher-worker",
		Short: "A price watcher worker",
		Long:  `Pricewatcher Worker is a worker for scraping webpages and reporting it back to the API.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := core.StartWorker(); err != nil {
				slogger.Fatal(err.Error())
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	registerRootCmd()
	registerListCmd()
	return rootCmd.Execute()
}

/*
registerRootCmd registers the flags and sets up everything for the root command
*/
func registerRootCmd() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		"the config file to use",
	)

	rootCmd.PersistentFlags().StringP(
		"verbose",
		"v",
		"info",
		"the minimum log verbosity level",
	)

	if err := viper.BindPFlag("log.minimum_level", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		slogger.Fatal(err.Error())
	}
	viper.SetDefault("log.minimum_level", "info")

	rootCmd.PersistentFlags().StringP(
		"dsn",
		"d",
		"",
		"the Sentry DSN endpoint to use",
	)

	if err := viper.BindPFlag("log.minimum_level", rootCmd.PersistentFlags().Lookup("dsn")); err != nil {
		slogger.Fatal(err.Error())
	}
	viper.SetDefault("log.sentry.dsn", "")

	rootCmd.PersistentFlags().StringVarP(&api.Queue, "queue", "q", "", "the queue to listen to")
	rootCmd.PersistentFlags().StringVarP(&api.Host, "api", "a", "http://localhost:8080", "the worker API to connect to")
}

/*
initConfig sets up and configures the application
*/
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			slogger.Fatal(err.Error())
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".pricewatcher-worker")
	}

	if err := viper.ReadInConfig(); err != nil {
		slogger.Fatal(err.Error())
	}

	slogger.Debug(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
}
