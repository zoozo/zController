package zctr

import (
	"html"
	"regexp"
)

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
func FilterNumber(data string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(data, "")
}
func FilterAlphabet(data string) string {
	re := regexp.MustCompile("[^a-zA-Z]")
	return re.ReplaceAllString(data, "")
}
func FilterWord(data string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9_]")
	return re.ReplaceAllString(data, "")
}
func FilterHtml(data string) string {
	re := regexp.MustCompile(`[<>'"&]`)
	return re.ReplaceAllString(data, "")
}
func HtmlEscape(data string) string {
	return html.EscapeString(data)
}
