package template

// goarch matches arch representations.
func goarch(s string) string {
	switch s {
	case "386":
		return "i386"

	case "amd64":
		return "x86_64"
	}

	return s
}
