package env

import "unicode"

func camelToSlice(s string) []string {
	pcs := []string{}
	start := 0
	for end, r := range s {
		if end != 0 && unicode.IsUpper(r) {
			pcs = append(pcs, s[start:end])
			start = end
		}
	}
	if start != len(s) {
		pcs = append(pcs, s[start:])
	}

	return pcs
}
