package thumbnails

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	logger "github.com/1602077/thumbnails/internal"
)

// HttpThumbnailDownloader downloads thumbnails over HTTP.
type HttpThumbnailDownloader struct {
	Config *ThumbnailDownloaderConfig
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

	logger.Debug(
		"requesting download of thumbnail",
		"url", thumbnailURL.String(),
		"filename", filename,
	)

	thumbnailToDownload, err := h.buildThumbnailURL(thumbnailURL)
	if err != nil {
		logger.Error(
			"failed to build thumbnail url",
			"error", err,
		)
		return err
	}

	resp, err := http.Get(thumbnailToDownload.String())
	if err != nil {
		logger.Error(
			"failed to download thumbnail",
			"error", err,
		)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error(
			fmt.Sprintf("%v", ErrFailedToDownloadImage),
			"status_code", resp.StatusCode,
		)
		return ErrFailedToDownloadImage
	}

	if err = os.MkdirAll(h.Config.DownloadDirectory, os.ModePerm); err != nil {
		logger.Error(
			"failed to create download directory",
			"error", err,
		)
		return err
	}

	file, err := os.Create(path.Join(h.Config.DownloadDirectory, filename))
	if err != nil {
		logger.Error(
			"failed to create download filename",
			"error", err,
		)
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		logger.Error(
			"failed to write response to filename",
			"error", err,
		)
		return err
	}

	logger.Info("thumbnail downloaded")
	return nil
}

func (h *HttpThumbnailDownloader) buildThumbnailURL(youtubeURL url.URL) (*url.URL, error) {
	videoID := youtubeURL.Query().Get("v")

	thumbnailURLStr, err := url.JoinPath(h.Config.ThumbnailStem.String(), videoID, h.Config.ThumbnailSuffix)
	if err != nil {
		return &url.URL{}, err
	}

	thumbnailURL, err := url.Parse(thumbnailURLStr)
	if err != nil {
		return &url.URL{}, err
	}

	logger.Debug(
		"built thumbnail url",
		"url", thumbnailURL.String(),
	)
	return thumbnailURL, nil
}
