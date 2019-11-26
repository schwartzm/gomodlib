package gomodlib

import "rsc.io/quote"

func Happy() string {
	return "What a great day."
}

func Sad() string {
	return "My cat ate a Lego brick."
}

func Joyous() string {
	return "For Christmas, I got the Metal Beards Sea Cow Lego set."
}

func Mad() string {
	return "I am infuriated."
}

func Hello() string {
	return quote.Hello()
}
