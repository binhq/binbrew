package github

import (
	"fmt"
	"text/template"

	"github.com/binhq/binbrew/pkg"
)

const urlPrefix = "https://github.com/{{.FullName}}/releases/download/%s"

var rules = pkg.RuleSet{
	"Masterminds/glide": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://glide.sh",
				Description: "Package Management for Golang",
				URL:         template.Must(template.New("").Funcs(pkg.TplFuncs).Parse(fmt.Sprintf(urlPrefix, "v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz"))),
				File:        template.Must(template.New("").Funcs(pkg.TplFuncs).Parse("{{.Os}}-{{.Arch}}/glide")),
			},
		},
	},
	"mattes/migrate": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Description: "Database migrations. CLI and Golang library.",
				URL:         template.Must(template.New("").Funcs(pkg.TplFuncs).Parse(fmt.Sprintf(urlPrefix, "v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz"))),
				File:        template.Must(template.New("").Funcs(pkg.TplFuncs).Parse("migrate.{{.Os}}-{{.Arch}}")),
			},
		},
	},
	"goreleaser/goreleaser": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint("*"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://goreleaser.github.io/",
				Description: "Deliver Go binaries as fast and easily as possible",
				URL:         template.Must(template.New("").Funcs(pkg.TplFuncs).Parse(fmt.Sprintf(urlPrefix, "v{{.Version}}/goreleaser_{{.Os | title}}_{{.Arch | goarch}}.tar.gz"))),
			},
		},
	},
	"golang/dep": []*pkg.Rule{
		{
			Constraint: pkg.MustConstraint(">0.3.0"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://github.com/golang/dep",
				Description: "Go dependency management tool",
				URL:         template.Must(template.New("").Funcs(pkg.TplFuncs).Parse(fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}"))),
			},
		},
		{
			Constraint: pkg.MustConstraint("<=0.3.0"),
			Template: &pkg.BinaryTemplate{
				Homepage:    "https://github.com/golang/dep",
				Description: "Go dependency management tool",
				URL:         template.Must(template.New("").Funcs(pkg.TplFuncs).Parse(fmt.Sprintf(urlPrefix, "v{{.Version}}/dep-{{.Os}}-{{.Arch}}.zip"))),
			},
		},
	},
}
