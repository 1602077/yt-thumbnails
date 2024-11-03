package thumbnails

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetThumbnail(t *testing.T) {
	thumbnailStem, err := url.Parse("https://i.ytimg.com/vi/")
	assert.NoError(t, err)

	downloader := HttpThumbnailDownloader{
		&ThumbnailDownloaderConfig{
			ThumbnailStem:     *thumbnailStem,
			ThumbnailSuffix:   "hq720.jpg",
			DownloadDirectory: "/tmp/thumbnails",
		},
	}

	thumbnailURL, err := url.Parse("https://www.youtube.com/watch?v=2AwmwORrepc")
	assert.NoError(t, err)

	t.Run("can download a thumbnail without error", func(t *testing.T) {
		err = downloader.GetThumbnail(*thumbnailURL, "tinzo-dark-synthwave.jpg")
		assert.NoError(t, err, "failed to download thumbnail")
	})

	t.Run("errors if nil filename", func(t *testing.T) {
		err = downloader.GetThumbnail(*thumbnailURL, "")
		assert.Error(t, err, "expected to error as no filename provided")
		assert.ErrorIs(t, err, ErrInvalidFilename)
	})
}

func TestBuildThumbnailURL(t *testing.T) {
	thumbnailStem, err := url.Parse("https://i.ytimg.com/vi/")
	assert.NoError(t, err)

	downloader := HttpThumbnailDownloader{
		&ThumbnailDownloaderConfig{
			ThumbnailStem:     *thumbnailStem,
			ThumbnailSuffix:   "hq720.jpg",
			DownloadDirectory: "/tmp/thumbnails",
		},
	}

	t.Run("can build a url from valid youtube url", func(t *testing.T) {
		youtubeURL, err := url.Parse("https://www.youtube.com/watch?v=2AwmwORrepc")
		assert.NoError(t, err)

		actualURL, err := downloader.buildThumbnailURL(*youtubeURL)
		assert.NoError(t, err)

		expectedURL, err := url.Parse("https://i.ytimg.com/vi/2AwmwORrepc/hq720.jpg")
		assert.NoError(t, err)

		assert.Equal(t, expectedURL, actualURL)
	})
}
