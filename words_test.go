package gomodlib

import "testing"

func TestShortWord(t *testing.T) {
	want := "hi"
	if got := ShortWord(); got != want {
		t.Errorf("ShortWord() = %q; want %q", got, want)
	}
}

func TestMediumWord(t *testing.T) {
	want := "goodbye"
	if got := MediumWord(); got != want {
		t.Errorf("MediumWord = %q, want %q", got, want)
	}
}

func TestLongWord(t *testing.T) {
	want := "Mississippi"
	if got := LongWord(); got != want {
		t.Errorf("LongWord = %q, want = %q", got, want)
	}
}

func TestVeryLongWord(t *testing.T) {
	want := "supercalifragilisticexpialidocious"
	if got := VeryLongWord(); got != want {
		t.Errorf("VeryLongWord = %q, want = %q", got, want)
	}
}

func TestIsShort(t *testing.T) {
	want := true
	if got := IsShort("tool"); got != want {
		t.Errorf("IsShort = %t, want = %t", got, want)
	}
}
