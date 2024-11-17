package main

import (
	"github.com/1602077/thumbnails/internal/cli"
	"github.com/1602077/thumbnails/pkg/logger"
)

func main() {
	logger.Init(logger.DEBUG, "json")

	client, err := cli.FromFlags()
	if err != nil {
		logger.Fatal(
			"failed to create cli client",
			"error", err,
		)
	}

	if err := client.Run(); err != nil {
		logger.Fatal(
			"failed to download thumbnail",
			"error", err,
		)
	}
}
