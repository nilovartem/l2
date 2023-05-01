package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

}

// FindAnagrams searches anagrams in a slice.
func FindAnagrams(slice []string) map[string][]string {
	result := make(map[string][]string)
	var keys []string
f1:
	for _, v := range slice {
		word := strings.ToLower(v)
		for _, key := range keys {
			if makeRaw(word) == makeRaw(key) {
				result[key] = sortedAppend(result[key], strings.ToLower(word))
				continue f1
			}
		}
		keys = append(keys, v)
	}

	return result
}
func makeRaw(val string) string {
	letters := strings.Split(val, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func sortedAppend(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}
