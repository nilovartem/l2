package sort

import (
	"sort"
	"strconv"
	"strings"
)

type Comparator interface {
	int | string | rune
}

func compare[comparator Comparator](i, j comparator) (result bool) {
	if i < j {
		result = true
	}
	return
}
func unify(lines []string) (result []string) {
	m := make(map[string]bool)
	for _, line := range lines {
		m[line] = true
	}
	for key := range m {
		result = append(result, key)
	}
	return
}
func Sort(lines []string, key int, numeric, reverse, unique bool) []string {
	if unique {
		lines = unify(lines)
	}
	//Обеспечим валидное значение key
	if key <= 0 {
		key = 1
	}
	sort.Slice(lines, func(i, j int) bool {
		line_i := strings.Split(lines[i], " ") //слайс из слов в строке [i]
		line_j := strings.Split(lines[j], " ") //слайс из слов в строке [j]
		column_i := line_i[key-1]
		column_j := line_j[key-1]
		result := compare(line_i[key-1], line_j[key-1])
		if numeric {
			/*если нужно сравнивать числа,
			то конвертируем колонку в число и сравниваем*/
			number_i, err_i := strconv.Atoi(column_i)
			number_j, err_j := strconv.Atoi(column_j)
			if err_i == nil && err_j == nil {
				result = compare(number_i, number_j)
			}
		}
		return result
	})
	if reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	return lines
}
