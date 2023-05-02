package gogrep

import (
	"strconv"
	"strings"
)

type Flags struct {
	After      uint
	Before     uint
	Context    uint
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func markCorrectLines(lines []int, up, down uint) []bool {
	result := make([]bool, len(lines))
	for i, v := range lines {
		if v != -1 {
			left := i - int(up)
			right := i + int(down)
			if left < 0 {
				left = 0
			}
			if right > len(lines)-1 {
				right = len(lines) - 1
			}
			for j := left; j <= right; j++ {
				result[j] = true
			}
		}
	}
	return result
}

func findMatches(strs []string, pattern string, ignoreCase bool) []int {
	// -1 - не матчится. Остальное - матчится
	result := make([]int, len(strs))
	for i, v := range strs {
		if ignoreCase {
			result[i] = strings.Index(strings.ToLower(v), pattern)
			continue
		}
		result[i] = strings.Index(v, pattern)
	}
	return result
}

func Grep(strs []string, pattern string, flags Flags) string {
	if flags.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	indexes := findMatches(strs, pattern, flags.IgnoreCase)

	if flags.Fixed {
		for i := 0; i < len(indexes); i++ {
			idx := indexes[i]
			if idx != -1 {
				if strs[i] != pattern {
					indexes[i] = -1
				}
			}
		}
	}

	if flags.Invert {
		for i, v := range indexes {
			if v == -1 {
				indexes[i] = 0
			} else {
				indexes[i] = -1
			}
		}
	}

	if flags.Count {
		count := 0
		for _, v := range indexes {
			if v != -1 {
				count++
			}
		}
		return strconv.Itoa(count) + "\n"
	}

	if flags.Context > 0 {
		flags.Before, flags.After = flags.Context, flags.Context
	}

	finalLines := markCorrectLines(indexes, flags.Before, flags.After)
	var result string
	for i, v := range finalLines {
		if v {
			if flags.LineNum {
				result += strconv.Itoa(i+1) + ":" + strs[i]
			} else {
				result += strs[i]
			}
		}
	}

	return result
}
