package libwhois

import (
	"golang.org/x/net/idna"
	"log"

	"bufio"
	_ "embed"
	"sort"
	"strings"
)

type PublicSuffixEx struct {
	PunyCode string
	Utf8Name string
}

//go:embed public_suffix_list.dat
var publicSuffixListContent string

func GetAllPublicSuffixes() []string {
	result := make([]string, 0, 1024*2)

	scanner := bufio.NewScanner(strings.NewReader(publicSuffixListContent))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "//") {
			continue
		}

		result = append(result, strings.TrimSpace(line))
	}

	sort.Strings(result)

	return result
}

func GetAllTopLevelPublicSuffixes() []string {
	topLevelSuffixes := GetAllTopLevelPublicSuffixesEx()

	result := make([]string, 0, len(topLevelSuffixes))
	for _, s := range topLevelSuffixes {
		result = append(result, s.PunyCode)
	}

	sort.Strings(result)

	return result
}

func GetAllTopLevelPublicSuffixesEx() []PublicSuffixEx {
	allSuffixes := GetAllPublicSuffixes()

	result := make([]PublicSuffixEx, 0, len(allSuffixes))
	for _, suffix := range allSuffixes {
		if strings.Contains(suffix, ".") {
			continue
		}

		res := PublicSuffixEx{
			PunyCode: suffix,
		}

		if !isLatinLettersDigitsDash(suffix) {
			asciiDomain, err := idna.ToASCII(suffix)
			if err != nil {
				log.Fatalf("Failed to convert to ASCII: %v", err)
			}
			res.PunyCode = asciiDomain
			res.Utf8Name = suffix
		}

		result = append(result, res)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].PunyCode < result[j].PunyCode
	})

	return result
}

func isLatinLettersDigitsDash(s string) bool {
	for _, r := range s {
		isLatinSymbol := r >= 'a' && r <= 'z'
		isDigit := r >= '0' && r <= '9'

		if !isLatinSymbol && !isDigit && r != '-' {
			return false
		}
	}
	return true
}
