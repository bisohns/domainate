package generate

import (
	"math/rand"
	"strings"
	"unicode"
)

var tlds = []string{
	"com",
	"net",
	"me",
	"org",
}

const (
	allowedChars        = "abcdefghijklmnopqrstuvwxyz01234556789_-"
	duplicateVowel bool = false
	removeVowel    bool = true
)

func randBool() bool {
	return rand.Intn(2) == 0
}

// transform : Extend to use a more elaborate method to provide alternate
// domain names
func transform(keyword string) []string {
	return []string{
		keyword,
		keyword + "app",
		"go" + keyword,
		keyword + "site",
		keyword + "time",
		"get" + keyword,
		keyword + "hq",
		"lets" + keyword,
	}
}

// vowelizer : Replaces vowels to give more options of domain names
// e.g biz => beez
func vowelizer(keyword string) (voweled []string) {
	bytekeyword := []byte(keyword)
	for iter := 1; iter <= 3; iter++ {
		if randBool() {
			var vI int = -1
			for i, char := range bytekeyword {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vI = i
					}
				}
			}
			if vI > 0 {
				switch randBool() {
				case duplicateVowel:
					bytekeyword = append(bytekeyword[:vI+1], bytekeyword[vI:]...)
				case removeVowel:
					bytekeyword = append(bytekeyword[:vI], bytekeyword[vI+1:]...)
				}
			}
			keyword = string(bytekeyword)
			voweled = append(voweled, keyword)
		}
	}
	return
}

// Domains : Returns domain names using a given string keyword e.g "one (two) three =>
// [one-two-three.com, ...]
func Domains(keyword string) (domainarray []string) {
	rand.Seed(7)
	// check for weird characters in the keystatement
	var newText []rune
	for _, r := range strings.ToLower(keyword) {
		if unicode.IsSpace(r) {
			r = '-'
		}
		if !strings.ContainsRune(allowedChars, r) {
			continue
		}
		newText = append(newText, r)
	}
	// vowelize to replace vowels randomly
	for _, indText := range vowelizer(string(newText)) {
		//transform to include common domain terms like 'get'
		tempdomainarray := transform(indText)

		for ind := range tempdomainarray {
			tempdomainarray[ind] = tempdomainarray[ind] + "." + tlds[rand.Intn(len(tlds))]
		}
		domainarray = append(domainarray, tempdomainarray...)
	}
	return
}
