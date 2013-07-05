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

}
