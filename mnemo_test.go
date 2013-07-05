package mnemo

import (
	"fmt"
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
	fmt.Printf("takeshi = %d\n", ToInteger("takeshi"))
	fmt.Printf("%d = %s\n", ToInteger("takeshi"), FromInteger(ToInteger("takeshi")))
}

func TestIsMnemoWord(t *testing.T) {
	if !IsMnemoWord("takeshi") {
		t.Fatalf("%s is not a word and should be", "takeshi")

	}

}
