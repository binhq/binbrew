package pkg

import (
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

type Downloader struct {
	cache *Cache
}

// NewDownloader returns a new Downloader instance.
func NewDownloader(cache *Cache) *Downloader {
	return &Downloader{
		cache: cache,
	}
}

// Download downloads a file if necessary (not found in cache) and copies it to the target directory.
func (d *Downloader) Download(binary *Binary, dst string) error {
	if d.cache.Miss(binary) {
		err := d.cache.Prepare(binary)
		if err != nil {
			return err
		}

		err = getter.GetAny(d.cache.Path(binary), binary.URL)
		if err != nil {
			return err
		}
	}

	// TODO: handle ErrMiss which means download error not handled properly
	from, err := d.cache.GetFile(binary)
	if err != nil {
		return err
	}
	defer from.Close()

	err = os.MkdirAll(dst, 0744)
	if err != nil {
		return err
	}

	to, err := os.OpenFile(filepath.Join(dst, binary.Name), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}
