package thumbnails

import "net/url"

type Downloader interface {
	GetThumbnail(url url.URL, filename string) error
}
