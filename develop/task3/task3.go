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

	filename string
)

func Flags() {
	flag.IntVar(&column, "k", 0, "выберите колонку для сортировки")
	flag.BoolVar(&number, "n", false, "сортировать по номеру")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&month, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&backspace, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&check, "c", false, "проверять, отсортированные ли данные")
	flag.BoolVar(&numberSuff, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.StringVar(&filename, "f", "", "путь до файла")
	flag.Parse()
}

func readFile(filename string) (res []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		if column > 0 {
			strArr := strings.Split(scan.Text(), " ")
			fmt.Print(len(strArr))
			fmt.Println(strArr)
			res = append(res, strArr[column])
		} else {
			res = append(res, scan.Text())
		}
	}

	sort.Strings(res)

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

	return
}

// func setColumn(column int) {
// }

func main() {
	Flags()
	result, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
