package zctr

import (
	"crypto/sha256"
	"fmt"
	"html"
	"reflect"
	"regexp"
)

func InSlice(value interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if value == s.Index(i).Interface() {
			return true
		}
	}
	return false
}

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
func FilterNumber2(data string) string {
	re := regexp.MustCompile("[^0-9.]")
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
func FilterSQL(data string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9_+-/*//()]")
	return re.ReplaceAllString(data, "")
}
func FilterHtml(data string) string {
	re := regexp.MustCompile(`[<>'"&]`)
	return re.ReplaceAllString(data, "")
}
func HtmlEscape(data string) string {
	return html.EscapeString(data)
}
func Sha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))

}
