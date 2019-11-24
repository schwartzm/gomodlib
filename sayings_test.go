package gomodlib

import "testing"

func TestHappy(t *testing.T) {
	want := "What a great day."
	if got := Happy(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestSad(t *testing.T) {
	want := "My pet cat is very sick."
	if got := Sad(); got != want {
		t.Errorf("Sad() = %q, want %q", got, want)
	}
}

func TestJoyous(t *testing.T) {
	want := "We won the playoff game."
	if got := Joyous(); got != want {
		t.Errorf("Joyous() = %q, want %q", got, want)
	}
}

func TestMad(t *testing.T) {
	want := "I am infuriated."
	if got := Mad(); got != want {
		t.Errorf("Mad() = %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
