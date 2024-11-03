package thumbnails

import "net/url"

type ThumbnailDownloader interface {
	GetThumbnail(url url.URL, filename string) error
}
