package main

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	testName string
	slice    []string
	result   map[string][]string
}{
	{"common_test", []string{"пятак", "листок", "тяпка", "пятка", "слиток", "столик"}, map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}},
	{"capitalized_test", []string{"Дурка", "дУрак", "Рудак", "коСти", "исток"}, map[string][]string{"дурка": {"дурак", "рудак"}, "кости": {"исток"}}},
	{"all_unique_test", []string{"go", "is", "a", "good", "compiled", "language"}, nil},
	{"empty_test", []string{}, nil},
}

func TestFindAnagramsRun(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()
			r := FindAnagrams(tt.slice)
			if reflect.DeepEqual(r, tt.result) {
				fmt.Println(r)
				t.Fatal("maps are not equal")
			}
		})
	}
}
