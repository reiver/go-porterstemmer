package porterstemmer



import (
//	"log"
	//"unicode"
)

var step1Suffixes [10][]byte = [10][]byte{
	[]byte("sses"),
	[]byte("ies"),
	[]byte("ss"),
	[]byte("s"),
	[]byte("eed"),
	[]byte("ed"),
	[]byte("at"),
	[]byte("bl"),
	[]byte("iz"),
	[]byte("ing"),
}

var step2Suffixes [21][]byte = [21][]byte{
	[]byte("ational"),
	[]byte("tional"),
	[]byte("enci"),
	[]byte("anci"),
	[]byte("izer"),
	[]byte("bli"),
	[]byte("alli"),
	[]byte("entli"),
	[]byte("eli"),
	[]byte("ousli"),
	[]byte("ization"),
	[]byte("ation"),
	[]byte("ator"),
	[]byte("alism"),
	[]byte("iveness"),
	[]byte("fulness"),
	[]byte("ousness"),
	[]byte("aliti"),
	[]byte("iviti"),
	[]byte("biliti"),
	[]byte("logi"),
}

var step3Suffixes [7][]byte = [7][]byte{
	[]byte("icate"),
	[]byte("ative"),
	[]byte("alize"),
	[]byte("iciti"),
	[]byte("ical"),
	[]byte("ful"),
	[]byte("ness"),
}

var step4Suffixes [19][]byte = [19][]byte{
	[]byte("al"),
	[]byte("ance"),
	[]byte("ence"),
	[]byte("er"),
	[]byte("ic"),
	[]byte("able"),
	[]byte("ible"),
	[]byte("ant"),
	[]byte("ement"),
	[]byte("ment"),
	[]byte("ent"),
	[]byte("ion"),
	[]byte("ou"),
	[]byte("ism"),
	[]byte("ate"),
	[]byte("iti"),
	[]byte("ous"),
	[]byte("ive"),
	[]byte("ize"),
}


func isConsonant(s []byte, i int) bool {

	//DEBUG
	//log.Printf("isConsonant: [%+v]", string(s[i]))

	result := true

	switch (  s[i]  ) {
		case 'a', 'e', 'i', 'o', 'u':
			result = false
		case 'y':
			if 0 == i {
				result = true
			} else {
				result = !isConsonant(s, i-1)
			}
		default: 
			result = true
   }

	return result
}



func measure(s []byte) uint {

	// Initialize.
		lenS := len(s)
		result := uint(0)
		i := 0


	// Short Circuit.
		if 0 == lenS {
/////////// RETURN
			return result
		}


	// Ignore (potential) consonant sequence at the beginning of word.
		for isConsonant(s, i) {

			//DEBUG
			//log.Printf("[measure([%s])] Eat Consonant [%d] -> [%s]", string(s), i, string(s[i]))

			i++
			if i >= lenS {
/////////////// RETURN
				return result
			}
		}


	// For each pair of a vowel sequence followed by a consonant sequence, increment result.
		Outer:
		for i < lenS {

			for !isConsonant(s, i) {

				//DEBUG
				//log.Printf("[measure([%s])] VOWEL [%d] -> [%s]", string(s), i, string(s[i]))

				i++
				if i >= lenS {
		/////////// BREAK
					break Outer
				}
			}
			for isConsonant(s, i) {

				//DEBUG
				//log.Printf("[measure([%s])] CONSONANT [%d] -> [%s]", string(s), i, string(s[i]))

				i++
				if i >= lenS {
					result++
		/////////// BREAK
					break Outer
				}
			}
			result++
		}


	// Return
		return result
}



func hasSuffix(s, suffix []byte) bool {

	lenSMinusOne      := len(s)      - 1
	lenSuffixMinusOne := len(suffix) - 1

	if lenSMinusOne < lenSuffixMinusOne {
		return false
	} else if s[lenSMinusOne] != suffix[lenSuffixMinusOne] { // I suspect checking this first should speed this function up in practice.
/////// RETURN
		return false
	} else {

		for i := 0; i < lenSuffixMinusOne ; i++ {

			if suffix[i] != s[lenSMinusOne-lenSuffixMinusOne+i] {
/////////////// RETURN
				return false
			}

		}

	}


	return true
}



func containsVowel(s []byte) bool {

	lenS := len(s)

	for i := 0 ; i < lenS ; i++ {

		if !isConsonant(s, i) {
/////////// RETURN
			return true
		}

	}

	return false
}



func hasRepeatDoubleConsonantSuffix(s []byte) bool {

	// Initialize.
		lenS := len(s)

		result := false


	// Do it!
		if 2 > lenS {
			result = false
		} else if s[lenS-1] == s[lenS-2] && isConsonant(s, lenS-1) { // Will using isConsonant() cause a problem with "YY"?
			result = true
		} else {
			result = false
		}


	// Return,
		return result
}



func hasConsonantVowelConsonantSuffix(s []byte) bool {

	// Initialize.
		lenS := len(s)

		result := false


	// Do it!
		if 3 > lenS {
			result = false
		} else if isConsonant(s, lenS-3) && !isConsonant(s, lenS-2) && isConsonant(s, lenS-1) {
			result = true
		} else  {
			result = false
		}


	// Return
		return result
}



func step1a(s []byte) []byte {

	// Initialize.
		var result []byte = s

		lenS := len(s)


	// Do it!
		if suffix := step1Suffixes[0] ; hasSuffix(s, suffix) {

			lenTrim := 2

			subSlice := s[:lenS-lenTrim]

			result = subSlice
		} else if suffix := step1Suffixes[1] ; hasSuffix(s, suffix) {
			lenTrim := 2

			subSlice := s[:lenS-lenTrim]

			result = subSlice
		} else if suffix := step1Suffixes[2] ; hasSuffix(s, suffix) {

			result = s
		} else if suffix := step1Suffixes[3] ; hasSuffix(s, suffix) {

			lenSuffix := 1

			subSlice := s[:lenS-lenSuffix]

			result = subSlice
		}	


	// Return.
		return result
}



func step1b(s []byte) []byte {

	// Initialize.
		var result []byte = s

		lenS := len(s)


	// Do it!
		if suffix := step1Suffixes[4] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if  0 < m {
				lenTrim := 1

				result = s[:lenS-lenTrim]
			}
		} else if suffix := step1Suffixes[5] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			if containsVowel(subSlice) {

				if suffix2 := step1Suffixes[6] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]
				} else if suffix2 := step1Suffixes[7] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]
				} else if suffix2 := step1Suffixes[8] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]
				} else if c := subSlice[len(subSlice)-1] ; 'l' != c && 's' != c && 'z' != c && hasRepeatDoubleConsonantSuffix(subSlice) {
					lenTrim := 1

					lenSubSlice := len(subSlice)

					result = subSlice[:lenSubSlice-lenTrim]
				} else if c := subSlice[len(subSlice)-1] ; 1 == measure(subSlice) && hasConsonantVowelConsonantSuffix(subSlice) && 'w' != c && 'x' != c && 'y' != c {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]

					result[len(result)-1] = 'e'
				} else {
					result = subSlice
				}

			}
		} else if suffix := step1Suffixes[9] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			if containsVowel(subSlice) {

				if suffix2 := step1Suffixes[6] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]

					result[len(result)-1] = 'e'
				} else if suffix2 := step1Suffixes[7] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]

					result[len(result)-1] = 'e'
				} else if suffix2 := step1Suffixes[8] ; hasSuffix(subSlice, suffix2) {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]

					result[len(result)-1] = 'e'
				} else if c := subSlice[len(subSlice)-1] ; 'l' != c && 's' != c && 'z' != c && hasRepeatDoubleConsonantSuffix(subSlice) {
					lenTrim := 1

					lenSubSlice := len(subSlice)

					result = subSlice[:lenSubSlice-lenTrim]
				} else if c := subSlice[len(subSlice)-1] ; 1 == measure(subSlice) && hasConsonantVowelConsonantSuffix(subSlice) && 'w' != c && 'x' != c && 'y' != c {
					lenTrim := -1

					result = s[:lenS-lenSuffix-lenTrim]

					result[len(result)-1] = 'e'
				} else {
					result = subSlice
				}

			}
		}


	// Return.
		return result
}



func step1c(s []byte) []byte {

	// Initialize.
		lenS := len(s)

		result := s


	// Do it!
		if 2 > lenS {
/////////// RETURN
			return result
		}

		if 'y' == s[lenS-1] && containsVowel(s[:lenS-1])  {

			result[lenS-1] = 'i';

		} else if 'Y' == s[lenS-1] && containsVowel(s[:lenS-1])  {

			result[lenS-1] = 'I';

		}


	// Return.
		return result
}



func step2(s []byte) []byte {

	// Initialize.
		lenS := len(s)

		result := s


	// Do it!
	var suffix []byte
		if suffix = step2Suffixes[0] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-5] = 'e'
				result = result[:lenS-4]
			}
		} else if suffix = step2Suffixes[1] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = result[:lenS-2]
			}
		} else if suffix = step2Suffixes[2] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-1] = 'e'
			}
		} else if suffix = step2Suffixes[3] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-1] = 'e'
			}
		} else if suffix = step2Suffixes[4] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-1]
			}
		} else if suffix = step2Suffixes[5] ; hasSuffix(s, suffix) { // --DEPARTURE--
//		} else if suffix := []byte("abli") ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-1] = 'e'
			}
		} else if suffix = step2Suffixes[6] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-2]
			}
		} else if suffix = step2Suffixes[7] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-2]
			}
		} else if suffix = step2Suffixes[8] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-2]
			}
		} else if suffix = step2Suffixes[9] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-2]
			}
		} else if suffix = step2Suffixes[10] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-5] = 'e'

				result = s[:lenS-4]
			}
		} else if suffix = step2Suffixes[11] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-3] = 'e'

				result = s[:lenS-2]
			}
		} else if suffix = step2Suffixes[12] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-2] = 'e'

				result = s[:lenS-1]
			}
		} else if suffix = step2Suffixes[13] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-3]
			}
		} else if suffix = step2Suffixes[14] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-4]
			}
		} else if suffix = step2Suffixes[15] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-4]
			}
		} else if suffix = step2Suffixes[16] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-4]
			}
		} else if suffix = step2Suffixes[17] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result = s[:lenS-3]
			}
		} else if suffix = step2Suffixes[18] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-3] = 'e'

				result = result[:lenS-2]
			}
		} else if suffix = step2Suffixes[19] ; hasSuffix(s, suffix) {
			if 0 < measure(s[:lenS-len(suffix)]) {
				result[lenS-5] = 'l'
				result[lenS-4] = 'e'

				result = result[:lenS-3]
			}
		} else if suffix = step2Suffixes[20] ; hasSuffix(s, suffix) { // --DEPARTURE--
			if 0 < measure(s[:lenS-len(suffix)]) {
				lenTrim := 1

				result = s[:lenS-lenTrim]
			}
		}


	// Return.
		return result
}



func step3(s []byte) []byte {

	// Initialize.
		lenS := len(s)
		result := s


	// Do it!
		if suffix := step3Suffixes[0] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			if 0 < measure(s[:lenS-lenSuffix]) {
				result = result[:lenS-3]
			}
		} else if suffix := step3Suffixes[1] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 0 < m {
				result = subSlice
			}
		} else if suffix := step3Suffixes[2] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			if 0 < measure(s[:lenS-lenSuffix]) {
				result = result[:lenS-3]
			}
		} else if suffix := step3Suffixes[3] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			if 0 < measure(s[:lenS-lenSuffix]) {
				result = result[:lenS-3]
			}
		} else if suffix := step3Suffixes[4] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			if 0 < measure(s[:lenS-lenSuffix]) {
				result = result[:lenS-2]
			}
		} else if suffix := step3Suffixes[5] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 0 < m {
				result = subSlice
			}
		} else if suffix := step3Suffixes[6] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 0 < m {
				result = subSlice
			}
		}


	// Return.
		return result
}



func step4(s []byte) []byte {

	// Initialize.
		lenS := len(s)
		result := s


	// Do it!
	var suffix []byte
		if suffix = step4Suffixes[0] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = result[:lenS-lenSuffix]
			}
		} else if suffix = step4Suffixes[1] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = result[:lenS-lenSuffix]
			}
		} else if suffix = step4Suffixes[2] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = result[:lenS-lenSuffix]
			}
		} else if suffix = step4Suffixes[3] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[4] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[5] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[6] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[7] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[8] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[9] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[10] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[11] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			if len(subSlice) == 0 {
				result = s
			} else {
				m := measure(subSlice)

				c := subSlice[len(subSlice)-1]

				if 1 < m && ('s' == c || 't' == c) {
					result = subSlice
				}
			}
		} else if suffix = step4Suffixes[12] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[13] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[14] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[15] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[16] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[17] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		} else if suffix = step4Suffixes[18] ; hasSuffix(s, suffix) {
			lenSuffix := len(suffix)

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		}


	// Return.
		return result
}



func step5a(s []byte) []byte {

	// Initialize.
		lenS := len(s)
		result := s


	// Do it!
		if 'e' == s[lenS-1] {
			lenSuffix := 1

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			} else if c := subSlice[len(subSlice)-1] ; 1 == m && !( hasConsonantVowelConsonantSuffix(subSlice) && 'w' != c && 'x' != c && 'y' != c)  {
				result = subSlice
			}
		}


	// Return.
		return result
}



func step5b(s []byte) []byte {

	// Initialize.
		lenS := len(s)
		result := s


	// Do it!
		if 2 < lenS && 'l' == s[lenS-2] && 'l' == s[lenS-1] {

			lenSuffix := 1

			subSlice := s[:lenS-lenSuffix]

			m := measure(subSlice)

			if 1 < m {
				result = subSlice
			}
		}


	// Return.
		return result
}
 



func StemString(s string) string {

	// Convert string to []byte
		byteArr := []byte(s)

	// Stem.
		byteArr = Stem(byteArr)

	// Convert []byte to string
		str := string(byteArr)

	// Return.
		return str
}

func Stem(s []byte) []byte {

	// Initialize.
		lenS := len(s)


	// Short circuit.
		if 0 == lenS {
/////////// RETURN
			return s
		}


	// Make all bytes lowercase.
		for i := 0 ; i < lenS ; i++ {
			//s[i] = unicode.ToLower(s[i])
		}


	// Stem
		result := StemWithoutLowerCasing(s)


	// Return.
		return result
}

func StemWithoutLowerCasing(s []byte) []byte {

	// Initialize.
		lenS := len(s)


	// Words that are of length 2 or less is already stemmed.
	// Don't do anything.
		if 2 >= lenS {
/////////// RETURN
			return s
		}


	// Stem
		s = step1a(s)
		s = step1b(s)
		s = step1c(s)
		s = step2(s)
		s = step3(s)
		s = step4(s)
		s = step5a(s)
		s = step5b(s)


	// Return.
		return s
}

