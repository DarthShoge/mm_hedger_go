package test

import "testing"

func TestBoot(t *testing.T){
	want := "mmbot bootstrapped 🚀"
	got := "mmbot bootstrapped 🚀"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}