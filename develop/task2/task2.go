package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Создаём интерфейс распаковки
type Unpacker interface {
	// Этот интерфейс включает себя только один метод
	Unpack()
}

// Также нам понадобится структура распакованной строки
type PackedString string

// Главная функция принимает строку пользователя
func main() {
	var pkdString PackedString

	fmt.Println("Введите строку для распаковки: ")

	_, err := fmt.Scanf("%s", &pkdString)
	if err != nil {
		log.Println(err)
		// Если строка некорректна, выбрасываем панику
		// согласно условиям задачи
	} else if pkdString.Unpack() == "" {
		panic("Некорректная строка!")
	} else {
		fmt.Println("Распакованная строка: ", pkdString.Unpack())
	}
}

// "a4bc2d5e" => "aaaabccddddde"

func (s PackedString) Unpack() string {
	// Инициализируем несколько переменных:
	// базовый билдер строки,
	// ещё один для обработки тех случаев, когда n > 9
	// временную строку
	// строку, в которую будет записано число n,
	// и флаг, фиксирующий экранирование
	var (
		sb      strings.Builder
		sbA     strings.Builder
		temp    string
		amount  string
		escaped bool
	)
	// Строка "amount" необходима для n > 9: сначала мы объединим два числа в строковом выражении
	// (например, 1+2, составив 12), а затем переведём их в числовой тип
	for i, v := range s {
		// Делаем проверку на экранирование: если оно есть,
		// записываем текущий символ во временную строку
		// и переключаем флаг
		if escaped {
			temp = string(v)
			sb.WriteRune(v)
			escaped = false
			continue
		}

		// Если текущий символ -- число, делаем следующее:
		if _, err := strconv.ParseInt(string(v), 0, 64); err == nil {
			// записываем это число в переменную количества
			amount = string(v)
			// Как только дошли до последнего символа в строке
			// запоминаем количество символов
			if i == len(s)-1 {
				amountN, _ := strconv.ParseInt(string(amount), 0, 64)
				// и проходим циклом по всем символам,
				// собирая строку в переменную temp
				for i := int64(0); i < amountN-1; i++ {
					sb.WriteString(temp)
					continue
				}
				continue
			}
			// Фиксируем количество повторений символа строки
			sbA.WriteString(amount)
			continue
			// И если там что-то записано, переводим строку в число
			// и циклом, равным этому числу, записываем символ в temp
		} else if amount != "" {
			amountN, _ := strconv.ParseInt(string(amount), 0, 64)
			for i := int64(0); i < amountN-1; i++ {
				sb.WriteString(temp)
				continue
			}
			// После чего очищаем amount
			amount = ""
			// И обнуляем счётчик повторений
			sbA.Reset()
		}

		// Если попался символ экранирования, переключаем флаг
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
