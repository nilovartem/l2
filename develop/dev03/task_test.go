package main

import (
	"testing"

	"github.com/nilovartem/l2/develop/dev03/sort"
)

var test_case_numbers_source = []string{"5", "1", "2", "3", "4"}
var test_case_numbers_valid = []string{"1", "2", "3", "4", "5"}

func TestKeyNumbers(t *testing.T) {
	result := sort.Sort(test_case_numbers_source, 0, true, false, false)
	if !SliceCompare(result, test_case_numbers_valid) {
		t.Error("got different slices")
	}
}
func SliceCompare(first, second []string) (result bool) {
	result = true
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			result = false
		}
	}
	return
}
