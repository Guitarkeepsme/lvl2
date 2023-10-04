package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/* "a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
*/

type Unpacker interface {
	Unpack()
}

type PackedString string

func main() {
	var pkdString PackedString

	fmt.Println("Введите строку для распаковки: ")

	_, err := fmt.Scanf("%s", &pkdString)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Распакованная строка: ", pkdString.Unpack())
	}

}

func (s PackedString) Unpack() string {
	var lastRune, lastLetter rune
	var res, num strings.Builder
	var esc bool

	res.Reset()
	num.Reset()
	lastRune = 0
	lastLetter = 0

	for i, rune := range s {
		if unicode.IsDigit(rune) && i == 0 {
			return ""
		}
		if unicode.IsLetter(rune) {
			// letter after digit
			if unicode.IsLetter(lastRune) {
				numRunes, err := strconv.Atoi(num.String())
				if err != nil {
					log.Fatal(err)
				}
				for j := 0; j < numRunes-1; j++ {
					res.WriteRune(lastLetter)
				}
				num.Reset()
			}
			// any letter
			res.WriteRune(rune)
			lastLetter = rune
			lastRune = rune
		}
		// write to digit sequence or flush letters to result
		if unicode.IsDigit(rune) {
			// espace digit
			if esc {
				res.WriteRune(rune)
				lastLetter = rune
				lastRune = rune
				esc = false
			} else {
				// first digit of new digit sequence
				if unicode.IsLetter(lastRune) {
					num.Reset()
				}
				num.WriteRune(rune)
				lastRune = rune
				// last digit in input string
				if i == utf8.RuneCountInString(string(s))-1 {
					numRunes, err := strconv.Atoi(num.String())
					if err != nil {
						log.Fatal(err)
					}
					for j := 0; j < numRunes-1; j++ {
						res.WriteRune(lastLetter)
					}
				}

			}
		}
		if rune == '\\' {
			if lastRune == '\\' {
				res.WriteRune(rune)
				lastLetter = rune
				lastRune = rune
				esc = false
			} else {
				esc = false
				lastRune = rune
			}
		}

	}

	return res.String()
}
