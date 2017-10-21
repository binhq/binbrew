package template

import (
	"text/template"

	"github.com/Masterminds/sprig"
)

var FuncMap template.FuncMap

func init() {
	FuncMap = sprig.TxtFuncMap()

	for name, fn := range funcMap {
		FuncMap[name] = fn
	}
}

// funcMap contains custom functions passed to templates.
var funcMap = template.FuncMap{
	"goarch": goarch,
}
