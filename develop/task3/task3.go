/* Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
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
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Line struct {
	Text string
	Key  string
}

type ByKey []Line

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func main() {
	// Объявляем и инициализируем флаги командной строки
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	M := flag.Bool("M", false, "сортировать по названию месяца")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверять, отсортированы ли данные")
	h := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Формат использования: sort [-k column] [-n] [-r] [-u] [-M] [-b] [-c] [-h] имя файла")
		return
	}

	fileName := flag.Args()[0]

	lines, err := readLines(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Применяем указанные ключи для сортировки строк
	sortedLines := sortLines(lines, *k, *n, *r, *u, *M, *b, *c, *h)

	// Записываем отсортированные строки в файл
	err = writeLines(sortedLines, fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Файл успешно отсортирован.")
}

func readLines(fileName string) ([]Line, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, Line{Text: line})
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func sortLines(lines []Line, k int, n, r, u, M, b, c, h bool) []Line {
	for i := range lines {
		text := lines[i].Text

		// Если указан ключ -b, игнорируем хвостовые пробелы
		if b {
			text = strings.TrimRight(text, " ")
		}

		columns := strings.Fields(text)

		if k > 0 && k <= len(columns) {
			lines[i].Key = columns[k-1]
		} else {
			lines[i].Key = text
		}
	}

	if n {
		sort.Slice(lines, func(i, j int) bool {
			num1, err1 := strconv.ParseFloat(lines[i].Key, 64)
			num2, err2 := strconv.ParseFloat(lines[j].Key, 64)

			if err1 == nil && err2 == nil {
				return num1 < num2
			}

			// Если указан ключ -h, сортируем числовые значения с учетом суффиксов
			if h {
				num1 = parseNumericValue(lines[i].Key)
				num2 = parseNumericValue(lines[j].Key)
				return num1 < num2
			}

			return lines[i].Key < lines[j].Key
		})
	} else if M {
		sort.Slice(lines, func(i, j int) bool {
			fmt.Println(lines)
			// Если указан ключ -M, сортируем по названию месяца

			// Это пока не работает
			date1, err1 := parseMonthName(lines[i].Key)
			date2, err2 := parseMonthName(lines[j].Key)

			if err1 == nil && err2 == nil {
				return date1.Before(date2)
			}

			return lines[i].Key < lines[j].Key
		})
	} else {
		sort.Sort(ByKey(lines))
	}

	if r {
		reverse(lines)
	}

	if u {
		lines = removeDuplicates(lines)
	}

	if c {
		// Если указан ключ -c, проверяем, отсортированы ли данные
		if !isSorted(lines) {
			fmt.Println("Данные не отсортированы.")
		} else {
			fmt.Println("Данные отсортированы.")
		}
	}

	return lines
}

func reverse(lines []Line) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func removeDuplicates(lines []Line) []Line {
	uniqueLines := make([]Line, 0, len(lines))
	seen := make(map[string]bool)

	for _, line := range lines {
		if !seen[line.Key] {
			uniqueLines = append(uniqueLines, line)
			seen[line.Key] = true
		}
	}

	return uniqueLines
}

func writeLines(lines []Line, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line.Text + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func parseMonthName(monthName string) (time.Time, error) {
	// Парсим название месяца в объект time.Time
	// Предполагаем, что строки имеют формат "MonthName Year"
	// Например, "January 2022"

	monthName = strings.TrimSpace(monthName)
	fmt.Println(monthName)
	date, err := time.Parse("January 2006", monthName)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func parseNumericValue(value string) float64 {
	// Парсим числовое значение с учетом суффиксов
	// Поддерживаемые суффиксы: K, M, G, T (кило, мега, гига, тера)
	// Например, "1.5K" будет преобразовано в 1500.0

	suffixes := map[string]float64{
		"K": 1e3,
		"M": 1e6,
		"G": 1e9,
		"T": 1e12,
	}

	value = strings.TrimSpace(value)
	for suffix, multiplier := range suffixes {
		if strings.HasSuffix(value, suffix) {
			numStr := strings.TrimSuffix(value, suffix)
			num, err := strconv.ParseFloat(numStr, 64)
			if err == nil {
				return num * multiplier
			}
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}

	return num
}

func isSorted(lines []Line) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i].Key < lines[i-1].Key {
			return false
		}
	}

	return true
}
