package zctr

import "regexp"

func IsNumber(data string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", data); m {
		return true
	}
	return false
}
func IsAlphabet(data string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", data); m {
		return true
	}
	return false
}
func IsWord(data string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", data); m {
		return true
	}
	return false
}
