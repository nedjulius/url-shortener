package main

import "testing"

func Test(t *testing.T) {
	encoding := encode(1231)
	if encoding != "3Vx" {
		t.Errorf("Encoding is incorrect, got: %s, want: %s", encoding, "3Vx")
	}
}
