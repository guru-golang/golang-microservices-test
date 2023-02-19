package strings_lib

import (
	"strings"
)

func Unescape(str string) (result string) {
	// https://www.rapidtables.com/code/text/unicode-characters.html
	result = strings.ReplaceAll(str, "\\u002b", "+")
	result = strings.ReplaceAll(result, "\\u002f", "/")
	result = strings.ReplaceAll(result, "\\u003a", ":")

	return result
}
