package pkg_test

import (
	"testing"

	"github.com/binhq/binbrew/pkg"
	"github.com/binhq/binbrew/pkg/github"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestResolver_Resolve(t *testing.T) {
	resolver := &pkg.Resolver{
		Providers: map[string]pkg.Provider{
			"github": &github.Provider{},
		},
	}

	binary, err := resolver.Resolve("golang/dep@0.3.2")

	require.NoError(t, err)
	assert.Equal(t, "dep", binary.Name)
}
