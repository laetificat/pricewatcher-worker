package main

import (
	"github.com/laetificat/pricewatcher-worker/cmd"
	"github.com/laetificat/slogger/pkg/slogger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		slogger.Fatal(err.Error())
	}
}
