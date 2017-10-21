package pkg_test

import (
	"testing"

	"github.com/Masterminds/semver"
	"github.com/binhq/binbrew/pkg"
	"github.com/binhq/binbrew/pkg/template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryTemplate_Resolve(t *testing.T) {
	tests := map[string]struct {
		template *pkg.BinaryTemplate
		fullName string
		version  *semver.Version
		ctx      map[string]string
		binary   *pkg.Binary
		err      string
	}{
		"Defaults": {
			&pkg.BinaryTemplate{
				URL:  template.ParseNew("http://example.com/name-{{.Version}}.tar.gz"),
				File: template.ParseNew("name-{{.Version}}"),
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			&pkg.Binary{
				Name:     "name",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/name-1.0.0.tar.gz",
				File:     "name-1.0.0",
			},
			"",
		},
		"AllFields": {
			&pkg.BinaryTemplate{
				Homepage:    "http://example.com",
				Description: "Lorem ipsum dolor",
				URL:         template.ParseNew("http://example.com/name-{{.Version}}.tar.gz"),
				File:        template.ParseNew("name-{{.Version}}"),
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			&pkg.Binary{
				Name:        "name",
				FullName:    "repo/name",
				Homepage:    "http://example.com",
				Description: "Lorem ipsum dolor",
				Version:     semver.MustParse("1.0.0"),
				URL:         "http://example.com/name-1.0.0.tar.gz",
				File:        "name-1.0.0",
			},
			"",
		},
		"OverridesName": {
			&pkg.BinaryTemplate{
				Name: "alias",
				URL:  template.ParseNew("http://example.com/name-{{.Version}}.tar.gz"),
				File: template.ParseNew("name-{{.Version}}"),
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			&pkg.Binary{
				Name:     "alias",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/name-1.0.0.tar.gz",
				File:     "name-1.0.0",
			},
			"",
		},
		"Context": {
			&pkg.BinaryTemplate{
				URL:  template.ParseNew("http://example.com/{{.FullName}}-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz"),
				File: template.ParseNew("{{.Name}}-{{.Version}}"),
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			map[string]string{
				pkg.CONTEXT_OS:   "darwin",
				pkg.CONTEXT_ARCH: "amd64",
			},
			&pkg.Binary{
				Name:     "name",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/repo/name-1.0.0-darwin-amd64.tar.gz",
				File:     "name-1.0.0",
			},
			"",
		},
		"ContextOverride": {
			&pkg.BinaryTemplate{
				URL:  template.ParseNew("http://example.com/{{.FullName}}-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz"),
				File: template.ParseNew("{{.Name}}-{{.Version}}"),
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			map[string]string{
				pkg.CONTEXT_OS:       "darwin",
				pkg.CONTEXT_ARCH:     "amd64",
				pkg.CONTEXT_NAME:     "lorem",
				pkg.CONTEXT_FULLNAME: "ipsum/lorem",
				pkg.CONTEXT_VERSION:  "2.0.0",
			},
			&pkg.Binary{
				Name:     "name",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/repo/name-1.0.0-darwin-amd64.tar.gz",
				File:     "name-1.0.0",
			},
			"",
		},
		"MissingUrlTemplate": {
			&pkg.BinaryTemplate{},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			nil,
			"missing URL template",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := test.template.Resolve(test.fullName, test.version, test.ctx)

			if test.err != "" {
				require.Error(t, err)
				assert.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.binary, b)
			}
		})
	}
}

func TestRuleSet_Resolve(t *testing.T) {
	tests := map[string]struct {
		ruleSet  pkg.RuleSet
		fullName string
		version  *semver.Version
		ctx      map[string]string
		binary   *pkg.Binary
	}{
		"Defaults": {
			pkg.RuleSet{
				"repo/name": []*pkg.Rule{
					{
						Constraint: pkg.MustConstraint("^1.0.0"),
						Template: &pkg.BinaryTemplate{
							URL:  template.ParseNew("http://example.com/name-{{.Version}}.tar.gz"),
							File: template.ParseNew("name-{{.Version}}"),
						},
					},
				},
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			&pkg.Binary{
				Name:     "name",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/name-1.0.0.tar.gz",
				File:     "name-1.0.0",
			},
		},
		"Fallback": {
			pkg.RuleSet{
				"repo/name": []*pkg.Rule{
					{
						Constraint: pkg.MustConstraint("^2.0.0"),
						Template: &pkg.BinaryTemplate{
							URL:  template.ParseNew("http://anotherexample.com/name-{{.Version}}.tar.gz"),
							File: template.ParseNew("name-{{.Version}}"),
						},
					},
					{
						Constraint: pkg.MustConstraint("*"),
						Template: &pkg.BinaryTemplate{
							URL:  template.ParseNew("http://example.com/name-{{.Version}}.tar.gz"),
							File: template.ParseNew("name-{{.Version}}"),
						},
					},
				},
			},
			"repo/name",
			semver.MustParse("1.0.0"),
			nil,
			&pkg.Binary{
				Name:     "name",
				FullName: "repo/name",
				Version:  semver.MustParse("1.0.0"),
				URL:      "http://example.com/name-1.0.0.tar.gz",
				File:     "name-1.0.0",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := test.ruleSet.Resolve(test.fullName, test.version, test.ctx)

			require.NoError(t, err)
			assert.Equal(t, test.binary, b)
		})
	}
}
