/*
	Утилита grep

Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "Печатать +N строк после совпадения")
	before := flag.Int("B", 0, "Печатать +N строк до совпадения")
	context := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения (до и после)")
	count := flag.Bool("c", false, "Печатать количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invert := flag.Bool("v", false, "Вместо совпадения исключать")
	fixed := flag.Bool("F", false, "Точное совпадение со строкой")
	lineNum := flag.Bool("n", false, "Печатать номер строки")

	flag.Parse()

	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	if pattern == "" {
		fmt.Println("Необходима опция")
		return
	}

	if len(files) == 0 {
		fmt.Println("Введите команду")
		return
	}

	for _, file := range files {
		err := filterFile(file, pattern, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)
		if err != nil {
			fmt.Printf("Ошибка при фильтрации файла %s: %s\n", file, err)
		}
	}
}

func filterFile(file, pattern string, after, before, context int, count, ignoreCase, invert, fixed, lineNum bool) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	matchCount := 0
	lineCount := 0
	lines := []string{}

	mapLines := make(map[int]string)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
		lineCount++

		mapLines[lineCount] = line

		if ignoreCase {
			if fixed && strings.EqualFold(line, pattern) {
				handleMatch(line, lineCount, count, lineNum)
				matchCount++
			} else if !fixed && strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				handleMatch(line, lineCount, count, lineNum)
				matchCount++
			} else if invert {
				handleNonMatch(line, lineCount, count, lineNum)
			}
		} else {
			if fixed && line == pattern {
				handleMatch(line, lineCount, count, lineNum)
				matchCount++
			} else if !fixed && strings.Contains(line, pattern) {
				handleMatch(line, lineCount, count, lineNum)
				matchCount++
			} else if invert {
				handleNonMatch(line, lineCount, count, lineNum)
			}
		}

		test := []string{}

		if after > 0 && lineCount > after {

			for key, value := range mapLines {
				if key < lineCount {
					continue
				}
				for afterCount := 0; afterCount < after; afterCount++ {
					// fmt.Println(value)
					test = append(test, value)
				}

			}

			fmt.Println(strings.Join(test, "\n"))

			// fmt.Println("Ok!")
			lines = lines[1:]
			// fmt.Println(lines)
		}

		if before > 0 || context > 0 {
			lines = append(lines, line)
			if len(lines) > before+context {
				lines = lines[:before+context]
			}
		}
	}

	fmt.Println(mapLines)

	if count {
		fmt.Printf("Совпадений в файле %s: %d\n", file, matchCount)
	}

	return nil
}

func handleMatch(line string, lineCount int, count, lineNum bool) {
	if count {
		return
	}

	if lineNum {
		fmt.Printf("%d: %s\n", lineCount, line)
	} else {
		fmt.Println(line)
	}
}

func handleNonMatch(line string, lineCount int, count, lineNum bool) {
	if count {
		return
	}

	if lineNum {
		fmt.Printf("%d: %s\n", lineCount, line)
	} else {
		fmt.Println(line)
	}
}
