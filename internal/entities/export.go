package entities

import "strings"

func EscapeBookFileName(n string) string {
	const (
		replacer  = ""
		maxLength = 100
	)

	// TODO: заменить на strings.Replacer
	for _, e := range []string{`\`, `/`, `|`, `:`, `"`, `*`, `?`, `<`, `>`, `.`, "\t"} {
		n = strings.ReplaceAll(n, e, replacer)
	}

	if len([]rune(n)) > maxLength {
		n = string([]rune(n)[:maxLength])
	}

	return n
}
