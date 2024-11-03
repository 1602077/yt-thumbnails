package thumbnails

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

// HttpThumbnailDownloader downloads thumbnails over HTTP.
type HttpThumbnailDownloader struct {
	*ThumbnailDownloaderConfig
}

type ThumbnailDownloaderConfig struct {
	ThumbnailStem     url.URL
	ThumbnailSuffix   string
	DownloadDirectory string
}

func (h *HttpThumbnailDownloader) GetThumbnail(thumbnailURL url.URL, filename string) error {
	if filename == "" {
		return ErrInvalidFilename
	}

	thumbnailToDownload, err := h.buildThumbnailURL(thumbnailURL)
	if err != nil {
		return err
	}

	resp, err := http.Get(thumbnailToDownload.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrFailedToDownloadImage
	}

	if err = os.MkdirAll(h.DownloadDirectory, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path.Join(h.DownloadDirectory, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}

func (h *HttpThumbnailDownloader) buildThumbnailURL(youtubeURL url.URL) (*url.URL, error) {
	videoID := youtubeURL.Query().Get("v")

	thumbnailURLStr, err := url.JoinPath(h.ThumbnailStem.String(), videoID, h.ThumbnailSuffix)
	if err != nil {
		return &url.URL{}, err
	}

	thumbnailURL, err := url.Parse(thumbnailURLStr)
	if err != nil {
		return &url.URL{}, err
	}

	return thumbnailURL, nil
}
