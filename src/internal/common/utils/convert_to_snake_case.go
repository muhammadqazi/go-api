package utils

import (
	"bytes"
	"unicode"
)

func ConvertToSnakeCase(input string) string {
	var buf bytes.Buffer
	var last rune
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 && last != '_' && !unicode.IsUpper(last) {
				buf.WriteRune('_')
			}
			r = unicode.ToLower(r)
		}
		buf.WriteRune(r)
		last = r
	}
	return buf.String()
}
