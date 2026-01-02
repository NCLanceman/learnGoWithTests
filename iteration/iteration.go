package iteration

import "strings"

//const repeatCount = 5

func Repeat(character string, repeats int) string {
	repeatCount := repeats
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
