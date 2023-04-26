package main

import "testing"

func TestIncorrent(t *testing.T) {
	const testCase = ""
	res, err := Unpack(testCase)
	if err != nil {
		t.Errorf("Expected error, got %v", res)
	}
}
func TestDigit(t *testing.T) {
	const testCase = "a4bc2d5e"
	const expected = "aaaabccddddde"
	res, err := Unpack(testCase)
	if res != expected {
		t.Errorf("Expected %v, got %v, error is %v", expected, res, err)
	}
}
func TestBase(t *testing.T) {
	const testCase = "abcd"
	const expected = "abcd"
	res, err := Unpack(testCase)
	if res != expected {
		t.Errorf("Expected %v, got %v, error is %v", expected, res, err)
	}
}
func TestEscaped(t *testing.T) {
	const testCase = "qwe\\45"
	const expected = "qwe44444"
	res, err := Unpack(testCase)
	if res != expected {
		t.Errorf("Expected %v, got %v, error is %v", expected, res, err)
	}
}
