package main

import (
	"fmt"
)

func match(regexp string, text string) bool {
	if len(regexp) > 0 && regexp[0] == '^' {
		return matchhere(regexp[1:], text)
	}
	for i := 0; i < len(text); i++ {
		if matchhere(regexp, text[i:]) {
			return true
		}
	}
	return false
}

func matchhere(regexp string, text string) bool {
	if len(regexp) == 0 {
		return true
	}
	if len(regexp) > 1 && regexp[1] == '*' {
		return matchstar(regexp[0], regexp[2:], text)
	}
	if len(regexp) == 1 && regexp[0] == '$' {
		return len(text) == 0
	}
	if len(regexp) > 0 && len(text) > 0 && (regexp[0] == '.' || regexp[0] == text[0]) {
		return matchhere(regexp[1:], text[1:])
	}
	return false
}

func matchstar(c byte, regexp string, text string) bool {
	for i := 0; i < len(text); i++ {
		if matchhere(regexp, text[i:]) {
			return true
		}
		if i < len(text) && (text[i] != c && c != '.') {
			break
		}
	}
	return false
}

func main() {
	fmt.Println(match("x*", "abcxyz"))
	fmt.Println(match("^xyz", "xyzabc"))
	fmt.Println(match("^xyz", "xyzxyz"))
	fmt.Println(match("^xyz", "xy"))
	fmt.Println(match("^xyz", "pqr"))
	fmt.Println(match("l*", "abcxyz"))
	fmt.Println(match("$xyz", "xyzabc"))
	fmt.Println(match("$xyz", "xyzxyz"))
	fmt.Println(match("xyz", "xyzsdk"))
	fmt.Println(match(".x", "pqr"))
	fmt.Println(match(".x", "pqrx"))
}
