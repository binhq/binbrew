// Package template provides helpers for binary templates.
package template

import "text/template"

// ParseNew parses a new template and sets the function map.
func ParseNew(tpl string) *template.Template {
	return template.Must(template.New("").Funcs(FuncMap).Parse(tpl))
}
