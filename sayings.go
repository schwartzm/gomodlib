package gomodlib

import "rsc.io/quote"

func Happy() string {
	return "What a great day."
}

func Sad() string {
	return "My pet cat is very sick."
}

func Joyous() string {
	return "We won the playoff game."
}

func Mad() string {
	return "I am infuriated."
}

func Hello() string {
	return quote.Hello()
}
