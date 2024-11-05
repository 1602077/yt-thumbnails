package main

import (
	"fmt"
	"os"

	logger "github.com/1602077/thumbnails/internal"
	"github.com/1602077/thumbnails/internal/cli"
)

func main() {
	logger.Init(logger.DEBUG, "json")

	client, err := cli.FromFlags()
	if err != nil {
		fmt.Printf("failed to create cli client: %v", err)
		os.Exit(1)
	}

	if err := client.Run(); err != nil {
		fmt.Printf("failed to download thumbnail: %v", err)
		os.Exit(1)
	}
}
