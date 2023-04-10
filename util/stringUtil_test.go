package util

import "testing"

func TestMatch(t *testing.T) {
	var result = false
	result = Match("arpc", "", "")
	if !result {
		t.Error("empty regex error")
	}
	result = Match("arpc", "arp", "")
	if !result {
		t.Error("correct regex error")
	}
	result = Match("arpc", "arc", "")
	if result {
		t.Error("incorrect regex error")
	}
	result = Match("arpc", "arp", "ap")
	if !result {
		t.Error("incorrect invertRegex error")
	}
	result = Match("arpc", "arp", "ar")
	if result {
		t.Error("correct invertRegex error")
	}
}

func TestRightCut(t *testing.T) {
	var result = false
	result = "ABCDEFG" == RightCut("ABCDEFG", 8)
	if !result {
		t.Error("ABCDEFG,8 error")
	}
	result = "ABCDEFG" == RightCut("ABCDEFG", 7)
	if !result {
		t.Error("ABCDEFG,7 error")
	}
	result = "ABCD.." == RightCut("ABCDEFG", 6)
	if !result {
		t.Error("ABCDEFG,6 error")
	}
	result = "" == RightCut("ABCDEFG", 0)
	if !result {
		t.Error("ABCDEFG,0 error")
	}
	result = "" == RightCut("ABCDEFG", -1)
	if !result {
		t.Error("ABCDEFG,0 error")
	}
}

func TestLeftCut(t *testing.T) {
	var result = false
	result = "ABCDEFG" == LeftCut("ABCDEFG", 8)
	if !result {
		t.Error("ABCDEFG,8 error")
	}
	result = "ABCDEFG" == LeftCut("ABCDEFG", 7)
	if !result {
		t.Error("ABCDEFG,7 error")
	}
	result = "..DEFG" == LeftCut("ABCDEFG", 6)
	if !result {
		t.Error("ABCDEFG,6 error")
	}
	result = "" == LeftCut("ABCDEFG", 0)
	if !result {
		t.Error("ABCDEFG,0 error")
	}
	result = "" == LeftCut("ABCDEFG", -1)
	if !result {
		t.Error("ABCDEFG,0 error")
	}
}
