package pkg

import (
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter"
	"github.com/sirupsen/logrus"
)

type Downloader struct {
	cache  *Cache
	logger logrus.FieldLogger
}

// NewDownloader returns a new Downloader instance.
func NewDownloader(cache *Cache, logger logrus.FieldLogger) *Downloader {
	return &Downloader{
		cache:  cache,
		logger: logger,
	}
}

// Download downloads a file if necessary (not found in cache) and copies it to the target directory.
func (d *Downloader) Download(binary *Binary, dst string) error {
	d.logger.Debug("looking for binary in the cache")

	if d.cache.Miss(binary) {
		d.logger.Debug("cache missed")

		err := d.cache.Prepare(binary)
		if err != nil {
			return err
		}

		d.logger.Debug("downloading binary")

		err = getter.GetAny(d.cache.Path(binary), binary.URL)
		if err != nil {
			return err
		}
	} else {
		d.logger.Debug("cache hit")
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
