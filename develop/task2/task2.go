package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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
	} else if pkdString.Unpack() == "" {
		fmt.Println("Некорректная строка!")
	} else {
		fmt.Println("Распакованная строка: ", pkdString.Unpack())
	}
}

// "a4bc2d5e" => "aaaabccddddde"

func (s PackedString) Unpack() string {
	var (
		sb      strings.Builder
		sbA     strings.Builder
		temp    string
		amount  string
		escaped bool
	)
	for i, v := range s {
		if escaped {
			temp = string(v)
			sb.WriteRune(v)
			escaped = false
			continue
		}
		if _, err := strconv.ParseInt(string(v), 0, 64); err == nil {
			amount = string(v)
			if i == len(s)-1 {
				amountN, _ := strconv.ParseInt(string(amount), 0, 64)
				for i := int64(0); i < amountN-1; i++ {
					sb.WriteString(temp)
					continue
				}
				continue
			}
			sbA.WriteString(amount)
			continue
		} else if amount != "" {
			amountN, _ := strconv.ParseInt(string(amount), 0, 64)
			for i := int64(0); i < amountN-1; i++ {
				sb.WriteString(temp)
				continue
			}
			amount = ""
			sbA.Reset()
		}

		if string(v) == `\` {
			escaped = true
			continue
		}
		temp = string(v)
		sb.WriteRune(v)
	}
	if len(sb.String()) == 0 && len(s) != 0 {
		return ""
	}

	return sb.String()
}
