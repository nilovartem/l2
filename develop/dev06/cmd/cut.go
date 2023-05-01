package cmd

import "strings"

func Cut(str, delimiter string, fields uint, separated bool) string {
	// убираем лишние пробелы, если в качестве разделителя выбран пробел
	if delimiter == " " {
		str = strings.Join(strings.Fields(str), delimiter)
	}
	if separated {
		if !strings.Contains(str, delimiter) {
			return ""
		}
	}
	split := strings.Split(str, delimiter)
	if uint(len(split)) > fields && split[fields] != "" {
		return split[fields]
	}
	return ""

}
