// mnemo.go
package mnemo

import (
	"errors"
	"fmt"
	"strings"
)

const NEG string = "wi"

func SYL() []string {
	return []string{"ba", "be", "bi", "bo", "bu",
		"da", "de", "di", "do", "du",
		"ga", "ge", "gi", "go", "gu",
		"ha", "he", "hi", "ho", "hu",
		"ja", "je", "ji", "jo", "ju",
		"ka", "ke", "ki", "ko", "ku",
		"ma", "me", "mi", "mo", "mu",
		"na", "ne", "ni", "no", "nu",
		"pa", "pe", "pi", "po", "pu",
		"ra", "re", "ri", "ro", "ru",
		"sa", "se", "si", "so", "su",
		"ta", "te", "ti", "to", "tu",
		"za", "ze", "zi", "zo", "zu",
		"wa", "wo", "ya", "yo", "yu"}
}

func SPECIALS() map[string]string {
	return map[string]string{
		"hu": "fu",
		"si": "shi",
		"ti": "chi",
		"tu": "tsu",
		"zi": "tzu"
	}
}

func toNumber(syl string) (int, error) {
	for ix, testSyl := range SYL() {
		if syl == testSyl {
			return ix, nil
		}
	}
	return -1, errors.New(fmt.Sprint("Syllabe \"%s\" not found", syl))
}

func toInt(value string) int {
	if len(value) == 0 {
		return 0
	}
	if strings.HasPrefix(value, NEG) {
		return -1 * toInt(value[len(NEG):])
	}
	n, _ := toNumber(value[len(value)-2 : 2])
	return len(SYL())*toInt(value[0:len(value)-3]) + n
}

func toSpecial(value string) string {

}

func FromInteger(value int) string {
	if value < 0 {
		return fmt.Sprint("%s%s", NEG, FromInteger(-value))
	}

	return toSpecial(fromInteger(value))
}

func main() {
	fmt.Println("Hello World!")
}
