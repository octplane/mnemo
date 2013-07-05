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
		"zi": "tzu",
	}
}

func toNumber(syl string) (int, error) {
	for ix, testSyl := range SYL() {
		if syl == testSyl {
			return ix, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Syllabe \"%s\" not found", syl))
}

func toInt(value string) (int, error) {
	var val int
	var base int
	var err error

	if len(value) == 0 {
		return 0, nil
	}

	if strings.HasPrefix(value, NEG) {
		val, err = toInt(value[len(NEG):])
		if err != nil {
			return -1, err
		}
		return -1 * val, nil
	}
	val, err = toNumber(value[len(value)-2 : len(value)])
	if err != nil {
		return -1, err
	}
	base, err = toInt(value[:len(value)-2])
	if err != nil {
		return -1, err
	}

	return len(SYL())*base + val, nil
}

func toSpecial(value string) string {
	for s, d := range SPECIALS() {
		value = strings.Replace(value, s, d, -1)
	}
	return value
}

func fromSpecial(value string) string {
	for d, s := range SPECIALS() {
		value = strings.Replace(value, s, d, -1)
	}
	return value
}

func fromInteger(value int) string {
	if value == 0 {
		return ""
	}
	mod := value % len(SYL())
	rest := value / len(SYL())

	return FromInteger(rest) + SYL()[mod]
}

func IsMnemoWord(value string) bool {
	_, err := ToInteger(value)
	if err != nil {
		return false
	}
	return true
}

func FromInteger(value int) string {
	if value < 0 {
		return fmt.Sprintf("%s%s", NEG, FromInteger(-value))
	}

	return toSpecial(fromInteger(value))
}

func Must(val int, err error) int {
	if err != nil {
		panic(err)
	}
	return val
}

func ToInteger(mnemo string) (int, error) {
	return toInt(fromSpecial(mnemo))
}

func main() {
	fmt.Println(FromInteger(12))
}
