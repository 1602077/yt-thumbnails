package cli

import (
	"flag"
	"fmt"
	"net/url"

	logger "github.com/1602077/thumbnails/internal"
	"github.com/1602077/thumbnails/internal/thumbnails"
)

var (
	urlPtr             = flag.String("url", "", "url of thumbnail icon to download")
	directoryPtr       = flag.String("directory", "/Users/jcmunday/thumbnails/", "directory to download thumbnails to")
	thumbnailStemPtr   = flag.String("stem", "https://i.ytimg.com/vi/", "stem of youtube thumbnail hosting")
	thumbnailSuffixPtr = flag.String("suffix", "hq720.jpg", "suffix of youtube thumbnail hosting")
	filenamePtr        = flag.String("filename", "", "filename to save image under")
)

// FromFlags creates a new CLI client from flags.
func FromFlags() (CLI, error) {
	flag.Parse()

	thumbnailSuffix, err := url.Parse(*thumbnailStemPtr)
	if err != nil {
		return CLI{}, err
	}

	downloader := thumbnails.HttpThumbnailDownloader{
		Config: &thumbnails.ThumbnailDownloaderConfig{
			ThumbnailStem:     *thumbnailSuffix,
			ThumbnailSuffix:   *thumbnailSuffixPtr,
			DownloadDirectory: *directoryPtr,
		},
	}

	logger.Debug(
		"downloader initialised",
		"config",
		fmt.Sprintf("%+v", downloader.Config),
	)

	return CLI{
		ThumbnailDownload: &downloader,
	}, nil
}
