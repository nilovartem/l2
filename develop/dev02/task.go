package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func Unpack(s string) (r string, err error) {
	const slash = "\\"
	//Если бы передано число, то выбрасываем ошибку
	if _, err := strconv.Atoi(s); err == nil {
		return r, errors.New("некорректная строка")
	}
	var last rune
	var escaped bool
	var b strings.Builder //самый эффективный способ работы со строками
	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			m := int(char - '0') //'0' = 48, переводим символ в число
			r := strings.Repeat(string(last), m-1)
			b.WriteString(r)
		} else {
			escaped = string(char) == slash && string(last) != slash
			if !escaped {
				b.WriteRune(char)
			}
			last = char
		}
	}
	return b.String(), err
}
