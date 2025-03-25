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

func TestLazyTruthy(t *testing.T) {
	if got := CmpLazy(true, func () string {
		return T
	}, func () string {
		t.Error(`lazilyEvaluate failed`)
		return F
	); got != T {
		t.Errorf(`want(%q) != got(%q)`, T, got)
	}
}

func TestLazyFalsy(t *testing.T) {
	if got := CmpLazy(false, func () string {
		t.Error(`lazilyEvaluate failed`)
		return T
	}, func () string {
		return F
	); got != F {
		t.Errorf(`want(%q) != got(%q)`, F, got)
	}
}
