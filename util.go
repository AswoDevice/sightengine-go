package sightengine_go

func join(a []Model, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return string(a[0])
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(a[0]) + sep + string(a[1])
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(a[0]) + sep + string(a[1]) + sep + string(a[2])
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
