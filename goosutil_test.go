package goosutil

import "testing"

func TestExists(t *testing.T) {
	b := Exists("Makefile")
	if b == false {
		t.Error("file doesnt exist.")
	}

	b = Exists("hogefugatemp")
	if b == true {
		t.Error("illegal file check.")
	}
}
