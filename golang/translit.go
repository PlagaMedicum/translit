// Package translit describes rules of the translation
package translit

var cyrJotConsonants []rune = []rune{
	'б', 'в', 'г', 'д', 'ж', 'з', 'к',
	'л', 'м', 'н', 'п', 'р', 'с', 'т', 'ф',
	'х', 'ц', 'ч', 'ш', 'щ',
}

var cyrVowels []rune = []rune{
	'а', 'е', 'ё', 'Ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я',
}

var unambPairs map[rune]string = map[rune]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'ж': "zh",
	'з': "z",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "shh",
	'ы': "y",
}

var latModifyableByH []rune = []rune{
	's', 'h', 'k', 'c', 'z', 'j', 'e',
}

func capitalize(r rune) rune {
	if (r >= 1072 && r <= 1103) ||
		(r >= 97 && r <= 122) {
		return r - 32
	} else if r >= 1104 && r <= 1119 {
		return r - 80
	}
	return r
}

func decapitalize(r rune) rune {
	if (r >= 1040 && r <= 1071) ||
		(r >= 65 && r <= 90) {
		return r + 32
	} else if r >= 1024 && r <= 1039 {
		return r + 80
	}
	return r
}

func isInArr(r rune, arr []rune) bool {
	for _, e := range arr {
		if r == decapitalize(e) || r == capitalize(e) {
			return true
		}
	}
	return false
}

func isCyr(r rune) bool {
	if (r >= 1040 && r <= 1103) ||
		(r == 'ё' || r == 'Ё') {
		return true
	}
	return false
}

func isWordEnding(i int, txt []rune) bool {
	if i == len(txt)-1 ||
		(i < len(txt)-1 && !isCyr(txt[i+1])) {
		return true
	}
	return false
}

func isWordStart(i int, txt []rune) bool {
	if i == 0 ||
		(i > 0 && i < len(txt) && !isCyr(txt[i-1])) {
		return true
	}
	return false
}

func isCyrCapitalized(r rune) bool {
	if (r >= 1040 && r <= 1071) ||
		(r >= 65 && r <= 90) ||
		(r >= 1024 && r <= 1039) {
		return true
	}
	return false
}

func isAfter(i int, input []rune, l rune) bool {
	return i > 0 && (input[i-1] == decapitalize(l) || input[i-1] == capitalize(l))
}

func isAfterOneOf(i int, input []rune, letters []rune) bool {
	return i > 0 && isInArr(input[i-1], letters)
}

func isBefore(i int, input []rune, l rune) bool {
	return i < len(input)-1 && (input[i+1] == decapitalize(l) || input[i+1] == capitalize(l))
}

func isBeforeOneOf(i int, input []rune, letters []rune) bool {
	return i < len(input)-1 && isInArr(input[i+1], letters)
}

// CyrToLat ...
func CyrToLat(s string) string {
	var res []rune
	input := []rune(s)

	for i, r := range input {
		if !isCyr(r) {
			res = append(res, r)
			continue
		}
		if p, b := unambPairs[r]; b {
			res = append(res, []rune(p)...)
			continue
		} else if p, b := unambPairs[r+32]; b {
			rp := []rune(p)
			rp[0] -= 32
			res = append(res, rp...)
			continue
		}

		inr := r
		r = decapitalize(r)

		var seq []rune
		switch r {
		case 'е':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'e'}
				break
			}
			seq = []rune{'j', 'e'}
		case 'ё':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'i', 'o'}
				break
			}
			seq = []rune{'j', 'o'}
		case 'и':
			if isAfterOneOf(i, input, cyrJotConsonants) &&
				isBeforeOneOf(i, input, []rune{'а', 'о', 'у'}) {
				seq = []rune{'j', 'i'}
				break
			}
			seq = []rune{'i'}
		case 'й':
			if (isWordStart(i, input) || isAfterOneOf(i, input, cyrVowels) || isAfter(i, input, 'ь')) &&
				isBeforeOneOf(i, input, []rune{'э', 'а', 'о', 'у'}) {
				seq = []rune{'j', 'i'}
				break
			} else if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'j', 'h'}
				break
			}
			seq = []rune{'j'}
		case 'х':
			if isInArr(res[len(res)-1], latModifyableByH) {
				seq = []rune{'k', 'h'}
				break
			}
			seq = []rune{'h'}
		case 'ь':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				if isBeforeOneOf(i, input, []rune{'э', 'а', 'о', 'у'}) {
					seq = []rune{'j', 'i'}
				} else {
					seq = []rune{'j'}
				}
				break
			}
			seq = []rune{'j', 'h'}
		case 'ъ':
			if isAfterOneOf(i, input, cyrJotConsonants) && !isWordEnding(i, input) {
				continue
			}
			seq = []rune{'y', 'h'}
		case 'э':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'e', 'h'}
				break
			}
			seq = []rune{'e'}
		case 'ю':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'i', 'u'}
				break
			}
			seq = []rune{'j', 'u'}
		case 'я':
			if isAfterOneOf(i, input, cyrJotConsonants) {
				seq = []rune{'i', 'a'}
				break
			}
			seq = []rune{'j', 'a'}
		}

		if isCyrCapitalized(inr) {
			seq[0] = capitalize(seq[0])
		}
		res = append(res, seq...)
	}

	return string(res)
}
