package mnemo

import (
	"testing"
)

func TestSpecial(t *testing.T) {
	if toSpecial("hu") != "fu" {
		t.Fatalf("%s should be special of %s, got %s", "fu", "hu", toSpecial("hu"))
	}

	if fromSpecial("fu") != "hu" {
		t.Fatalf("%s should be from special of %s, got %s", "hu", "fu", fromSpecial("fu"))
	}
}

func TestBackAndForth(t *testing.T) {
	if FromInteger(Must(ToInteger("takeshi"))) != "takeshi" {
		t.Fatalf("'takeshi' != %s", FromInteger(Must(ToInteger("takeshi"))))
	}

}

func TestIsMnemoWord(t *testing.T) {
	if !IsMnemoWord("takeshi") {
		t.Fatalf("%s is not a word and should be", "takeshi")
	}
	if !IsMnemoWord("tsunasima") {
		t.Fatalf("%s is not a word and should be", "tsunasima")
	}
	if !IsMnemoWord("tunashima") {
		t.Fatalf("%s is not a word and should be", "tunashima")
	}
	if IsMnemoWord("dsfadf") {
		t.Fatalf("%s is a word and should no be", "dsfadf")
	}
	if IsMnemoWord("takeshin") {
		t.Fatalf("%s is a word and should no be", "takeshin")
	}
}

func TestSplit(t *testing.T) {
	test1 := []string{"ko", "chi", "pi", "ga"}
	test2 := []string{"ko", "na", "de", "tzu"}

	kochi := Split("kochipiga")
	for p, item := range kochi {
		if item != test1[p] {
			t.Fatalf("kochipiga is not split correctly: %v", kochi)
		}
	}
	kona := Split("konadetzu")
	for p, item := range kona {
		if item != test2[p] {
			t.Fatalf("konadetzy is not split correctly: %v", kona)
		}
	}
}

func TestZero(t *testing.T) {
	if FromInteger(0) != "" {
		t.Fatalf("FromInteger(0) != '' (== %s)", FromInteger(0))
	}
	if Must(ToInteger("")) != 0 {
		t.Fatalf("ToInteger(\"\") != 0 (== %d)", Must(ToInteger("")))
	}
}

func backAndForth(t *testing.T, mnemo string, value int) {
	if FromInteger(value) != mnemo {
		t.Fatalf("FromInteger(%d) != \"%s\" (got %s)", value, mnemo, FromInteger(value))
	}
}

func TestNegative(t *testing.T) {
	backAndForth(t, "wina", -35)
	backAndForth(t, "wibe", -1)
}

func TestWiBadPosition(t *testing.T) {
	testValues := []string{"wi", "wiwi", "bewi", "nawi", "nabewi", "nawibe", "nawiwi"}

	for _, testValue := range testValues {
		val, err := ToInteger(testValue)
		if err == nil {
			t.Fatalf("%s was parsed as %d, should have failed.", testValue, val)
		}
		t.Log(err)
	}
}
