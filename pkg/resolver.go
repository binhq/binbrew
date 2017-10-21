package pkg

import (
	"regexp"

	"github.com/pkg/errors"
)

const defaultProvider = "github"

const binaryUrlPattern = `^(?:([a-zA-Z0-9_]+):\/\/)?([a-zA-Z0-9-_\.\/]+)(?:@(.+))?$`

var binaryUrlRegex *regexp.Regexp

func init() {
	binaryUrlRegex = regexp.MustCompile(binaryUrlPattern)
}

// Resolver is responsible for resolving a binary URL into an actual binary.
type Resolver struct {
	Providers map[string]Provider
}

// Resolve resolves a binary URL into an actual binary.
func (r *Resolver) Resolve(binaryUrl string) (*Binary, error) {
	match := binaryUrlRegex.FindStringSubmatch(binaryUrl)
	if match == nil {
		return nil, errors.New("cannot parse binary URL")
	}

	providerName, binaryName, binaryVersion := match[1], match[2], match[3]

	// Fall back to the default provider
	if providerName == "" {
		providerName = defaultProvider
	}

	// When not provided, get the latest version
	if binaryVersion == "" || binaryVersion == "latest" {
		binaryVersion = "*"
	}

	// Try to locate the provider
	provider, ok := r.Providers[providerName]
	if !ok {
		return nil, errors.New("unknown provider: " + providerName)
	}

	binary, err := provider.Find(binaryName, binaryVersion)
	if err != nil {
		return binary, err
	}

	return binary, nil
}
