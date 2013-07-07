// This package provides methods for turning integer into words and vice-verse
// It is a port of the ruby library rufus-mnemo
package mnemo

import (
	"errors"
	"fmt"
	"strings"
)

var globals = struct {
	Neg      string
	Syl      []string
	Specials map[string]string
}{
	"wi",
	[]string{"ba", "be", "bi", "bo", "bu",
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
		"wa", "wo", "ya", "yo", "yu"},
	map[string]string{
		"hu": "fu",
		"si": "shi",
		"ti": "chi",
		"tu": "tsu",
		"zi": "tzu",
	}}

func toNumber(sym string) (int, error) {
	for ix, testSyl := range globals.Syl {
		if sym == testSyl {
			return ix, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Syllabe \"%s\" not found", sym))
}

func toInt(value string) (int, error) {
	var val int
	var base int
	var err error

	if len(value) == 0 {
		return 0, nil
	}

	if strings.HasPrefix(value, globals.Neg) && len(value) > len(globals.Neg) {
		val, err = toInt(value[len(globals.Neg):])
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

	return len(globals.Syl)*base + val, nil
}

func arrayToSpecial(components []string) []string {
	ret := make([]string, len(components))
	for pos, normal := range components {
		ret[pos] = normal
		for s, d := range globals.Specials {
			if normal == s {
				ret[pos] = d
			}
		}
	}
	return ret
}

func toSpecial(value string) string {
	for s, d := range globals.Specials {
		value = strings.Replace(value, s, d, -1)
	}
	return value
}

func fromSpecial(value string) string {
	for d, s := range globals.Specials {
		value = strings.Replace(value, s, d, -1)
	}
	return value
}

func fromInteger(value int) string {
	if value == 0 {
		return ""
	}
	mod := value % len(globals.Syl)
	rest := value / len(globals.Syl)

	return FromInteger(rest) + globals.Syl[mod]
}

func stringSplit(mnemo string, components []string) []string {
	if len(mnemo) < 1 {
		return components
	}
	components = append(components, mnemo[0:2])

	return stringSplit(mnemo[2:], components)
}

// Split a string into its syllabes
func Split(mnemo string) []string {
	components := make([]string, 0)
	return arrayToSpecial(stringSplit(fromSpecial(mnemo), components))
}

// Indicate whether this string is parsable or not
func IsMnemoWord(mnemo string) bool {
	_, err := ToInteger(mnemo)
	if err != nil {
		return false
	}
	return true
}

// Convert an integer into its mnemo string
func FromInteger(value int) string {
	if value < 0 {
		return fmt.Sprintf("%s%s", globals.Neg, FromInteger(-value))
	}

	return toSpecial(fromInteger(value))
}

// Convert a string into an integer.
// Panics if not parsable
func Must(val int, err error) int {
	if err != nil {
		panic(err)
	}
	return val
}

// Convert a string into an integer
// returns an error if something went wrong
func ToInteger(mnemo string) (int, error) {
	return toInt(fromSpecial(mnemo))
}
