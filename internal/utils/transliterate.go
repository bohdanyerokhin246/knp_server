package utils

import (
	"strings"
	"unicode"
)

func Transliterate(input string) string {
	if input == "" {
		return "123"
	}

	translitMap := map[rune]string{
		'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'є': "ye", 'ж': "zh",
		'з': "z", 'и': "y", 'і': "i", 'ї': "yi", 'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
		'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts",
		'ч': "ch", 'ш': "sh", 'щ': "shch", 'ь': "", 'ю': "yu", 'я': "ya", ' ': "_",
	}

	var output strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) && unicode.Is(unicode.Cyrillic, r) {
			output.WriteString(translitMap[unicode.ToLower(r)])
		} else {
			output.WriteRune(r)
		}
	}
	return output.String()
}
