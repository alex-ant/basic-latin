package bl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToBasicLatin(t *testing.T) {
	require.Equal(t, "AAAAAAAAAAAAAA", ToBasicLatin("ÀÁÂÃÄÅĀĂĄǞǠǺȀȂ"),
		"Failed to convert string to basic latin")
	require.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", ToBasicLatin("ÀƂĈĐÊƑĢĦĨĴĶĽMŇØƤQŘŞŦŨƲŴXŶƵ"),
		"Failed to convert string to basic latin")
	require.Equal(t, "abcdefghijklmnopqrstuvwxyz", ToBasicLatin("áƃćđèƒġĥĩǰƙŀmņòƥqȓșţùvŵxýź"),
		"Failed to convert string to basic latin")
}

func TestGetNonBasicCapitalLatinLetters(t *testing.T) {
	errorMsg := "Failed to obtain non-basic latin runes for character"

	require.Equal(t, "ÀÁÂÃÄÅĀĂĄǞǠǺȀȂ", string(GetNonBasicLatinLetters('A')),
		fmt.Sprintf("%s A", errorMsg))
	require.Equal(t, "ƁƂ", string(GetNonBasicLatinLetters('B')),
		fmt.Sprintf("%s B", errorMsg))
	require.Equal(t, "ÇĆĈĊČƇ", string(GetNonBasicLatinLetters('C')),
		fmt.Sprintf("%s C", errorMsg))
	require.Equal(t, "ĎĐƊƋ", string(GetNonBasicLatinLetters('D')),
		fmt.Sprintf("%s D", errorMsg))
	require.Equal(t, "ÈÉÊËĒĔĖĘĚȄȆ", string(GetNonBasicLatinLetters('E')),
		fmt.Sprintf("%s E", errorMsg))
	require.Equal(t, "Ƒ", string(GetNonBasicLatinLetters('F')),
		fmt.Sprintf("%s F", errorMsg))
	require.Equal(t, "ĜĞĠĢƓǤǦǴ", string(GetNonBasicLatinLetters('G')),
		fmt.Sprintf("%s G", errorMsg))
	require.Equal(t, "ĤĦ", string(GetNonBasicLatinLetters('H')),
		fmt.Sprintf("%s H", errorMsg))
	require.Equal(t, "ÌÍÎÏĨĪĬĮİƗȈȊ", string(GetNonBasicLatinLetters('I')),
		fmt.Sprintf("%s I", errorMsg))
	require.Equal(t, "Ĵ", string(GetNonBasicLatinLetters('J')),
		fmt.Sprintf("%s J", errorMsg))
	require.Equal(t, "ĶƘǨ", string(GetNonBasicLatinLetters('K')),
		fmt.Sprintf("%s K", errorMsg))
	require.Equal(t, "ĹĻĽĿŁ", string(GetNonBasicLatinLetters('L')),
		fmt.Sprintf("%s L", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('M')),
		fmt.Sprintf("%s M", errorMsg))
	require.Equal(t, "ÑŃŅŇƝǸ", string(GetNonBasicLatinLetters('N')),
		fmt.Sprintf("%s N", errorMsg))
	require.Equal(t, "ÒÓÔÕÖØŌŎŐƟƠǪǬǾȌȎ", string(GetNonBasicLatinLetters('O')),
		fmt.Sprintf("%s O", errorMsg))
	require.Equal(t, "Ƥ", string(GetNonBasicLatinLetters('P')),
		fmt.Sprintf("%s P", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('Q')),
		fmt.Sprintf("%s Q", errorMsg))
	require.Equal(t, "ŔŖŘȐȒ", string(GetNonBasicLatinLetters('R')),
		fmt.Sprintf("%s R", errorMsg))
	require.Equal(t, "ŚŜŞŠȘ", string(GetNonBasicLatinLetters('S')),
		fmt.Sprintf("%s S", errorMsg))
	require.Equal(t, "ŢŤŦƬƮȚ", string(GetNonBasicLatinLetters('T')),
		fmt.Sprintf("%s T", errorMsg))
	require.Equal(t, "ÙÚÛÜŨŪŬŮŰŲƯȔȖ", string(GetNonBasicLatinLetters('U')),
		fmt.Sprintf("%s U", errorMsg))
	require.Equal(t, "Ʋ", string(GetNonBasicLatinLetters('V')),
		fmt.Sprintf("%s V", errorMsg))
	require.Equal(t, "Ŵ", string(GetNonBasicLatinLetters('W')),
		fmt.Sprintf("%s W", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('X')),
		fmt.Sprintf("%s X", errorMsg))
	require.Equal(t, "ÝŶŸƳ", string(GetNonBasicLatinLetters('Y')),
		fmt.Sprintf("%s Y", errorMsg))
	require.Equal(t, "ŹŻŽƵ", string(GetNonBasicLatinLetters('Z')),
		fmt.Sprintf("%s Z", errorMsg))
}

func TestGetNonBasicSmallLatinLetters(t *testing.T) {
	errorMsg := "Failed to obtain non-basic latin runes for character"

	require.Equal(t, "àáâãäåāăąǟǡǻȁȃ", string(GetNonBasicLatinLetters('a')),
		fmt.Sprintf("%s a", errorMsg))
	require.Equal(t, "ƀƃ", string(GetNonBasicLatinLetters('b')),
		fmt.Sprintf("%s b", errorMsg))
	require.Equal(t, "çćĉċčƈ", string(GetNonBasicLatinLetters('c')),
		fmt.Sprintf("%s c", errorMsg))
	require.Equal(t, "ďđƌ", string(GetNonBasicLatinLetters('d')),
		fmt.Sprintf("%s d", errorMsg))
	require.Equal(t, "èéêëēĕėęěȅȇ", string(GetNonBasicLatinLetters('e')),
		fmt.Sprintf("%s e", errorMsg))
	require.Equal(t, "ƒ", string(GetNonBasicLatinLetters('f')),
		fmt.Sprintf("%s f", errorMsg))
	require.Equal(t, "ĝğġģǥǧǵ", string(GetNonBasicLatinLetters('g')),
		fmt.Sprintf("%s g", errorMsg))
	require.Equal(t, "ĥħ", string(GetNonBasicLatinLetters('h')),
		fmt.Sprintf("%s h", errorMsg))
	require.Equal(t, "ìíîïĩīĭįıȉȋ", string(GetNonBasicLatinLetters('i')),
		fmt.Sprintf("%s i", errorMsg))
	require.Equal(t, "ĵǰ", string(GetNonBasicLatinLetters('j')),
		fmt.Sprintf("%s j", errorMsg))
	require.Equal(t, "ķƙǩ", string(GetNonBasicLatinLetters('k')),
		fmt.Sprintf("%s k", errorMsg))
	require.Equal(t, "ĺļľŀłƚ", string(GetNonBasicLatinLetters('l')),
		fmt.Sprintf("%s l", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('m')),
		fmt.Sprintf("%s m", errorMsg))
	require.Equal(t, "ñńņňŉƞǹ", string(GetNonBasicLatinLetters('n')),
		fmt.Sprintf("%s n", errorMsg))
	require.Equal(t, "òóôõöøōŏőơǫǭǿȍȏ", string(GetNonBasicLatinLetters('o')),
		fmt.Sprintf("%s o", errorMsg))
	require.Equal(t, "ƥ", string(GetNonBasicLatinLetters('p')),
		fmt.Sprintf("%s p", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('q')),
		fmt.Sprintf("%s q", errorMsg))
	require.Equal(t, "ŕŗřȑȓ", string(GetNonBasicLatinLetters('r')),
		fmt.Sprintf("%s r", errorMsg))
	require.Equal(t, "śŝşšș", string(GetNonBasicLatinLetters('s')),
		fmt.Sprintf("%s s", errorMsg))
	require.Equal(t, "ţťŧƫƭț", string(GetNonBasicLatinLetters('t')),
		fmt.Sprintf("%s t", errorMsg))
	require.Equal(t, "ùúûüũūŭůűųưȕȗ", string(GetNonBasicLatinLetters('u')),
		fmt.Sprintf("%s u", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('v')),
		fmt.Sprintf("%s v", errorMsg))
	require.Equal(t, "ŵ", string(GetNonBasicLatinLetters('w')),
		fmt.Sprintf("%s w", errorMsg))
	require.Equal(t, "", string(GetNonBasicLatinLetters('x')),
		fmt.Sprintf("%s x", errorMsg))
	require.Equal(t, "ýÿŷƴ", string(GetNonBasicLatinLetters('y')),
		fmt.Sprintf("%s y", errorMsg))
	require.Equal(t, "źżžƶ", string(GetNonBasicLatinLetters('z')),
		fmt.Sprintf("%s z", errorMsg))
}
