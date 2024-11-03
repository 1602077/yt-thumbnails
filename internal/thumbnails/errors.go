package thumbnails

import "errors"

var (
	ErrFailedToDownloadImage = errors.New("failed to download image")
	ErrInvalidFilename       = errors.New("filename is invalid")
)
