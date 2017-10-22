package pkg

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

// ErrMiss is returned when the cache is accessed when it's empty.
var ErrMiss = errors.New("cache missed")

type Cache struct {
	dir string
}

// NewCache returns a new cache instance.
func NewCache() *Cache {
	var dir string

	u, err := user.Current()
	if err != nil {
		dir = os.TempDir()
	} else {
		dir = u.HomeDir
	}

	return &Cache{
		dir: filepath.Join(dir, ".binbrew", "downloads"),
	}
}

// Path returns the cache path for a binary.
func (c *Cache) Path(binary *Binary) string {
	return filepath.Join(c.dir, binary.FullName, binary.Version.String())
}

// File returns the cached file path for a binary.
func (c *Cache) File(binary *Binary) string {
	return filepath.Join(c.Path(binary), binary.File)
}

// Hit checks whether the binary can be found in the cache.
func (c *Cache) Hit(binary *Binary) bool {
	stat, err := os.Stat(c.File(binary))

	// Check if the file exists or stat failed
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}

	// Check if the file is actually a directory
	if stat.IsDir() {
		return false
	}

	return true
}

// Miss checks whether a binary cannot be found in the cache.
func (c *Cache) Miss(binary *Binary) bool {
	return !c.Hit(binary)
}

// GetFile returns the binary file.
//
// Returns ErrMiss when accessing the cache when it's empty.
func (c *Cache) GetFile(binary *Binary) (*os.File, error) {
	if c.Hit(binary) == false {
		return nil, ErrMiss
	}

	file, err := os.Open(c.File(binary))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Prepare prepares the cache directory for a binary.
func (c *Cache) Prepare(binary *Binary) error {
	return os.MkdirAll(c.Path(binary), 0744)
}
