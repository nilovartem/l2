package main

import "testing"

const (
	INVALID_SERVER = ""
	VALID_SERVER   = "0.beevik-ntp.pool.ntp.org"
)

func TestInvalidServer(t *testing.T) {
	_, err := Time(INVALID_SERVER)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestValidServer(t *testing.T) {
	_, err := Time(VALID_SERVER)
	if err != nil {
		t.Errorf("Expected time, got error")
	}
}
