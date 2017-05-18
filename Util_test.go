package zctr

import (
	"testing"
)

func TestIsNumber(t *testing.T) { //{{{
	if !IsNumber("a1A") {
		t.Log("is number Testing OK")
	} else {
		t.Error("is number Testing fail!!")
	}
	if IsNumber("123") {
		t.Log("is number Testing OK")
	} else {
		t.Error("is number Testing fail!!")
	}
}                                   //}}}
func TestIsAlphabet(t *testing.T) { //{{{
	if !IsAlphabet("a1A") {
		t.Log("is Alphabet Testing OK")
	} else {
		t.Error("is Alphabet Testing fail!!")
	}
	if IsAlphabet("AabC") {
		t.Log("is Alphabet Testing OK")
	} else {
		t.Error("is Alphabet Testing fail!!")
	}
}                               //}}}
func TestIsWord(t *testing.T) { //{{{
	if !IsWord("#%^") {
		t.Log("is Word Testing OK")
	} else {
		t.Error("is Word Testing fail!!")
	}
	if IsWord("abc_3s") {
		t.Log("is Word Testing OK")
	} else {
		t.Error("is Word Testing fail!!")
	}
}                                     //}}}
func TestFilterNumber(t *testing.T) { //{{{
	if FilterNumber("a1b2c3#%") == "123" {
		t.Log("filter number Testing OK")
	} else {
		t.Error("filter number Testing fail!!")
	}
}                                       //}}}
func TestFilterAlphabet(t *testing.T) { //{{{
	if FilterAlphabet("0A3a1bC%#_") == "AabC" {
		t.Log("filter Alphabet Testing OK")
	} else {
		t.Error("filter Alphabet Testing fail!!")
	}
}                                   //}}}
func TestFilterWord(t *testing.T) { //{{{
	if FilterWord("<#!abc_3S") == "abc_3S" {
		t.Log("Filter Word Testing OK")
	} else {
		t.Error("Filter Word Testing fail!!")
	}
}                                   //}}}
func TestFilterHtml(t *testing.T) { //{{{
	if FilterHtml("<#!abc_ 3S") == "#!abc_ 3S" {
		t.Log("Filter Html Testing OK")
	} else {
		t.Error("Filter Html Testing fail!!")
	}
}                                   //}}}
func TestHtmlEscape(t *testing.T) { //{{{
	if HtmlEscape("<#!abc_ 3S") == "&lt;#!abc_ 3S" {
		t.Log("Html Escape Testing OK")
	} else {
		t.Error("Html Escape Testing fail!!")
	}
} //}}}
