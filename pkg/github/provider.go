package github

import (
	"github.com/Masterminds/semver"
	"github.com/binhq/binbrew/pkg"
	"github.com/pkg/errors"
)

type Provider struct {
}

func (p *Provider) Find(fullName string, version string) (*pkg.Binary, error) {
	// Try parsing version as an exact version
	v, err := semver.NewVersion(version)
	if err != nil {
		_, err := semver.NewConstraint(version)
		if err != nil {
			return nil, err
		}

		// fetch the latest version which matches the constraint
		return nil, errors.New("fetching latest version is not implemented yet")
	}

	return rules.Resolve(fullName, v, nil)
}
