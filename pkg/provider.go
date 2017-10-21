package pkg

import "github.com/Masterminds/semver"

// Binary represents information for a specific version.
type Binary struct {
	// Name of the binary
	Name string

	// Full name of the binary (as in the binary URL)
	FullName string

	// Homepage of the binary (if any)
	Homepage string

	// Simple description about what this software does
	Description string

	// Binary version
	Version *semver.Version

	// Download URL
	URL string

	// Effective file name (when extracting from an archive)
	//
	// Falls back to binary name when not provided.
	File string
}

// Provider returns a binary.
type Provider interface {
	Find(fullName string, version string) (*Binary, error)
}
