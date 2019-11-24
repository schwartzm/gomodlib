package gomodlib

func ShortWord() string {
	return "hi"
}

func MediumWord() string {
	return "goodbye"
}

func LongWord() string {
	return "Mississippi"
}

func VeryLongWord() string {
	return "supercalifragilisticexpialidocious"
}

func IsShort(w string) bool {
	if len(w) <= 5 {
		return true
	}
	return false
}
