package cmd

import (
	"sort"
	"strconv"
	"strings"
)

func makeUnique(s []string) []string {
	allKeys := make(map[string]bool)
	slice := []string{}
	for _, item := range s {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			slice = append(slice, item)
		}
	}
	return slice
}
func Sort(strs []string, key int, numeric, reverse, unique bool) []string {
	if unique {
		strs = makeUnique(strs)
	}

	sort.Slice(strs, func(i, j int) bool {
		if strings.TrimSpace(strs[i]) == "" {
			return false
		}
		iColumn := strings.Split(strs[i], " ")
		jColumn := strings.Split(strs[j], " ")
		if len(iColumn) <= key || len(jColumn) <= key {
			return true
		}
		switch numeric {
		case true:
			iNum, ierr := strconv.Atoi(iColumn[key])
			jNum, jerr := strconv.Atoi(jColumn[key])
			if ierr == nil && jerr == nil {
				return iNum < jNum
			} else if ierr != nil && jerr != nil {
				return iColumn[key] < jColumn[key]
			} else if ierr != nil {
				return true
			} else if jerr != nil {
				return false
			}
		case false:
			return iColumn[key] < jColumn[key]
		}

		return false
	})

	if reverse {
		for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
			strs[i], strs[j] = strs[j], strs[i]
		}
	}
	return strs
}
