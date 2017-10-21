package pkg

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/hashicorp/go-getter"
	"io"
)

type Downloader struct {
}

func (d *Downloader) Download(binary *Binary, dst string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	home := u.HomeDir
	if home == "" {
		home = os.TempDir()
	}

	cache := filepath.Join(home, ".binbrew", "cache", binary.FullName, binary.Version.String())

	err = os.MkdirAll(cache, 0744)
	if err != nil {
		return err
	}

	err = getter.GetAny(cache, binary.URL)
	if err != nil {
		return err
	}

	from, err := os.Open(filepath.Join(cache, binary.File))
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
