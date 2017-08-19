package bl

import (
	"log"
	"strconv"
	"strings"
)

type runeRange struct {
	min rune
	max rune
}

func newRuneRange(min, max int) runeRange {
	return runeRange{
		min: rune(min),
		max: rune(max),
	}
}

type supplementRunes struct {
	ranges []runeRange
}

func (sr supplementRunes) getAllRunes() (rr []rune) {
	for _, rng := range sr.ranges {
		for r := rng.min; r < rng.max+1; r++ {
			rr = append(rr, r)
		}
	}
	return
}

func (sr *supplementRunes) appendRange(rng runeRange) {
	sr.ranges = append(sr.ranges, rng)
}

func (sr *supplementRunes) runeWithinRange(r rune) bool {
	for _, rng := range sr.ranges {
		if r >= rng.min && r <= rng.max {
			return true
		}
	}
	return false
}

var configMap map[rune]supplementRunes

func init() {
	configMap = make(map[rune]supplementRunes)

	for i, line := range strings.Split(config, "\n") {
		// Remove space from the line.
		lineTrimmed := strings.TrimSpace(line)

		// Skip empty lines.
		if lineTrimmed == "" {
			continue
		}

		// Extract basic latin character and rules.
		lineSl := strings.Split(lineTrimmed, ":")
		if len(lineSl) != 2 {
			log.Fatalf("configuration line %d: failed to split by column", i)
		}

		var character rune
		for j, cRune := range lineSl[0] {
			if j > 0 {
				log.Fatalf("configuration line %d: multiple configuration characters provided [%s]", i, lineSl[0])
			}
			character = cRune
		}

		if !((character >= 65 && character <= 90) || (character >= 97 && character <= 122)) {
			log.Fatalf("configuration line %d: invalid configuration character provided", i)
		}

		rules := lineSl[1]

		var sr supplementRunes

		// Read rules.
		rulesSl := strings.Split(rules, ",")
		for _, rule := range rulesSl {
			// Check if range.
			if strings.Contains(rule, "-") {
				rangeSl := strings.Split(rule, "-")
				if len(rangeSl) != 2 {
					log.Fatalf("configuration line %d: failed to split range", i)
				}

				min, minErr := strconv.Atoi(rangeSl[0])
				if minErr != nil {
					log.Fatalf("configuration line %d: invalid range min [%s]", i, rangeSl[0])
				}
				max, maxErr := strconv.Atoi(rangeSl[1])
				if maxErr != nil {
					log.Fatalf("configuration line %d: invalid range max [%s]", i, rangeSl[1])
				}

				sr.appendRange(newRuneRange(min, max))
				continue
			}

			// If a single number provided.
			val, valErr := strconv.Atoi(rule)
			if valErr != nil {
				log.Fatalf("configuration line %d: invalid range value [%s]", i, rule)
			}

			// Skip null rules.
			if val > 0 {
				sr.appendRange(newRuneRange(val, val))
			}
		}

		// Assign character rules to the configuration map.
		configMap[character] = sr
	}
}

func getBasicLatinCharacterForRune(r rune) rune {
	// Latin-1 Unicode Supplement letters start at rune 192 representing the 'À'
	// character. Any rune less than that represents either a basic latin
	// character, a symbol or a punctuation which doesn't require conversion.
	if r < 192 {
		return r
	}

	for bl, supplement := range configMap {
		if supplement.runeWithinRange(r) {
			return bl
		}
	}
	return r
}

// GetNonBasicLatinLetters returns all non-basic latin runes representing the
// passed basic latin rune.
func GetNonBasicLatinLetters(blCharacter rune) []rune {
	return configMap[blCharacter].getAllRunes()
}

// ToBasicLatin converts all the characters of the provided text to Basic Latin (Unicode block).
// e.g. Alizée -> Alizee
func ToBasicLatin(text string) (normalized string) {
	for _, c := range text {
		normalized += string(getBasicLatinCharacterForRune(c))
	}
	return
}
