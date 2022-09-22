package helpers

import "unicode"

func LowerFirstChar(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
