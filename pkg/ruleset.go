package pkg

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/sprig"
	tpl "github.com/binhq/binbrew/pkg/template"
	"github.com/pkg/errors"
)

var TplFuncs template.FuncMap

func init() {
	TplFuncs = sprig.TxtFuncMap()

	tpl.Merge(TplFuncs)
}

// Predefined context keys.
const (
	CONTEXT_OS       = "Os"
	CONTEXT_ARCH     = "Arch"
	CONTEXT_VERSION  = "Version"
	CONTEXT_NAME     = "Name"
	CONTEXT_FULLNAME = "FullName"
)

// BinaryTemplate represents information used for creating an actual binary instance.
type BinaryTemplate struct {
	// Overrides the automatically calculated name based on the full name of the binary
	Name string

	// Homepage of the binary (if any)
	Homepage string

	// Simple description about what this software does
	Description string

	// Download URL
	URL *template.Template

	// Effective file name (when extracting from an archive)
	//
	// Falls back to binary name when not provided.
	File *template.Template
}

func (t *BinaryTemplate) Resolve(fullName string, version *semver.Version, ctx map[string]string) (*Binary, error) {
	name := t.Name

	// No name override, calculate the name
	if name == "" {
		nameParts := strings.Split(fullName, "/")
		name = nameParts[len(nameParts)-1]
	}

	if ctx == nil {
		ctx = make(map[string]string)
	}

	// Fill the context
	ctx[CONTEXT_NAME] = name
	ctx[CONTEXT_FULLNAME] = fullName
	ctx[CONTEXT_VERSION] = version.String()

	// Fallback to GOOS
	if _, ok := ctx[CONTEXT_OS]; !ok {
		ctx[CONTEXT_OS] = runtime.GOOS
	}

	// Fallback to GOARCH
	if _, ok := ctx[CONTEXT_ARCH]; !ok {
		ctx[CONTEXT_ARCH] = runtime.GOARCH
	}

	buf := new(bytes.Buffer)

	if err := t.URL.Execute(buf, ctx); err != nil {
		return nil, errors.Wrap(err, "cannot render URL template")
	}

	url := buf.String()

	file := name

	if t.File != nil {
		buf.Reset()

		if err := t.File.Execute(buf, ctx); err != nil {
			return nil, errors.Wrap(err, "cannot render File template")
		}

		file = buf.String()
	}

	return &Binary{
		Name:        name,
		FullName:    fullName,
		Homepage:    t.Homepage,
		Description: t.Description,
		Version:     version,
		URL:         url,
		File:        file,
	}, nil
}

// Rule binds a template and it's version condition together.
type Rule struct {
	Constraint *semver.Constraints
	Template   *BinaryTemplate
}

// MustConstraint creates a new constraint and panics on error.
func MustConstraint(c string) *semver.Constraints {
	constraint, err := semver.NewConstraint(c)
	if err != nil {
		panic(err)
	}

	return constraint
}

// RuleSet contains the whole set of builtin rules.
type RuleSet map[string][]*Rule

func (r RuleSet) Resolve(fullName string, version *semver.Version, ctx map[string]string) (*Binary, error) {
	rules, ok := r[fullName]
	if !ok {
		return nil, errors.New("rules not found for binary: " + fullName)
	}

	var currentRule *Rule

	for _, rule := range rules {
		if rule.Constraint.Check(version) {
			currentRule = rule
			break
		}
	}

	if currentRule == nil {
		return nil, errors.New(fmt.Sprintf("cannot find rule for binary: %s@%s", fullName, version.String()))
	}

	return currentRule.Template.Resolve(fullName, version, ctx)
}
