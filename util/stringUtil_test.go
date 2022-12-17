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
