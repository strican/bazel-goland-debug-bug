package main

import "testing"

func TestFoo(t *testing.T) {
	if want, got := 3, Foo(1, 2); want != got {
		t.Fatalf("\nwant: %d\ngot : %d", want, got)
	}
}
