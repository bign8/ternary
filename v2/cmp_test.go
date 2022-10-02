package ternary

import "testing"

const (
	T = `true`
	F = `false`
)

func TestTruthy(t *testing.T) {
	if got := Cmp(true, T, F); got != T {
		t.Errorf(`want(%q) != got(%q)`, T, got)
	}
}

func TestFalsy(t *testing.T) {
	if got := Cmp(false, T, F); got != F {
		t.Errorf(`want(%q) != got(%q)`, F, got)
	}
}
