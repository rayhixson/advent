package main

import "testing"

func TestA(t *testing.T) {
	compare(t, "aa", true, false)
	compare(t, "aabb", true, false)
	compare(t, "aabbb", true, true)
	compare(t, "aaaa", false, false)
	compare(t, "", false, false)
	compare(t, "abcdefg", false, false)
	compare(t, "aaxaaxaax", false, true)
	compare(t, "aaxbbbyxbbb", true, false)
	compare(t, "axayaza", false, false)
}

func compare(t *testing.T, s string, shouldHaveTwo, shouldHaveThree bool) {
	twos, threes := countDupes(s)
	if twos != shouldHaveTwo {
		t.Errorf("Failed [%v] ==> %v, %v", s, twos, threes)
	}
}

func TestDiffer(t *testing.T) {
	compareDiffers(t, "abcd", "abxd", "abd", true)
	compareDiffers(t, "abcd", "xbcd", "bcd", true)
	compareDiffers(t, "abcd", "abcx", "abc", true)
}

func compareDiffers(t *testing.T, a, b, match string, shouldBe bool) {
	same, v := differByOne(a, b)
	if same != shouldBe || v != match {
		t.Errorf("nope: %v, %v, %v", a, b, same)
	}
}
