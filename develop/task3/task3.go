/* Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать
в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column     int
	number     bool
	reverse    bool
	unique     bool
	month      bool
	backspace  bool
	check      bool
	numberSuff bool
	filename   string
)

func Flags() {
	// Основные функции
	flag.IntVar(&column, "k", 0, "выбрать колонку для сортировки")
	flag.BoolVar(&number, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")

	// Дополнительные функции
	flag.BoolVar(&month, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&backspace, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&check, "c", false, "проверять, отсортированы ли данные")
	flag.BoolVar(&numberSuff, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.StringVar(&filename, "f", "", "путь до файла")
	flag.Parse()
}

type ByColumn struct {
	lines []string
	colmn int
}

func (b ByColumn) Len() int {
	return len(b.lines)
}

func (b ByColumn) Swap(i, j int) {
	b.lines[j], b.lines[i] = b.lines[i], b.lines[j]
}

func (b ByColumn) Less(i, j int) bool {
	colmnI := strings.Fields(b.lines[i])[b.colmn-1]
	colmnJ := strings.Fields(b.lines[j])[b.colmn-1]
	return colmnI < colmnJ
}

func sortByColumn(strings []string, column int) {
	sort.Sort(ByColumn{strings, column})
}

func readFile(filename string) (res []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		res = append(res, scan.Text())
	}

	sort.Strings(res)

	if column > 0 {
		sort.Sort(ByColumn{res, column})
	}

	if unique {
		uniqueLines := make(map[string]bool)
		for _, line := range res {
			uniqueLines[line] = true
		}

		res = []string{}

		for line := range uniqueLines {
			res = append(res, line)
		}
	}

	if number {
		ints := make([]int, len(res))

		for i, s := range res {
			ints[i], _ = strconv.Atoi(s)
		}

		sort.Ints(ints)

		res = []string{}

		for i := range ints {
			number := ints[i]
			text := strconv.Itoa(number)
			res = append(res, text)
		}
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(res)))
	}

	// if month {

	// }

	return
}

func main() {
	Flags()
	result, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
