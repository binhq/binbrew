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

// protobuf_goarch matches arch representations.
func protobuf_goarch(s string) string {
	switch s {
	case "386":
		return "x86_32"

	case "amd64":
		return "x86_64"
	}

	return s
}

// protobuf_goos matches os representations.
func protobuf_goos(s string) string {
	switch s {
	case "darwin":
		return "osx"
	}

	return s
}
