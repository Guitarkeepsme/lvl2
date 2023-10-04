package main

import (
	"fmt"
	"strings"
	"unicode"
)

/* "a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
*/

func main() {
	str := "a4bc2d5e"

	runes := []rune(str)

	res := strings.Builder{}

	for i := 0; i < len(str); i++ {
		// fmt.Println(string(str[i]))
		if unicode.IsDigit(runes[i]) {
			fmt.Println(string(runes[i]))
			res.WriteRune(runes[i])
		}
	}
	fmt.Println(res)

}
