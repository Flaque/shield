package tokenizer

import "unicode"

func IsWhitespace(str string) bool {
	for _, v := range str {
		if !unicode.IsSpace(v) {
			return false
		}
	}

	return true
}
