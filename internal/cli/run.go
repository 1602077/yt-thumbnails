package cli

import (
	"net/url"

	"github.com/1602077/thumbnails/internal/thumbnails"
)

// CLI wraps the downloading of thumbnails behind a CLI interface.
type CLI struct {
	ThumbnailDownload thumbnails.Downloader
}

func (c *CLI) Run() error {
	thumbnailURL, err := url.Parse(*urlPtr)
	if err != nil {
		return err
	}

	return c.ThumbnailDownload.GetThumbnail(*thumbnailURL, *filenamePtr)
}
