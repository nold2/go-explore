package utils

import "strings"

// MakeExcited transform a string to uppercase and append "!"
func MakeExcited(sentence string) string {
	upper := strings.ToUpper(sentence)
	return upper + "!"

}
