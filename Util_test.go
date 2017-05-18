package zctr

import "testing"

func TestIsNumber(t *testing.T) { //{{{
	if !IsNumber("a1A") {
		t.Log("number Testing OK")
	} else {
		t.Error("number Testing fail!!")
	}
	if IsNumber("123") {
		t.Log("number Testing OK")
	} else {
		t.Error("number Testing fail!!")
	}
}                                   //}}}
func TestIsAlphabet(t *testing.T) { //{{{
	if !IsAlphabet("a1A") {
		t.Log("Alphabet Testing OK")
	} else {
		t.Error("Alphabet Testing fail!!")
	}
	if IsAlphabet("AabC") {
		t.Log("Alphabet Testing OK")
	} else {
		t.Error("Alphabet Testing fail!!")
	}
}                               //}}}
func TestIsWord(t *testing.T) { //{{{
	if !IsWord("#%^") {
		t.Log("number Testing OK")
	} else {
		t.Error("number Testing fail!!")
	}
	if IsWord("abc_3s") {
		t.Log("number Testing OK")
	} else {
		t.Error("number Testing fail!!")
	}
} //}}}
