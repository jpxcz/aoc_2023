package main

import "testing"

func TestBackLash(t *testing.T) {
	if !hasSymbol("asdf@") {
		t.Fatal("expected backlash")
	}
}
