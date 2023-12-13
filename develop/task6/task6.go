/*
Реализовать утилиту -- аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
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
	// Определение флагов командной строки
	fields := flag.Int("f", 0, "выбрать номер поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	onlySeparated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Чтение строк из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка, содержит ли строка разделитель
		if !strings.Contains(line, *delimiter) {
			if *onlySeparated {
				continue
			}
			fmt.Println(line)
			continue
		}

		// Разбивка строки на поля
		fieldsSlice := strings.Split(line, *delimiter)

		// Выбор и вывод запрошенного поля
		if *fields == 0 {
			fmt.Println(line)
		} else if *fields > 0 && *fields <= len(fieldsSlice) {
			fmt.Println(fieldsSlice[*fields-1])
		} else {
			fmt.Fprintf(os.Stderr, "Ошибка: поля под номером %d нет в строке\n", *fields)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении STDIN:", err)
		os.Exit(1)
	}
}
