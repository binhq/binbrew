package template

import "text/template"

// funcMap contains custom functions passed to templates.
var funcMap = template.FuncMap{
	"goarch": goarch,
}

func Merge(targetFuncMap template.FuncMap) {
	for name, fn := range funcMap {
		targetFuncMap[name] = fn
	}
}
