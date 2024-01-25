package utils

import "testing"

func TestRandomString_Random(t *testing.T) {
	if len(DefaultRandomString.Random(10)) != 10 {
		t.Error("DefaultRandomString Length Error")
	}
	if DefaultRandomString.Random(10) == DefaultRandomString.Random(10) {
		t.Error("DefaultRandomString Value Error")
	}
}
