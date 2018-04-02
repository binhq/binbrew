package github

import (
	"fmt"

	"github.com/binhq/binbrew/pkg"
	"github.com/binhq/binbrew/pkg/template"
)

const urlPrefix = "https://github.com/{{.FullName}}/releases/download/%s"

var rules = pkg.RuleSet{
	"Masterminds/glide": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://glide.sh",
				Description: "Package Management for Golang",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz")),
				File:        template.ParseNew("{{.Os}}-{{.Arch}}/glide"),
			},
		},
	},
	"mattes/migrate": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Description: "Database migrations. CLI and Golang library.",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz")),
				File:        template.ParseNew("migrate.{{.Os}}-{{.Arch}}"),
			},
		},
	},
	"goreleaser/goreleaser": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://goreleaser.github.io/",
				Description: "Deliver Go binaries as fast and easily as possible",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/goreleaser_{{.Os | title}}_{{.Arch | goarch}}.tar.gz")),
			},
		},
	},
	"golang/dep": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint(">0.3.0"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://github.com/golang/dep",
				Description: "Go dependency management tool",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}")),
				File:        template.ParseNew("dep-{{.Os}}-{{.Arch}}"),
			},
		},
		{
			Constraint: pkg.MustConstraint("<=0.3.0"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://github.com/golang/dep",
				Description: "Go dependency management tool",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}.zip")),
			},
		},
	},
	"gobuffalo/packr": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Description: "The simple and easy way to embed static files into Go binaries.",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/packr_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz")),
			},
		},
	},
	"google/protobuf": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Name:        "protoc",
				Description: "Protocol Buffers - Google's data interchange format",
				URL:         template.ParseNew(fmt.Sprintf(urlPrefix, "v{{.Version}}/protoc-{{.Version}}-{{.Os|protobuf_goos}}-{{.Arch|protobuf_goarch}}.zip")),
				File:        template.ParseNew("bin/protoc"),
			},
		},
	},
}
