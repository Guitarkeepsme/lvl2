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
	"log"
	"os"
	"strings"
)

var (
	fields               int
	delimiter, separated bool
)

func main() {
	var (
		text string
		err  error
	)
	setFlags()

	fmt.Println(fields)
	fmt.Println(separated)

	for text != "\n" {
		text, err = readStdin()
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case delimiter:
		case separated:
		default:
			if fields > 0 {
				fmt.Printf("%s\n", cut(text, fields))
				os.Exit(1)
			} else if fields < 0 {
				fmt.Printf("cut: полей должно быть больше одного, но не %d\n", fields)
				os.Exit(1)
			}
			fmt.Printf("cut: необходим список байтов, символов или полей\n")
			os.Exit(1)

		}
	}
}

// Инициализируем необходимые флаги
func setFlags() {
	flag.IntVar(&fields, "f", 0, "установить номер поля или колонки")
	flag.BoolVar(&delimiter, "d", false, "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "выбрать только строки с разделителем")
	flag.Parse()
}

// Функция для чтения из стдин
func readStdin() (string, error) {
	// текст пользователя собирается билдером
	var text strings.Builder
	// ридер читает данные
	reader := bufio.NewReader(os.Stdin)
	tempText, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text.WriteString(tempText)

	return text.String(), nil

}

// Функция для дробления текста на фрагменты за счёт разделителя
func cut(text string, fields int) []string {
	// Делим текст по табуляциям
	strs := strings.Split(text, "\t")
	// Результатом будет слайс строк
	res := make([]string, 1)
	// Если отрезков больше одного,
	// то делаем проверку: в случае если отрезок
	// содержит \n, к результату аппендим
	if len(strs) > 1 {
		for _, v := range strs {
			if strings.Contains(v, "\n") {
				res = append(res, strs[fields])
			}
		}
	}
	return res
}
