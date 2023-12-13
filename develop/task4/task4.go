/* Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого --
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.


*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

func checkAnagram(words *[]string) map[string]*[]string {
	anagrams := make(map[string][]string)
	res := make(map[string]*[]string)

	for _, word := range *words {
		word = strings.ToLower(word)

		sortedW := sortString(word)

		anagrams[sortedW] = append(anagrams[sortedW], word)
	}

	for key, value := range anagrams {
		if len(value) <= 1 {
			delete(anagrams, key)
		}
	}

	sortMap(&anagrams)

	for _, k := range anagrams {
		key := k[0]
		res[key] = new([]string)
		*res[key] = append(*res[key], k...)
	}

	return res
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func sortMap(anagrams *map[string][]string) {
	for _, value := range *anagrams {
		sort.Strings(value)
	}
}

func main() {
	str := &[]string{"пятак", "пятка", "тяпка", "кот", "кто", "бор",
		"рука", "кура", "раку", "верфь", "умник"}

	// Для убоства чтения конвертируем в json
	res, err := json.Marshal(checkAnagram((str)))
	if err != nil {
		log.Fatal("Ошибка при переводе в json")
		return
	}
	fmt.Println(string(res))
}
