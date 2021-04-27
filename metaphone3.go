/*

Copyright 2010, Lawrence Philips
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

*/

/*
Translated by Andy Jackson.

 * A request from the author: Please comment and sign any changes you make to
 * the Metaphone 3 reference implementation.
 * <br>
 * Please do NOT reformat this module to Refine's coding standard,
 * but instead keep the original format so that it can be more easily compared
 * to any modified fork of the original.
*/

/**
 * Metaphone 3<br>
 * VERSION 2.1.3
 *
 * by Lawrence Philips<br>
 *
 * Metaphone 3 is designed to return an *approximate* phonetic key (and an alternate
 * approximate phonetic key when appropriate) that should be the same for English
 * words, and most names familiar in the United States, that are pronounced *similarly*.
 * The key value is *not* intended to be an *exact* phonetic, or even phonemic,
 * representation of the word. This is because a certain degree of 'fuzziness' has
 * proven to be useful in compensating for variations in pronunciation, as well as
 * misheard pronunciations. For example, although americans are not usually aware of it,
 * the letter 's' is normally pronounced 'z' at the end of words such as "sounds".<br><br>
 *
 * The 'approximate' aspect of the encoding is implemented according to the following rules:<br><br>
 *
 * (1) All vowels are encoded to the same value - 'A'. If the parameter encodeVowels
 * is set to false, only *initial* vowels will be encoded at all. If encodeVowels is set
 * to true, 'A' will be encoded at all places in the word that any vowels are normally
 * pronounced. 'W' as well as 'Y' are treated as vowels. Although there are differences in
 * the pronunciation of 'W' and 'Y' in different circumstances that lead to their being
 * classified as vowels under some circumstances and as consonants in others, for the purposes
 * of the 'fuzziness' component of the Soundex and Metaphone family of algorithms they will
 * be always be treated here as vowels.<br><br>
 *
 * (2) Voiced and un-voiced consonant pairs are mapped to the same encoded value. This
 * means that:<br>
 * 'D' and 'T' -> 'T'<br>
 * 'B' and 'P' -> 'P'<br>
 * 'G' and 'K' -> 'K'<br>
 * 'Z' and 'S' -> 'S'<br>
 * 'V' and 'F' -> 'F'<br><br>
 *
 * - In addition to the above voiced/unvoiced rules, 'CH' and 'SH' -> 'X', where 'X'
 * represents the "-SH-" and "-CH-" sounds in Metaphone 3 encoding.<br><br>
 *
 * - Also, the sound that is spelled as "TH" in English is encoded to '0' (zero symbol). (Although
 * Americans are not usually aware of it, "TH" is pronounced in a voiced (e.g. "that") as
 * well as an unvoiced (e.g. "theater") form, which are naturally mapped to the same encoding.)<br><br>
 *
 * The encodings in this version of Metaphone 3 are according to pronunciations common in the
 * United States. This means that they will be inaccurate for consonant pronunciations that
 * are different in the United Kingdom, for example "tube" -> "CHOOBE" -> XAP rather than american TAP.<br><br>
 *
 * Metaphone 3 was preceded by by Soundex, patented in 1919, and Metaphone and Double Metaphone,
 * developed by Lawrence Philips. All of these algorithms resulted in a significant number of
 * incorrect encodings. Metaphone3 was tested against a database of about 100 thousand English words,
 * names common in the United States, and non-English words found in publications in the United States,
 * with an emphasis on words that are commonly mispronounced, prepared by the Moby Words website,
 * but with the Moby Words 'phonetic' encodings algorithmically mapped to Double Metaphone encodings.
 * Metaphone3 increases the accuracy of encoding of english words, common names, and non-English
 * words found in american publications from the 89% for Double Metaphone, to over 98%.<br><br>
 *
 * DISCLAIMER:
 * Anthropomorphic Software LLC claims only that Metaphone 3 will return correct encodings,
 * within the 'fuzzy' definition of correct as above, for a very high percentage of correctly
 * spelled English and commonly recognized non-English words. Anthropomorphic Software LLC
 * warns the user that a number of words remain incorrectly encoded, that misspellings may not
 * be encoded 'properly', and that people often have differing ideas about the pronunciation
 * of a word. Therefore, Metaphone 3 is not guaranteed to return correct results every time, and
 * so a desired target word may very well be missed. Creators of commercial products should
 * keep in mind that systems like Metaphone 3 produce a 'best guess' result, and should
 * condition the expectations of end users accordingly.<br><br>
 *
 * METAPHONE3 IS PROVIDED "AS IS" WITHOUT
 * WARRANTY OF ANY KIND. LAWRENCE PHILIPS AND ANTHROPOMORPHIC SOFTWARE LLC
 * MAKE NO WARRANTIES, EXPRESS OR IMPLIED, THAT IT IS FREE OF ERROR,
 * OR ARE CONSISTENT WITH ANY PARTICULAR STANDARD OF MERCHANTABILITY,
 * OR THAT IT WILL MEET YOUR REQUIREMENTS FOR ANY PARTICULAR APPLICATION.
 * LAWRENCE PHILIPS AND ANTHROPOMORPHIC SOFTWARE LLC DISCLAIM ALL LIABILITY
 * FOR DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES RESULTING FROM USE
 * OF THIS SOFTWARE.
 *
 * @author Lawrence Philips
 *
 * Metaphone 3 is designed to return an <i>approximate</i> phonetic key (and an alternate
 * approximate phonetic key when appropriate) that should be the same for English
 * words, and most names familiar in the United States, that are pronounced "similarly".
 * The key value is <i>not</i> intended to be an exact phonetic, or even phonemic,
 * representation of the word. This is because a certain degree of 'fuzziness' has
 * proven to be useful in compensating for variations in pronunciation, as well as
 * misheard pronunciations. For example, although americans are not usually aware of it,
 * the letter 's' is normally pronounced 'z' at the end of words such as "sounds".<br><br>
 *
 * The 'approximate' aspect of the encoding is implemented according to the following rules:<br><br>
 *
 * (1) All vowels are encoded to the same value - 'A'. If the parameter encodeVowels
 * is set to false, only *initial* vowels will be encoded at all. If encodeVowels is set
 * to true, 'A' will be encoded at all places in the word that any vowels are normally
 * pronounced. 'W' as well as 'Y' are treated as vowels. Although there are differences in
 * the pronunciation of 'W' and 'Y' in different circumstances that lead to their being
 * classified as vowels under some circumstances and as consonants in others, for the purposes
 * of the 'fuzziness' component of the Soundex and Metaphone family of algorithms they will
 * be always be treated here as vowels.<br><br>
 *
 * (2) Voiced and un-voiced consonant pairs are mapped to the same encoded value. This
 * means that:<br>
 * 'D' and 'T' -> 'T'<br>
 * 'B' and 'P' -> 'P'<br>
 * 'G' and 'K' -> 'K'<br>
 * 'Z' and 'S' -> 'S'<br>
 * 'V' and 'F' -> 'F'<br><br>
 *
 * - In addition to the above voiced/unvoiced rules, 'CH' and 'SH' -> 'X', where 'X'
 * represents the "-SH-" and "-CH-" sounds in Metaphone 3 encoding.<br><br>
 *
 * - Also, the sound that is spelled as "TH" in English is encoded to '0' (zero symbol). (Although
 * americans are not usually aware of it, "TH" is pronounced in a voiced (e.g. "that") as
 * well as an unvoiced (e.g. "theater") form, which are naturally mapped to the same encoding.)<br><br>
 *
 * In the "Exact" encoding, voiced/unvoiced pairs are <i>not</i> mapped to the same encoding, except
 * for the voiced and unvoiced versions of 'TH', sounds such as 'CH' and 'SH', and for 'S' and 'Z',
 * so that the words whose metaph keys match will in fact be closer in pronunciation that with the
 * more approximate setting. Keep in mind that encoding settings for search strings should always
 * be exactly the same as the encoding settings of the stored metaph keys in your database!
 * Because of the considerably increased accuracy of Metaphone3, it is now possible to use this
 * setting and have a very good chance of getting a correct encoding.
 * <br><br>
 * In the Encode Vowels encoding, all non-initial vowels and diphthongs will be encoded to
 * 'A', and there will only be one such vowel encoding character between any two consonants.
 * It turns out that there are some surprising wrinkles to encoding non-initial vowels in
 * practice, pre-eminently in inversions between spelling and pronunciation such as e.g.
 * "wrinkle" => 'RANKAL', where the last two sounds are inverted when spelled.
 * <br><br>
 * The encodings in this version of Metaphone 3 are according to pronunciations common in the
 * United States. This means that they will be inaccurate for consonant pronunciations that
 * are different in the United Kingdom, for example "tube" -> "CHOOBE" -> XAP rather than american TAP.
 * <br><br>
 *
 */
package metaphone3

import (
	"strings"
)

/** Default size of key storage allocation */
var MAX_KEY_ALLOCATION = 32

/** Default maximum length of encoded key. */
var DEFAULT_MAX_KEY_LENGTH = 8

type M3 struct {
	/** Flag whether or not to encode non-initial vowels. */
	encodeVowels bool

	/** Flag whether or not to encode consonants as exactly
	* as possible. */
	encodeExact bool

	/** Length of word sent in to be encoded, as
	* measured at beginning of encoding. */
	length int

	/** Running copy of primary key. */
	primary strings.Builder

	/** Running copy of secondary key. */
	secondary strings.Builder

	/** Index of character in m.inWord currently being
	* encoded. */
	current int

	/** Index of last character in m.inWord. */
	last int

	/** Flag that an AL inversion has already been done. */
	flag_AL_inversion bool

	/** Length of encoded key string. */
	metaphLength int

	/** Internal copy of word to be encoded, allocated separately
	* from pointed to in incoming parameter string. */
	inWord string
}

////////////////////////////////////////////////////////////////////////////////
// Metaphone3 class definition
////////////////////////////////////////////////////////////////////////////////

/**
 * Constructor, default. This constructor is most convenient when
 * encoding more than one word at a time. New words to encode can
 * be set using SetWord(char *).
 *
 */
func New() *M3 {
	return &M3{
		metaphLength: DEFAULT_MAX_KEY_LENGTH,
	}
}

/**
 * Sets length allocated for output keys.
 * If incoming number is greater than maximum allowable
 * length returned by GetMaximumKeyLength(), set key length
 * to maximum key length and return false;  otherwise, set key
 * length to parameter value and return true.
 *
 * @param inKeyLength new length of key.
 * @return true if able to set key length to requested value.
 *
 */
func (m *M3) SetKeyLength(inKeyLength int) bool {
	if inKeyLength < 1 {
		// can't have that -
		// no room for terminating null
		inKeyLength = 1
	}

	if inKeyLength > MAX_KEY_ALLOCATION {
		m.metaphLength = MAX_KEY_ALLOCATION
		return false
	}

	m.metaphLength = inKeyLength
	return true
}

/**
 * Adds an encoding character to the encoded key value string - two parameter version
 *
 * @param main primary encoding character to be added to encoded key string
 * @param alt alternative encoding character to be added to encoded alternative key string
 *
 */
func (m *M3) metaphAdd(main string, alt string) {
	if !(main == "A" && (m.primary.Len() > 0) && (m.primary.String()[m.primary.Len()-1] == 'A')) {
		m.primary.WriteString(main)
	}

	if !(alt == "A" && (m.secondary.Len() > 0) && (m.secondary.String()[m.secondary.Len()-1] == 'A')) {
		if alt != "" {
			m.secondary.WriteString(alt)
		}
	}
}

/**
 * Adds an encoding character to the encoded key value string - Exact/Approx version
 *
 * @param mainExact primary encoding character to be added to encoded key if string
 * m.encodeExact is set
 *
 * @param altExact alternative encoding character to be added to encoded alternative
 * key if m.encodeExact is set string
 *
 * @param main primary encoding character to be added to encoded key string
 *
 * @param alt alternative encoding character to be added to encoded alternative key string
 *
 */
func (m *M3) metaphAddExactApprox4(mainExact string, altExact string, main string, alt string) {
	if m.encodeExact {
		m.metaphAdd(mainExact, altExact)
	} else {
		m.metaphAdd(main, alt)
	}
}

/**
 * Adds an encoding character to the encoded key value string - Exact/Approx version
 *
 * @param mainExact primary encoding character to be added to encoded key if string
 * m.encodeExact is set
 *
 * @param main primary encoding character to be added to encoded key string
 *
 */
func (m *M3) metaphAddExactApprox(mainExact string, main string) {
	if m.encodeExact {
		m.metaphAdd(mainExact, mainExact)
	} else {
		m.metaphAdd(main, main)
	}
}

/** Sets flag that causes Metaphone3 to encode non-initial vowels. However, even
 * if there are more than one vowel sound in a vowel sequence (i.e.
 * vowel diphthong, etc.), only one 'A' will be encoded before the next consonant or the
 * end of the word.
 *
 * @param inEncodeVowels Non-initial vowels encoded if true, not if false.
 */
func (m *M3) SetEncodeVowels(inEncodeVowels bool) { m.encodeVowels = inEncodeVowels }

/** Sets flag that causes Metaphone3 to encode consonants as exactly as possible.
 * This does not include 'S' vs. 'Z', since americans will pronounce 'S' at the
 * at the end of many words as 'Z', nor does it include "CH" vs. "SH". It does cause
 * a distinction to be made between 'B' and 'P', 'D' and 'T', 'G' and 'K', and 'V'
 * and 'F'.
 *
 * @param inEncodeExact consonants to be encoded "exactly" if true, not if false.
 */
func (m *M3) SetEncodeExact(inEncodeExact bool) { m.encodeExact = inEncodeExact }

/**
 * Test for close front vowels
 *
 * @return true if close front vowel
 */
func (m *M3) front_Vowel(at int) bool {
	return ((m.charAt(at) == 'E') || (m.charAt(at) == 'I') || (m.charAt(at) == 'Y'))
}

/**
 * Detect names or words that begin with spellings
 * typical of german or slavic words, for the purpose
 * of choosing alternate pronunciations correctly
 *
 */
func (m *M3) slavoGermanic() bool {
	return m.stringAt(0, 3, "SCH", "") || m.stringAt(0, 2, "SW", "") || (m.charAt(0) == 'J') || (m.charAt(0) == 'W')
}

/**
 * Tests if character is a vowel
 *
 * @param inChar character to be tested in to be encoded string
 * @return true if character is a vowel, false if not
 *
 */
func isVowel(inChar rune) bool {
	return (inChar == 'A') || (inChar == 'E') || (inChar == 'I') || (inChar == 'O') || (inChar == 'U') || (inChar == 'Y') || (inChar == 'À') || (inChar == 'Á') || (inChar == 'Â') || (inChar == 'Ã') || (inChar == 'Ä') || (inChar == 'Å') || (inChar == 'Æ') || (inChar == 'È') || (inChar == 'É') || (inChar == 'Ê') || (inChar == 'Ë') || (inChar == 'Ì') || (inChar == 'Í') || (inChar == 'Î') || (inChar == 'Ï') || (inChar == 'Ò') || (inChar == 'Ó') || (inChar == 'Ô') || (inChar == 'Õ') || (inChar == 'Ö') || (inChar == '') || (inChar == 'Ø') || (inChar == 'Ù') || (inChar == 'Ú') || (inChar == 'Û') || (inChar == 'Ü') || (inChar == 'Ý') || (inChar == '')
}

/**
 * skips over vowels in a string. Has exceptions for skipping consonants that
 * will not be encoded.
 *
 * @param at position, in to be encoded string, of character to start skipping from
 *
 * @return position of next consonant in to be encoded string
 */
func (m *M3) skipVowels(at int) int {
	if at < 0 {
		return 0
	}

	if at >= m.length {
		return m.length
	}
	it := m.charAt(at)

	for isVowel(it) || (it == 'W') {
		if m.stringAt(at, 4, "WICZ", "WITZ", "WIAK", "") || m.stringAt((at-1), 5, "EWSKI", "EWSKY", "OWSKI", "OWSKY", "") || (m.stringAt(at, 5, "WICKI", "WACKI", "") && ((at + 4) == m.last)) {
			break
		}

		at++
		if ((m.charAt(at-1) == 'W') && (m.charAt(at) == 'H')) && !(m.stringAt(at, 3, "HOP", "") || m.stringAt(at, 4, "HIDE", "HARD", "HEAD", "HAWK", "HERD", "HOOK", "HAND", "HOLE", "") || m.stringAt(at, 5, "HEART", "HOUSE", "HOUND", "") || m.stringAt(at, 6, "HAMMER", "")) {
			at++
		}

		if at > (m.length - 1) {
			break
		}
		it = m.charAt(at)
	}

	return at
}

/**
 * Advanced counter m.current so that it indexes the next character to be encoded
 *
 * @param ifNotEncodeVowels number of characters to advance if not encoding internal vowels
 * @param ifEncodeVowels number of characters to advance if encoding internal vowels
 *
 */
func (m *M3) advanceCounter(ifNotEncodeVowels, ifEncodeVowels int) {
	if !m.encodeVowels {
		m.current += ifNotEncodeVowels
	} else {
		m.current += ifEncodeVowels
	}
}

/**
     * Subscript safe charAt()
     *
	 * @param at index of character to access
	 * @return null if index out of bounds, .charAt() otherwise
*/
func (m *M3) charAt(at int) rune {
	// check subbounds string
	if (at < 0) || (at > (m.length - 1)) {
		return rune(0)
	}

	for i, r := range m.inWord {
		if i == at {
			return r
		}
	}
	return rune(0)
}

/**
 * Tests whether the word is the root or a regular english inflection
 * of it, e.g. "ache", "achy", "aches", "ached", "aching", "achingly"
 * This is for cases where we want to match only the root and corresponding
 * inflected forms, and not completely different words which may have the
 * same subin them string.
 */
func rootOrInflections(inWord string, root string) bool {
	var test string

	test = root + "S"
	if (inWord == root) || (inWord == test) {
		return true
	}

	if root[len(root)-1] != 'E' {
		test = root + "ES"
	}

	if inWord == test {
		return true
	}

	if root[len(root)-1] != 'E' {
		test = root + "ED"
	} else {
		test = root + "D"
	}

	if inWord == test {
		return true
	}

	if root[len(root)-1] == 'E' {
		root = root[:len(root)-1]
	}

	test = root + "ING"
	if inWord == test {
		return true
	}

	test = root + "INGLY"
	if inWord == test {
		return true
	}

	test = root + "Y"
	if inWord == test {
		return true
	}

	return false
}

/**
 * Determines if one of the substrings sent in is the same as
 * what is at the specified position in the being encoded string.
 *
 * @param start
 * @param length
 * @param compareStrings
 * @return
 */
func (m *M3) stringAt(start, length int, compareStrings ...string) bool {
	// check subbounds string
	if (start < 0) || (start > (m.length - 1)) || ((start + length - 1) > (m.length - 1)) {
		return false
	}

	b := strings.Builder{}
	for i, r := range m.inWord {
		if i >= start {
			b.WriteRune(r)
		}
		if i == start+length {
			break
		}
	}
	target := b.String()

	for _, strFragment := range compareStrings {
		if target == strFragment {
			return true
		}
	}
	return false
}

/**
 * Encodes input to one or two key values string according to Metaphone 3 rules.
 *
 */
func (m *M3) Encode(in string) (primary, secondary string) {
	m.flag_AL_inversion = false

	m.current = 0

	m.inWord = strings.ToUpper(in)
	m.primary.Reset()
	m.secondary.Reset()

	m.length = 0
	for _ = range in {
		m.length++
	}
	if m.length < 1 {
		return
	}

	//zero based index
	m.last = m.length - 1

	///////////main loop//////////////////////////
	for !(m.primary.Len() > m.metaphLength) && !(m.secondary.Len() > m.metaphLength) {
		if m.current >= m.length {
			break
		}

		switch m.charAt(m.current) {
		case 'B':

			m.encode_B()
			break

		case 'ß':
		case 'Ç':

			m.metaphAdd("S", "S")
			m.current++
			break

		case 'C':

			m.encode_C()
			break

		case 'D':

			m.encode_D()
			break

		case 'F':

			m.encode_F()
			break

		case 'G':

			m.encode_G()
			break

		case 'H':

			m.encode_H()
			break

		case 'J':

			m.encode_J()
			break

		case 'K':

			m.encode_K()
			break

		case 'L':

			m.encode_L()
			break

		case 'M':

			m.encode_M()
			break

		case 'N':

			m.encode_N()
			break

		case 'Ñ':

			m.metaphAdd("N", "N")
			m.current++
			break

		case 'P':

			m.encode_P()
			break

		case 'Q':

			m.encode_Q()
			break

		case 'R':

			m.encode_R()
			break

		case 'S':

			m.encode_S()
			break

		case 'T':

			m.encode_T()
			break

		case 'Ð': // eth
		case 'Þ': // thorn

			m.metaphAdd("0", "0")
			m.current++
			break

		case 'V':

			m.encode_V()
			break

		case 'W':

			m.encode_W()
			break

		case 'X':

			m.encode_X()
			break

		case '':

			m.metaphAdd("X", "X")
			m.current++
			break

		case '':

			m.metaphAdd("S", "S")
			m.current++
			break

		case 'Z':

			m.encode_Z()
			break

		default:

			if isVowel(m.charAt(m.current)) {
				m.encode_Vowels()
				break
			}

			m.current++

		}
	}

	primary, secondary = m.primary.String(), m.secondary.String()

	//only give back m.metaphLength number of chars in m.metaph
	if len(primary) > m.metaphLength {
		primary = primary[:m.metaphLength]
	}

	if len(secondary) > m.metaphLength {
		secondary = secondary[:m.metaphLength]
	}

	// it is possible for the two metaphs to be the same
	// after truncation. lose the second one if so
	if primary == secondary {
		secondary = ""
	}

	return primary, secondary
}

/**
 * Encodes all initial vowels to A.
 *
 * Encodes non-initial vowels to A if m.encodeVowels is true
 *
 *
 */
func (m *M3) encode_Vowels() {
	if m.current == 0 {
		// all init vowels map to 'A'
		// as of Double Metaphone
		m.metaphAdd("A", "A")
	} else if m.encodeVowels {
		if m.charAt(m.current) != 'E' {
			if m.skip_Silent_UE() {
				return
			}

			if m.o_Silent() {
				m.current++
				return
			}

			// encode all vowels and
			// diphthongs to the same value
			m.metaphAdd("A", "A")
		} else {
			m.encode_E_Pronounced()
		}
	}

	if !(!isVowel(m.charAt(m.current-2)) && m.stringAt((m.current-1), 4, "LEWA", "LEWO", "LEWI", "")) {
		m.current = m.skipVowels(m.current)
	} else {
		m.current++
	}
}

/**
 * Encodes cases where non-initial 'e' is pronounced, taking
 * care to detect unusual cases from the greek.
 *
 * Only executed if non initial vowel encoding is turned on
 *
 *
 */
func (m *M3) encode_E_Pronounced() {
	// special cases with two pronunciations
	// 'agape' 'lame' 'resume'
	if (m.stringAt(0, 4, "LAME", "SAKE", "PATE", "") && (m.length == 4)) || (m.stringAt(0, 5, "AGAPE", "") && (m.length == 5)) || ((m.current == 5) && m.stringAt(0, 6, "RESUME", "")) {
		m.metaphAdd("", "A")
		return
	}

	// special case "inge" => 'INGA', 'INJ'
	if m.stringAt(0, 4, "INGE", "") && (m.length == 4) {
		m.metaphAdd("A", "")
		return
	}

	// special cases with two pronunciations
	// special handling due to the difference in
	// the pronunciation of the '-D'
	if (m.current == 5) && m.stringAt(0, 7, "BLESSED", "LEARNED", "") {
		m.metaphAddExactApprox4("D", "AD", "T", "AT")
		m.current += 2
		return
	}

	// encode all vowels and diphthongs to the same value
	if (!m.e_Silent() && !m.flag_AL_inversion && !m.silent_Internal_E()) || m.e_Pronounced_Exceptions() {
		m.metaphAdd("A", "A")
	}

	// now that we've visited the vowel in question
	m.flag_AL_inversion = false
}

/**
 * Tests for cases where non-initial 'o' is not pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m.metaph key
 *
 */
func (m *M3) o_Silent() bool {
	// if "iron" at beginning or end of word and not "irony"
	if (m.charAt(m.current) == 'O') && m.stringAt((m.current-2), 4, "IRON", "") {
		if (m.stringAt(0, 4, "IRON", "") || (m.stringAt((m.current-2), 4, "IRON", "") && (m.last == (m.current + 1)))) && !m.stringAt((m.current-2), 6, "IRONIC", "") {
			return true
		}
	}

	return false
}

/**
 * Tests and encodes cases where non-initial 'e' is never pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m.metaph key
 *
 */
func (m *M3) e_Silent() bool {
	if m.e_Pronounced_At_End() {
		return false
	}

	// 'e' silent when last letter, altho
	// also silent if before plural 's'
	// or past tense or participle 'd', e.g.
	// 'grapes' and 'banished' => PNXT
	// and not e.g. "nested", "rises", or "pieces" => RASAS
	// e.g.  'wholeness', 'boneless', 'barely'
	if (m.current == m.last) || (m.stringAt(m.last, 1, "S", "D", "") && (m.current > 1) && ((m.current + 1) == m.last) && !(m.stringAt((m.current-1), 3, "TED", "SES", "CES", "") || m.stringAt(0, 9, "ANTIPODES", "ANOPHELES", "") || m.stringAt(0, 8, "MOHAMMED", "MUHAMMED", "MOUHAMED", "") || m.stringAt(0, 7, "MOHAMED", "") || m.stringAt(0, 6, "NORRED", "MEDVED", "MERCED", "ALLRED", "KHALED", "RASHED", "MASJED", "") || m.stringAt(0, 5, "JARED", "AHMED", "HAMED", "JAVED", "") || m.stringAt(0, 4, "ABED", "IMED", ""))) || (m.stringAt((m.current+1), 4, "NESS", "LESS", "") && ((m.current + 4) == m.last)) || (m.stringAt((m.current+1), 2, "LY", "") && ((m.current + 2) == m.last) && !m.stringAt(0, 6, "CICELY", "")) {
		return true
	}

	return false
}

/**
 * Tests for words where an 'E' at the end of the word
 * is pronounced
 *
 * special cases, mostly from the greek, spanish, japanese,
 * italian, and french words normally having an acute accent.
 * also, pronouns and articles
 *
 * Many Thanks to ali, QuentinCompson, JeffCO, ToonScribe, Xan,
 * Trafalz, and VictorLaszlo, all of them atriots from the Eschaton,
 * for all their fine contributions!
 *
 * @return true if 'E' at end is pronounced
 *
 */
func (m *M3) e_Pronounced_At_End() bool {
	if (m.current == m.last) && (m.stringAt((m.current-6), 7, "STROPHE", "") ||
		// if a vowel is before the 'E', vowel eater will have eaten it.
		//otherwise, consonant + 'E' will need 'E' pronounced
		(m.length == 2) || ((m.length == 3) && !isVowel(m.charAt(0))) ||
		// these german name endings can be relied on to have the 'e' pronounced
		(m.stringAt((m.last-2), 3, "BKE", "DKE", "FKE", "KKE", "LKE",
			"NKE", "MKE", "PKE", "TKE", "VKE", "ZKE", "") && !m.stringAt(0, 5, "FINKE", "FUNKE", "") && !m.stringAt(0, 6, "FRANKE", "")) || m.stringAt((m.last-4), 5, "SCHKE", "") || (m.stringAt(0, 4, "ACME", "NIKE", "CAFE", "RENE", "LUPE", "JOSE", "ESME", "") && (m.length == 4)) || (m.stringAt(0, 5, "LETHE", "CADRE", "TILDE", "SIGNE", "POSSE", "LATTE", "ANIME", "DOLCE", "CROCE",
		"ADOBE", "OUTRE", "JESSE", "JAIME", "JAFFE", "BENGE", "RUNGE",
		"CHILE", "DESME", "CONDE", "URIBE", "LIBRE", "ANDRE", "") && (m.length == 5)) || (m.stringAt(0, 6, "HECATE", "PSYCHE", "DAPHNE", "PENSKE", "CLICHE", "RECIPE",
		"TAMALE", "SESAME", "SIMILE", "FINALE", "KARATE", "RENATE", "SHANTE",
		"OBERLE", "COYOTE", "KRESGE", "STONGE", "STANGE", "SWAYZE", "FUENTE",
		"SALOME", "URRIBE", "") && (m.length == 6)) || (m.stringAt(0, 7, "ECHIDNE", "ARIADNE", "MEINEKE", "PORSCHE", "ANEMONE", "EPITOME",
		"SYNCOPE", "SOUFFLE", "ATTACHE", "MACHETE", "KARAOKE", "BUKKAKE",
		"VICENTE", "ELLERBE", "VERSACE", "") && (m.length == 7)) || (m.stringAt(0, 8, "PENELOPE", "CALLIOPE", "CHIPOTLE", "ANTIGONE", "KAMIKAZE", "EURIDICE",
		"YOSEMITE", "FERRANTE", "") && (m.length == 8)) || (m.stringAt(0, 9, "HYPERBOLE", "GUACAMOLE", "XANTHIPPE", "") && (m.length == 9)) || (m.stringAt(0, 10, "SYNECDOCHE", "") && (m.length == 10))) {
		return true
	}

	return false
}

/**
 * Detect internal silent 'E's e.g. "roseman",
 * "firestone"
 *
 */
func (m *M3) silent_Internal_E() bool {
	// 'olesen' but not 'olen'	RAKE BLAKE
	if (m.stringAt(0, 3, "OLE", "") && m.e_Silent_Suffix(3) && !m.e_Pronouncing_Suffix(3)) || (m.stringAt(0, 4, "BARE", "FIRE", "FORE", "GATE", "HAGE", "HAVE",
		"HAZE", "HOLE", "CAPE", "HUSE", "LACE", "LINE",
		"LIVE", "LOVE", "MORE", "MOSE", "MORE", "NICE",
		"RAKE", "ROBE", "ROSE", "SISE", "SIZE", "WARE",
		"WAKE", "WISE", "WINE", "") && m.e_Silent_Suffix(4) && !m.e_Pronouncing_Suffix(4)) || (m.stringAt(0, 5, "BLAKE", "BRAKE", "BRINE", "CARLE", "CLEVE", "DUNNE",
		"HEDGE", "HOUSE", "JEFFE", "LUNCE", "STOKE", "STONE",
		"THORE", "WEDGE", "WHITE", "") && m.e_Silent_Suffix(5) && !m.e_Pronouncing_Suffix(5)) || (m.stringAt(0, 6, "BRIDGE", "CHEESE", "") && m.e_Silent_Suffix(6) && !m.e_Pronouncing_Suffix(6)) || m.stringAt((m.current-5), 7, "CHARLES", "") {
		return true
	}

	return false
}

/**
 * Detect conditions required
 * for the 'E' not to be pronounced
 *
 */
func (m *M3) e_Silent_Suffix(at int) bool {
	if (m.current == (at - 1)) && (m.length > (at + 1)) && (isVowel(m.charAt(at+1)) || (m.stringAt(at, 2, "ST", "SL", "") && (m.length > (at + 2)))) {
		return true
	}

	return false
}

/**
 * Detect endings that will
 * cause the 'e' to be pronounced
 *
 */
func (m *M3) e_Pronouncing_Suffix(at int) bool {
	// e.g. 'bridgewood' - the other vowels will get eaten
	// up so we need to put one in here
	if (m.length == (at + 4)) && m.stringAt(at, 4, "WOOD", "") {
		return true
	}

	// same as above
	if (m.length == (at + 5)) && m.stringAt(at, 5, "WATER", "WORTH", "") {
		return true
	}

	// e.g. 'bridgette'
	if (m.length == (at + 3)) && m.stringAt(at, 3, "TTE", "LIA", "NOW", "ROS", "RAS", "") {
		return true
	}

	// e.g. 'olena'
	if (m.length == (at + 2)) && m.stringAt(at, 2, "TA", "TT", "NA", "NO", "NE",
		"RS", "RE", "LA", "AU", "RO", "RA", "") {
		return true
	}

	// e.g. 'bridget'
	if (m.length == (at + 1)) && m.stringAt(at, 1, "T", "R", "") {
		return true
	}

	return false
}

/**
 * Exceptions where 'E' is pronounced where it
 * usually wouldn't be, and also some cases
 * where 'LE' transposition rules don't apply
 * and the vowel needs to be encoded here
 *
 * @return true if 'E' pronounced
 *
 */
func (m *M3) e_Pronounced_Exceptions() bool {
	// greek names e.g. "herakles" or hispanic names e.g. "robles", where 'e' is pronounced, other exceptions
	if (((m.current + 1) == m.last) && (m.stringAt((m.current-3), 5, "OCLES", "ACLES", "AKLES", "") || m.stringAt(0, 4, "INES", "") || m.stringAt(0, 5, "LOPES", "ESTES", "GOMES", "NUNES", "ALVES", "ICKES",
		"INNES", "PERES", "WAGES", "NEVES", "BENES", "DONES", "") || m.stringAt(0, 6, "CORTES", "CHAVES", "VALDES", "ROBLES", "TORRES", "FLORES", "BORGES",
		"NIEVES", "MONTES", "SOARES", "VALLES", "GEDDES", "ANDRES", "VIAJES",
		"CALLES", "FONTES", "HERMES", "ACEVES", "BATRES", "MATHES", "") || m.stringAt(0, 7, "DELORES", "MORALES", "DOLORES", "ANGELES", "ROSALES", "MIRELES", "LINARES",
		"PERALES", "PAREDES", "BRIONES", "SANCHES", "CAZARES", "REVELES", "ESTEVES",
		"ALVARES", "MATTHES", "SOLARES", "CASARES", "CACERES", "STURGES", "RAMIRES",
		"FUNCHES", "BENITES", "FUENTES", "PUENTES", "TABARES", "HENTGES", "VALORES", "") || m.stringAt(0, 8, "GONZALES", "MERCEDES", "FAGUNDES", "JOHANNES", "GONSALES", "BERMUDES",
		"CESPEDES", "BETANCES", "TERRONES", "DIOGENES", "CORRALES", "CABRALES",
		"MARTINES", "GRAJALES", "") || m.stringAt(0, 9, "CERVANTES", "FERNANDES", "GONCALVES", "BENEVIDES", "CIFUENTES", "SIFUENTES",
		"SERVANTES", "HERNANDES", "BENAVIDES", "") || m.stringAt(0, 10, "ARCHIMEDES", "CARRIZALES", "MAGALLANES", ""))) || m.stringAt(m.current-2, 4, "FRED", "DGES", "DRED", "GNES", "") || m.stringAt((m.current-5), 7, "PROBLEM", "RESPLEN", "") || m.stringAt((m.current-4), 6, "REPLEN", "") || m.stringAt((m.current-3), 4, "SPLE", "") {
		return true
	}

	return false
}

/**
 * Encodes "-UE".
 *
 * @return true if encoding handled in this routine, false if not
 */
func (m *M3) skip_Silent_UE() bool {
	// always silent except for cases listed below
	if (m.stringAt((m.current-1), 3, "QUE", "GUE", "") && !m.stringAt(0, 8, "BARBEQUE", "PALENQUE", "APPLIQUE", "") &&
		// '-que' cases usually french but missing the acute accent
		!m.stringAt(0, 6, "RISQUE", "") && !m.stringAt((m.current-3), 5, "ARGUE", "SEGUE", "") && !m.stringAt(0, 7, "PIROGUE", "ENRIQUE", "") && !m.stringAt(0, 10, "COMMUNIQUE", "")) && (m.current > 1) && (((m.current + 1) == m.last) || m.stringAt(0, 7, "JACQUES", "")) {
		m.current = m.skipVowels(m.current)
		return true
	}

	return false
}

/**
 * Encodes 'B'
 *
 *
 */
func (m *M3) encode_B() {
	if m.encode_Silent_B() {
		return
	}

	// "-mb", e.g", "dumb", already skipped over under
	// 'M', altho it should really be handled here...
	m.metaphAddExactApprox("B", "P")

	if (m.charAt(m.current+1) == 'B') || ((m.charAt(m.current+1) == 'P') && ((m.current+1 < m.last) && (m.charAt(m.current+2) != 'H'))) {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encodes silent 'B' for cases not covered under "-mb-"
 *
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_B() bool {
	//'debt', 'doubt', 'subtle'
	if m.stringAt((m.current-2), 4, "DEBT", "") || m.stringAt((m.current-2), 5, "SUBTL", "") || m.stringAt((m.current-2), 6, "SUBTIL", "") || m.stringAt((m.current-3), 5, "DOUBT", "") {
		m.metaphAdd("T", "T")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes 'C'
 *
 */
func (m *M3) encode_C() {

	if m.encode_Silent_C_At_Beginning() || m.encode_CA_To_S() || m.encode_CO_To_S() || m.encode_CH() || m.encode_CCIA() || m.encode_CC() || m.encode_CK_CG_CQ() || m.encode_C_Front_Vowel() || m.encode_Silent_C() || m.encode_CZ() || m.encode_CS() {
		return
	}

	//else
	if !m.stringAt((m.current - 1), 1, "C", "K", "G", "Q", "") {
		m.metaphAdd("K", "K")
	}

	//name sent in 'mac caffrey', 'mac gregor
	if m.stringAt((m.current + 1), 2, " C", " Q", " G", "") {
		m.current += 2
	} else {
		if m.stringAt((m.current+1), 1, "C", "K", "Q", "") && !m.stringAt((m.current+1), 2, "CE", "CI", "") {
			m.current += 2
			// account for combinations such as Ro-ckc-liffe
			if m.stringAt((m.current), 1, "C", "K", "Q", "") && !m.stringAt((m.current+1), 2, "CE", "CI", "") {
				m.current++
			}
		} else {
			m.current++
		}
	}
}

/**
 * Encodes cases where 'C' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_C_At_Beginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(m.current, 2, "CT", "CN", "") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encodes exceptions where "-CA-" should encode to S
 * instead of K including cases where the cedilla has not been used
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CA_To_S() bool {
	// Special case: 'caesar'.
	// Also, where cedilla not used, as in "linguica" => LNKS
	if ((m.current == 0) && m.stringAt(m.current, 4, "CAES", "CAEC", "CAEM", "")) || m.stringAt(0, 8, "FRANCAIS", "FRANCAIX", "LINGUICA", "") || m.stringAt(0, 6, "FACADE", "") || m.stringAt(0, 9, "GONCALVES", "PROVENCAL", "") {
		m.metaphAdd("S", "S")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encodes exceptions where "-CO-" encodes to S instead of K
 * including cases where the cedilla has not been used
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CO_To_S() bool {
	// e.g. 'coelecanth' => SLKN0
	if (m.stringAt(m.current, 4, "COEL", "") && (isVowel(m.charAt(m.current+4)) || ((m.current + 3) == m.last))) || m.stringAt(m.current, 5, "COENA", "COENO", "") || m.stringAt(0, 8, "FRANCOIS", "MELANCON", "") || m.stringAt(0, 6, "GARCON", "") {
		m.metaphAdd("S", "S")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-CH-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CH() bool {
	if m.stringAt(m.current, 2, "CH", "") {
		if m.encode_CHAE() || m.encode_CH_To_H() || m.encode_Silent_CH() || m.encode_ARCH() ||
			// encode_CH_To_X() should be
			// called before the germanic
			// and greek encoding functions
			m.encode_CH_To_X() || m.encode_English_CH_To_K() || m.encode_Germanic_CH_To_K() || m.encode_Greek_CH_Initial() || m.encode_Greek_CH_Non_Initial() {
			return true
		}

		if m.current > 0 {
			if m.stringAt(0, 2, "MC", "") && (m.current == 1) {
				//e.g., "McHugh"
				m.metaphAdd("K", "K")
			} else {
				m.metaphAdd("X", "K")
			}
		} else {
			m.metaphAdd("X", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CHAE-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CHAE() bool {
	// e.g. 'michael'
	if (m.current > 0) && m.stringAt((m.current+2), 2, "AE", "") {
		if m.stringAt(0, 7, "RACHAEL", "") {
			m.metaphAdd("X", "X")
		} else if !m.stringAt((m.current - 1), 1, "C", "K", "G", "Q", "") {
			m.metaphAdd("K", "K")
		}

		m.advanceCounter(4, 2)
		return true
	}

	return false
}

/**
 * Encdoes transliterations from the hebrew where the
 * sound 'kh' is represented as "-CH-". The normal pronounciation
 * of this in english is either 'h' or 'kh', and alternate
 * spellings most often use "-H-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CH_To_H() bool {
	// hebrew => 'H', e.g. 'channukah', 'chabad'
	if ((m.current == 0) && (m.stringAt((m.current+2), 3, "AIM", "ETH", "ELM", "") || m.stringAt((m.current+2), 4, "ASID", "AZAN", "") || m.stringAt((m.current+2), 5, "UPPAH", "UTZPA", "ALLAH", "ALUTZ", "AMETZ", "") || m.stringAt((m.current+2), 6, "ESHVAN", "ADARIM", "ANUKAH", "") || m.stringAt((m.current+2), 7, "ALLLOTH", "ANNUKAH", "AROSETH", ""))) ||
		// and an irish name with the same encoding
		m.stringAt((m.current-3), 7, "CLACHAN", "") {
		m.metaphAdd("H", "H")
		m.advanceCounter(3, 2)
		return true
	}

	return false
}

/**
 * Encodes cases where "-CH-" is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_CH() bool {
	// '-ch-' not pronounced
	if m.stringAt((m.current-2), 7, "FUCHSIA", "") || m.stringAt((m.current-2), 5, "YACHT", "") || m.stringAt(0, 8, "STRACHAN", "") || m.stringAt(0, 8, "CRICHTON", "") || (m.stringAt((m.current-3), 6, "DRACHM", "")) && !m.stringAt((m.current-3), 7, "DRACHMA", "") {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to X
 * English language patterns
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CH_To_X() bool {
	// e.g. 'approach', 'beach'
	if (m.stringAt((m.current-2), 4, "OACH", "EACH", "EECH", "OUCH", "OOCH", "MUCH", "SUCH", "") && !m.stringAt((m.current-3), 5, "JOACH", "")) ||
		// e.g. 'dacha', 'macho'
		(((m.current + 2) == m.last) && m.stringAt((m.current-1), 4, "ACHA", "ACHO", "")) || (m.stringAt(m.current, 4, "CHOT", "CHOD", "CHAT", "") && ((m.current + 3) == m.last)) || ((m.stringAt((m.current-1), 4, "OCHE", "") && ((m.current + 2) == m.last)) && !m.stringAt((m.current-2), 5, "DOCHE", "")) || m.stringAt((m.current-4), 6, "ATTACH", "DETACH", "KOVACH", "") || m.stringAt((m.current-5), 7, "SPINACH", "") || m.stringAt(0, 6, "MACHAU", "") || m.stringAt((m.current-4), 8, "PARACHUT", "") || m.stringAt((m.current-5), 8, "MASSACHU", "") || (m.stringAt((m.current-3), 5, "THACH", "") && !m.stringAt((m.current-1), 4, "ACHE", "")) || m.stringAt((m.current-2), 6, "VACHON", "") {
		m.metaphAdd("X", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to K in contexts of
 * initial "A" or "E" follwed by "CH"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_English_CH_To_K() bool {
	//'ache', 'echo', alternate spelling of 'michael'
	if ((m.current == 1) && rootOrInflections(m.inWord, "ACHE")) || (((m.current > 3) && rootOrInflections(m.inWord[m.current-1:], "ACHE")) && (m.stringAt(0, 3, "EAR", "") || m.stringAt(0, 4, "HEAD", "BACK", "") || m.stringAt(0, 5, "HEART", "BELLY", "TOOTH", ""))) || m.stringAt((m.current-1), 4, "ECHO", "") || m.stringAt((m.current-2), 7, "MICHEAL", "") || m.stringAt((m.current-4), 7, "JERICHO", "") || m.stringAt((m.current-5), 7, "LEPRECH", "") {
		m.metaphAdd("K", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to K in mostly germanic context
 * of internal "-ACH-", with exceptions
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Germanic_CH_To_K() bool {
	// various germanic
	// "<consonant><vowel>CH-"implies a german word where 'ch' => K
	if ((m.current > 1) && !isVowel(m.charAt(m.current-2)) && m.stringAt((m.current-1), 3, "ACH", "") && !m.stringAt((m.current-2), 7, "MACHADO", "MACHUCA", "LACHANC", "LACHAPE", "KACHATU", "") && !m.stringAt((m.current-3), 7, "KHACHAT", "") && ((m.charAt(m.current+2) != 'I') && ((m.charAt(m.current+2) != 'E') || m.stringAt((m.current-2), 6, "BACHER", "MACHER", "MACHEN", "LACHER", ""))) ||
		// e.g. 'brecht', 'fuchs'
		(m.stringAt((m.current+2), 1, "T", "S", "") && !(m.stringAt(0, 11, "WHICHSOEVER", "") || m.stringAt(0, 9, "LUNCHTIME", ""))) ||
		// e.g. 'andromache'
		m.stringAt(0, 4, "SCHR", "") || ((m.current > 2) && m.stringAt((m.current-2), 5, "MACHE", "")) || ((m.current == 2) && m.stringAt((m.current-2), 4, "ZACH", "")) || m.stringAt((m.current-4), 6, "SCHACH", "") || m.stringAt((m.current-1), 5, "ACHEN", "") || m.stringAt((m.current-3), 5, "SPICH", "ZURCH", "BUECH", "") || (m.stringAt((m.current-3), 5, "KIRCH", "JOACH", "BLECH", "MALCH", "") &&
		// "kirch" and "blech" both get 'X'
		!(m.stringAt((m.current-3), 8, "KIRCHNER", "") || ((m.current + 1) == m.last))) || (((m.current + 1) == m.last) && m.stringAt((m.current-2), 4, "NICH", "LICH", "BACH", "")) || (((m.current + 1) == m.last) && m.stringAt((m.current-3), 5, "URICH", "BRICH", "ERICH", "DRICH", "NRICH", "") && !m.stringAt((m.current-5), 7, "ALDRICH", "") && !m.stringAt((m.current-6), 8, "GOODRICH", "") && !m.stringAt((m.current-7), 9, "GINGERICH", ""))) || (((m.current + 1) == m.last) && m.stringAt((m.current-4), 6, "ULRICH", "LFRICH", "LLRICH",
		"EMRICH", "ZURICH", "EYRICH", "")) ||
		// e.g., 'wachtler', 'wechsler', but not 'tichner'
		((m.stringAt((m.current-1), 1, "A", "O", "U", "E", "") || (m.current == 0)) && m.stringAt((m.current+2), 1, "L", "R", "N", "M", "B", "H", "F", "V", "W", " ", "")) {
		// "CHR/L-" e.g. 'chris' do not get
		// alt pronunciation of 'X'
		if m.stringAt((m.current+2), 1, "R", "L", "") || m.slavoGermanic() {
			m.metaphAdd("K", "K")
		} else {
			m.metaphAdd("K", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-ARCH-". Some occurances are from greek roots and therefore encode
 * to 'K', others are from english words and therefore encode to 'X'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ARCH() bool {
	if m.stringAt((m.current - 2), 4, "ARCH", "") {
		// "-ARCH-" has many combining forms where "-CH-" => K because of its
		// derivation from the greek
		if ((isVowel(m.charAt(m.current+2)) && m.stringAt((m.current-2), 5, "ARCHA", "ARCHI", "ARCHO", "ARCHU", "ARCHY", "")) || m.stringAt((m.current-2), 6, "ARCHEA", "ARCHEG", "ARCHEO", "ARCHET", "ARCHEL", "ARCHES", "ARCHEP",
			"ARCHEM", "ARCHEN", "") || (m.stringAt((m.current-2), 4, "ARCH", "") && ((m.current + 1) == m.last)) || m.stringAt(0, 7, "MENARCH", "")) && (!rootOrInflections(m.inWord, "ARCH") && !m.stringAt((m.current-4), 6, "SEARCH", "POARCH", "") && !m.stringAt(0, 9, "ARCHENEMY", "ARCHIBALD", "ARCHULETA", "ARCHAMBAU", "") && !m.stringAt(0, 6, "ARCHER", "ARCHIE", "") && !((((m.stringAt((m.current-3), 5, "LARCH", "MARCH", "PARCH", "") || m.stringAt((m.current-4), 6, "STARCH", "")) && !(m.stringAt(0, 6, "EPARCH", "") || m.stringAt(0, 7, "NOMARCH", "") || m.stringAt(0, 8, "EXILARCH", "HIPPARCH", "MARCHESE", "") || m.stringAt(0, 9, "ARISTARCH", "") || m.stringAt(0, 9, "MARCHETTI", ""))) || rootOrInflections(m.inWord, "STARCH")) && (!m.stringAt((m.current-2), 5, "ARCHU", "ARCHY", "") || m.stringAt(0, 7, "STARCHY", "")))) {
			m.metaphAdd("K", "X")
		} else {
			m.metaphAdd("X", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-CH-" to K when from greek roots
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Greek_CH_Initial() bool {
	// greek roots e.g. 'chemistry', 'chorus', ch at beginning of root
	if (m.stringAt(m.current, 6, "CHAMOM", "CHARAC", "CHARIS", "CHARTO", "CHARTU", "CHARYB", "CHRIST", "CHEMIC", "CHILIA", "") || (m.stringAt(m.current, 5, "CHEMI", "CHEMO", "CHEMU", "CHEMY", "CHOND", "CHONA", "CHONI", "CHOIR", "CHASM",
		"CHARO", "CHROM", "CHROI", "CHAMA", "CHALC", "CHALD", "CHAET", "CHIRO", "CHILO", "CHELA", "CHOUS",
		"CHEIL", "CHEIR", "CHEIM", "CHITI", "CHEOP", "") && !(m.stringAt(m.current, 6, "CHEMIN", "") || m.stringAt((m.current-2), 8, "ANCHONDO", ""))) || (m.stringAt(m.current, 5, "CHISM", "CHELI", "") &&
		// exclude spanish "machismo"
		!(m.stringAt(0, 8, "MACHISMO", "") ||
			// exclude some french words
			m.stringAt(0, 10, "REVANCHISM", "") || m.stringAt(0, 9, "RICHELIEU", "") || (m.stringAt(0, 5, "CHISM", "") && (m.length == 5)) || m.stringAt(0, 6, "MICHEL", ""))) ||
		// include e.g. "chorus", "chyme", "chaos"
		(m.stringAt(m.current, 4, "CHOR", "CHOL", "CHYM", "CHYL", "CHLO", "CHOS", "CHUS", "CHOE", "") && !m.stringAt(0, 6, "CHOLLO", "CHOLLA", "CHORIZ", "")) ||
		// "chaos" => K but not "chao"
		(m.stringAt(m.current, 4, "CHAO", "") && ((m.current + 3) != m.last)) ||
		// e.g. "abranchiate"
		(m.stringAt(m.current, 4, "CHIA", "") && !(m.stringAt(0, 10, "APPALACHIA", "") || m.stringAt(0, 7, "CHIAPAS", ""))) ||
		// e.g. "chimera"
		m.stringAt(m.current, 7, "CHIMERA", "CHIMAER", "CHIMERI", "") ||
		// e.g. "chameleon"
		((m.current == 0) && m.stringAt(m.current, 5, "CHAME", "CHELO", "CHITO", "")) ||
		// e.g. "spirochete"
		((((m.current + 4) == m.last) || ((m.current + 5) == m.last)) && m.stringAt((m.current-1), 6, "OCHETE", ""))) &&
		// more exceptions where "-CH-" => X e.g. "chortle", "crocheter"
		!((m.stringAt(0, 5, "CHORE", "CHOLO", "CHOLA", "") && (m.length == 5)) || m.stringAt(m.current, 5, "CHORT", "CHOSE", "") || m.stringAt((m.current-3), 7, "CROCHET", "") || m.stringAt(0, 7, "CHEMISE", "CHARISE", "CHARISS", "CHAROLE", "")) {
		// "CHR/L-" e.g. 'christ', 'chlorine' do not get
		// alt pronunciation of 'X'
		if m.stringAt((m.current + 2), 1, "R", "L", "") {
			m.metaphAdd("K", "K")
		} else {
			m.metaphAdd("K", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode a variety of greek and some german roots where "-CH-" => K
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Greek_CH_Non_Initial() bool {
	//greek & other roots e.g. 'tachometer', 'orchid', ch in middle or end of root
	if m.stringAt((m.current-2), 6, "ORCHID", "NICHOL", "MECHAN", "LICHEN", "MACHIC", "PACHEL", "RACHIF", "RACHID",
		"RACHIS", "RACHIC", "MICHAL", "") || m.stringAt((m.current-3), 5, "MELCH", "GLOCH", "TRACH", "TROCH", "BRACH", "SYNCH", "PSYCH",
		"STICH", "PULCH", "EPOCH", "") || (m.stringAt((m.current-3), 5, "TRICH", "") && !m.stringAt((m.current-5), 7, "OSTRICH", "")) || (m.stringAt((m.current-2), 4, "TYCH", "TOCH", "BUCH", "MOCH", "CICH", "DICH", "NUCH", "EICH", "LOCH",
		"DOCH", "ZECH", "WYCH", "") && !(m.stringAt((m.current-4), 9, "INDOCHINA", "") || m.stringAt((m.current-2), 6, "BUCHON", ""))) || m.stringAt((m.current-2), 5, "LYCHN", "TACHO", "ORCHO", "ORCHI", "LICHO", "") || (m.stringAt((m.current-1), 5, "OCHER", "ECHIN", "ECHID", "") && ((m.current == 1) || (m.current == 2))) || m.stringAt((m.current-4), 6, "BRONCH", "STOICH", "STRYCH", "TELECH", "PLANCH", "CATECH", "MANICH", "MALACH",
		"BIANCH", "DIDACH", "") || (m.stringAt((m.current-1), 4, "ICHA", "ICHN", "") && (m.current == 1)) || m.stringAt((m.current-2), 8, "ORCHESTR", "") || m.stringAt((m.current-4), 8, "BRANCHIO", "BRANCHIF", "") || (m.stringAt((m.current-1), 5, "ACHAB", "ACHAD", "ACHAN", "ACHAZ", "") && !m.stringAt((m.current-2), 7, "MACHADO", "LACHANC", "")) || m.stringAt((m.current-1), 6, "ACHISH", "ACHILL", "ACHAIA", "ACHENE", "") || m.stringAt((m.current-1), 7, "ACHAIAN", "ACHATES", "ACHIRAL", "ACHERON", "") || m.stringAt((m.current-1), 8, "ACHILLEA", "ACHIMAAS", "ACHILARY", "ACHELOUS", "ACHENIAL", "ACHERNAR", "") || m.stringAt((m.current-1), 9, "ACHALASIA", "ACHILLEAN", "ACHIMENES", "") || m.stringAt((m.current-1), 10, "ACHIMELECH", "ACHITOPHEL", "") ||
		// e.g. 'inchoate'
		(((m.current - 2) == 0) && (m.stringAt((m.current-2), 6, "INCHOA", "") ||
			// e.g. 'ischemia'
			m.stringAt(0, 4, "ISCH", ""))) ||
		// e.g. 'ablimelech', 'antioch', 'pentateuch'
		(((m.current + 1) == m.last) && m.stringAt((m.current-1), 1, "A", "O", "U", "E", "") && !(m.stringAt(0, 7, "DEBAUCH", "") || m.stringAt((m.current-2), 4, "MUCH", "SUCH", "KOCH", "") || m.stringAt((m.current-5), 7, "OODRICH", "ALDRICH", ""))) {
		m.metaphAdd("K", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes reliably italian "-CCIA-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CCIA() bool {
	//e.g., 'focaccia'
	if m.stringAt((m.current + 1), 3, "CIA", "") {
		m.metaphAdd("X", "S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-CC-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CC() bool {
	//double 'C', but not if e.g. 'McClellan'
	if m.stringAt(m.current, 2, "CC", "") && !((m.current == 1) && (m.charAt(0) == 'M')) {
		// exception
		if m.stringAt((m.current - 3), 7, "FLACCID", "") {
			m.metaphAdd("S", "S")
			m.advanceCounter(3, 2)
			return true
		}

		//'bacci', 'bertucci', other italian
		if (((m.current + 2) == m.last) && m.stringAt((m.current+2), 1, "I", "")) || m.stringAt((m.current+2), 2, "IO", "") || (((m.current + 4) == m.last) && m.stringAt((m.current+2), 3, "INO", "INI", "")) {
			m.metaphAdd("X", "X")
			m.advanceCounter(3, 2)
			return true
		}

		//'accident', 'accede' 'succeed'
		if m.stringAt((m.current+2), 1, "I", "E", "Y", "") &&
			//except 'bellocchio','bacchus', 'soccer' get K
			!((m.charAt(m.current+2) == 'H') || m.stringAt((m.current-2), 6, "SOCCER", "")) {
			m.metaphAdd("KS", "KS")
			m.advanceCounter(3, 2)
			return true

		} else {
			//Pierce's rule
			m.metaphAdd("K", "K")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode cases where the consonant following "C" is redundant
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CK_CG_CQ() bool {
	if m.stringAt(m.current, 2, "CK", "CG", "CQ", "") {
		// eastern european spelling e.g. 'gorecki' == 'goresky'
		if m.stringAt(m.current, 3, "CKI", "CKY", "") && ((m.current + 2) == m.last) && (m.length > 6) {
			m.metaphAdd("K", "SK")
		} else {
			m.metaphAdd("K", "K")
		}
		m.current += 2

		if m.stringAt(m.current, 1, "K", "G", "Q", "") {
			m.current++
		}
		return true
	}

	return false
}

/**
 * Encode cases where "C" preceeds a front vowel such as "E", "I", or "Y".
 * These cases most likely => S or X
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_C_Front_Vowel() bool {
	if m.stringAt(m.current, 2, "CI", "CE", "CY", "") {
		if m.encode_British_Silent_CE() || m.encode_CE() || m.encode_CI() || m.encode_Latinate_Suffixes() {
			m.advanceCounter(2, 1)
			return true
		}

		m.metaphAdd("S", "S")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_British_Silent_CE() bool {
	// english place names like e.g.'gloucester' pronounced glo-ster
	if (m.stringAt((m.current+1), 5, "ESTER", "") && ((m.current + 5) == m.last)) || m.stringAt((m.current+1), 10, "ESTERSHIRE", "") {
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CE() bool {
	// 'ocean', 'commercial', 'provincial', 'cello', 'fettucini', 'medici'
	if (m.stringAt((m.current+1), 3, "EAN", "") && isVowel(m.charAt(m.current-1))) ||
		// e.g. 'rosacea'
		(m.stringAt((m.current-1), 4, "ACEA", "") && ((m.current + 2) == m.last) && !m.stringAt(0, 7, "PANACEA", "")) ||
		// e.g. 'botticelli', 'concerto'
		m.stringAt((m.current+1), 4, "ELLI", "ERTO", "EORL", "") ||
		// some italian names familiar to americans
		(m.stringAt((m.current-3), 5, "CROCE", "") && ((m.current + 1) == m.last)) || m.stringAt((m.current-3), 5, "DOLCE", "") ||
		// e.g. 'cello'
		(m.stringAt((m.current+1), 4, "ELLO", "") && ((m.current + 4) == m.last)) {
		m.metaphAdd("X", "S")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CI() bool {
	// with consonant before C
	// e.g. 'fettucini', but exception for the americanized pronunciation of 'mancini'
	if ((m.stringAt((m.current+1), 3, "INI", "") && !m.stringAt(0, 7, "MANCINI", "")) && ((m.current + 3) == m.last)) ||
		// e.g. 'medici'
		(m.stringAt((m.current-1), 3, "ICI", "") && ((m.current + 1) == m.last)) ||
		// e.g. "commercial', 'provincial', 'cistercian'
		m.stringAt((m.current-1), 5, "RCIAL", "NCIAL", "RCIAN", "UCIUS", "") ||
		// special cases
		m.stringAt((m.current-3), 6, "MARCIA", "") || m.stringAt((m.current-2), 7, "ANCIENT", "") {
		m.metaphAdd("X", "S")
		return true
	}

	// with vowel before C (or at beginning?)
	if ((m.stringAt(m.current, 3, "CIO", "CIE", "CIA", "") && isVowel(m.charAt(m.current-1))) ||
		// e.g. "ciao"
		m.stringAt((m.current+1), 3, "IAO", "")) && !m.stringAt((m.current-4), 8, "COERCION", "") {
		if (m.stringAt(m.current, 4, "CIAN", "CIAL", "CIAO", "CIES", "CIOL", "CION", "") ||
			// exception - "glacier" => 'X' but "spacier" = > 'S'
			m.stringAt((m.current-3), 7, "GLACIER", "") || m.stringAt(m.current, 5, "CIENT", "CIENC", "CIOUS", "CIATE", "CIATI", "CIATO", "CIABL", "CIARY", "") || (((m.current + 2) == m.last) && m.stringAt(m.current, 3, "CIA", "CIO", "")) || (((m.current + 3) == m.last) && m.stringAt(m.current, 3, "CIAS", "CIOS", ""))) &&
			// exceptions
			!(m.stringAt((m.current-4), 11, "ASSOCIATION", "") || m.stringAt(0, 4, "OCIE", "") ||
				// exceptions mostly because these names are usually from
				// the spanish rather than the italian in america
				m.stringAt((m.current-2), 5, "LUCIO", "") || m.stringAt((m.current-2), 6, "MACIAS", "") || m.stringAt((m.current-3), 6, "GRACIE", "GRACIA", "") || m.stringAt((m.current-2), 7, "LUCIANO", "") || m.stringAt((m.current-3), 8, "MARCIANO", "") || m.stringAt((m.current-4), 7, "PALACIO", "") || m.stringAt((m.current-4), 9, "FELICIANO", "") || m.stringAt((m.current-5), 8, "MAURICIO", "") || m.stringAt((m.current-7), 11, "ENCARNACION", "") || m.stringAt((m.current-4), 8, "POLICIES", "") || m.stringAt((m.current-2), 8, "HACIENDA", "") || m.stringAt((m.current-6), 9, "ANDALUCIA", "") || m.stringAt((m.current-2), 5, "SOCIO", "SOCIE", "")) {
			m.metaphAdd("X", "S")
		} else {
			m.metaphAdd("S", "X")
		}

		return true
	}

	// exception
	if m.stringAt((m.current - 4), 8, "COERCION", "") {
		m.metaphAdd("J", "J")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Latinate_Suffixes() bool {
	if m.stringAt((m.current + 1), 4, "EOUS", "IOUS", "") {
		m.metaphAdd("X", "S")
		return true
	}

	return false
}

/**
 * Encodes some exceptions where "C" is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_C() bool {
	if m.stringAt((m.current + 1), 1, "T", "S", "") {
		if m.stringAt(0, 11, "CONNECTICUT", "") || m.stringAt(0, 6, "INDICT", "TUCSON", "") {
			m.current++
			return true
		}
	}

	return false
}

/**
 * Encodes slavic spellings or transliterations
 * written as "-CZ-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CZ() bool {
	if m.stringAt((m.current+1), 1, "Z", "") && !m.stringAt((m.current-1), 6, "ECZEMA", "") {
		if m.stringAt(m.current, 4, "CZAR", "") {
			m.metaphAdd("S", "S")
		} else {
			// otherwise most likely a czech word...
			m.metaphAdd("X", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * "-CS" special cases
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_CS() bool {
	// give an 'etymological' 2nd
	// encoding for "kovacs" so
	// that it matches "kovach"
	if m.stringAt(0, 6, "KOVACS", "") {
		m.metaphAdd("KS", "X")
		m.current += 2
		return true
	}

	if m.stringAt((m.current-1), 3, "ACS", "") && ((m.current + 1) == m.last) && !m.stringAt((m.current-4), 6, "ISAACS", "") {
		m.metaphAdd("X", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-D-"
 *
 */
func (m *M3) encode_D() {
	if m.encode_DG() || m.encode_DJ() || m.encode_DT_DD() || m.encode_D_To_J() || m.encode_DOUS() || m.encode_Silent_D() {
		return
	}

	if m.encodeExact {
		// "final de-voicing" in this case
		// e.g. 'missed' == 'mist'
		if (m.current == m.last) && m.stringAt((m.current-3), 4, "SSED", "") {
			m.metaphAdd("T", "T")
		} else {
			m.metaphAdd("D", "D")
		}
	} else {
		m.metaphAdd("T", "T")
	}
	m.current++
}

/**
 * Encode "-DG-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_DG() bool {
	if m.stringAt(m.current, 2, "DG", "") {
		// excludes exceptions e.g. 'edgar',
		// or cases where 'g' is first letter of combining form
		// e.g. 'handgun', 'waldglas'
		if m.stringAt((m.current+2), 1, "A", "O", "") ||
			// e.g. "midgut"
			m.stringAt((m.current+1), 3, "GUN", "GUT", "") ||
			// e.g. "handgrip"
			m.stringAt((m.current+1), 4, "GEAR", "GLAS", "GRIP", "GREN", "GILL", "GRAF", "") ||
			// e.g. "mudgard"
			m.stringAt((m.current+1), 5, "GUARD", "GUILT", "GRAVE", "GRASS", "") ||
			// e.g. "woodgrouse"
			m.stringAt((m.current+1), 6, "GROUSE", "") {
			m.metaphAddExactApprox("DG", "TK")
		} else {
			//e.g. "edge", "abridgment"
			m.metaphAdd("J", "J")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-DJ-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_DJ() bool {
	// e.g. "adjacent"
	if m.stringAt(m.current, 2, "DJ", "") {
		m.metaphAdd("J", "J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-DD-" and "-DT-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_DT_DD() bool {
	// eat redundant 'T' or 'D'
	if m.stringAt(m.current, 2, "DT", "DD", "") {
		if m.stringAt(m.current, 3, "DTH", "") {
			m.metaphAddExactApprox("D0", "T0")
			m.current += 3
		} else {
			if m.encodeExact {
				// devoice it
				if m.stringAt(m.current, 2, "DT", "") {
					m.metaphAdd("T", "T")
				} else {
					m.metaphAdd("D", "D")
				}
			} else {
				m.metaphAdd("T", "T")
			}
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode cases where "-DU-" "-DI-", and "-DI-" => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_D_To_J() bool {
	// e.g. "module", "adulate"
	if (m.stringAt(m.current, 3, "DUL", "") && (isVowel(m.charAt(m.current-1)) && isVowel(m.charAt(m.current+3)))) ||
		// e.g. "soldier", "grandeur", "procedure"
		(((m.current + 3) == m.last) && m.stringAt((m.current-1), 5, "LDIER", "NDEUR", "EDURE", "RDURE", "")) || m.stringAt((m.current-3), 7, "CORDIAL", "") ||
		// e.g.  "pendulum", "education"
		m.stringAt((m.current-1), 5, "NDULA", "NDULU", "EDUCA", "") ||
		// e.g. "individual", "individual", "residuum"
		m.stringAt((m.current-1), 4, "ADUA", "IDUA", "IDUU", "") {
		m.metaphAddExactApprox4("J", "D", "J", "T")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode latinate suffix "-DOUS" where 'D' is pronounced as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_DOUS() bool {
	// e.g. "assiduous", "arduous"
	if m.stringAt((m.current + 1), 4, "UOUS", "") {
		m.metaphAddExactApprox4("J", "D", "J", "T")
		m.advanceCounter(4, 1)
		return true
	}

	return false
}

/**
 * Encode silent "-D-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_D() bool {
	// silent 'D' e.g. 'wednesday', 'handsome'
	if m.stringAt((m.current-2), 9, "WEDNESDAY", "") || m.stringAt((m.current-3), 7, "HANDKER", "HANDSOM", "WINDSOR", "") ||
		// french silent D at end in words or names familiar to americans
		m.stringAt((m.current-5), 6, "PERNOD", "ARTAUD", "RENAUD", "") || m.stringAt((m.current-6), 7, "RIMBAUD", "MICHAUD", "BICHAUD", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-F-"
 *
 */
func (m *M3) encode_F() {
	// Encode cases where "-FT-" => "T" is usually silent
	// e.g. 'often', 'soften'
	// This should really be covered under "T"!
	if m.stringAt((m.current - 1), 5, "OFTEN", "") {
		m.metaphAdd("F", "FT")
		m.current += 2
		return
	}

	// eat redundant 'F'
	if m.charAt(m.current+1) == 'F' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("F", "F")

}

/**
 * Encode "-G-"
 *
 */
func (m *M3) encode_G() {
	if m.encode_Silent_G_At_Beginning() || m.encode_GG() || m.encode_GK() || m.encode_GH() || m.encode_Silent_G() || m.encode_GN() || m.encode_GL() || m.encode_Initial_G_Front_Vowel() || m.encode_NGER() || m.encode_GER() || m.encode_GEL() || m.encode_Non_Initial_G_Front_Vowel() || m.encode_GA_To_J() {
		return
	}

	if !m.stringAt((m.current - 1), 1, "C", "K", "G", "Q", "") {
		m.metaphAddExactApprox("G", "K")
	}

	m.current++
}

/**
 * Encode cases where 'G' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_G_At_Beginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(m.current, 2, "GN", "") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode "-GG-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GG() bool {
	if m.charAt(m.current+1) == 'G' {
		// italian e.g, 'loggia', 'caraveggio', also 'suggest' and 'exaggerate'
		if m.stringAt((m.current-1), 5, "AGGIA", "OGGIA", "AGGIO", "EGGIO", "EGGIA", "IGGIO", "") ||
			// 'ruggiero' but not 'snuggies'
			(m.stringAt((m.current-1), 5, "UGGIE", "") && !(((m.current + 3) == m.last) || ((m.current + 4) == m.last))) || (((m.current + 2) == m.last) && m.stringAt((m.current-1), 4, "AGGI", "OGGI", "")) || m.stringAt((m.current-2), 6, "SUGGES", "XAGGER", "REGGIE", "") {
			// expection where "-GG-" => KJ
			if m.stringAt((m.current - 2), 7, "SUGGEST", "") {
				m.metaphAddExactApprox("G", "K")
			}

			m.metaphAdd("J", "J")
			m.advanceCounter(3, 2)
		} else {
			m.metaphAddExactApprox("G", "K")
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode "-GK-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GK() bool {
	// 'gingko'
	if m.charAt(m.current+1) == 'K' {
		m.metaphAdd("K", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-GH-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH() bool {
	if m.charAt(m.current+1) == 'H' {
		if m.encode_GH_After_Consonant() || m.encode_Initial_GH() || m.encode_GH_To_J() || m.encode_GH_To_H() || m.encode_UGHT() || m.encode_GH_H_Part_Of_Other_Word() || m.encode_Silent_GH() || m.encode_GH_To_F() {
			return true
		}

		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_After_Consonant() bool {
	// e.g. 'burgher', 'bingham'
	if (m.current > 0) && !isVowel(m.charAt(m.current-1)) &&
		// not e.g. 'greenhalgh'
		!(m.stringAt((m.current-3), 5, "HALGH", "") && ((m.current + 1) == m.last)) {
		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_GH() bool {
	if m.current < 3 {
		// e.g. "ghislane", "ghiradelli"
		if m.current == 0 {
			if m.charAt(m.current+2) == 'I' {
				m.metaphAdd("J", "J")
			} else {
				m.metaphAddExactApprox("G", "K")
			}
			m.current += 2
			return true
		}
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_To_J() bool {
	// e.g., 'greenhalgh', 'dunkenhalgh', english names
	if m.stringAt((m.current-2), 4, "ALGH", "") && ((m.current + 1) == m.last) {
		m.metaphAdd("J", "")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_To_H() bool {
	// special cases
	// e.g., 'donoghue', 'donaghy'
	if (m.stringAt((m.current-4), 4, "DONO", "DONA", "") && isVowel(m.charAt(m.current+2))) || m.stringAt((m.current-5), 9, "CALLAGHAN", "") {
		m.metaphAdd("H", "H")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_UGHT() bool {
	//e.g. "ought", "aught", "daughter", "slaughter"
	if m.stringAt((m.current - 1), 4, "UGHT", "") {
		if (m.stringAt((m.current-3), 5, "LAUGH", "") && !(m.stringAt((m.current-4), 7, "SLAUGHT", "") || m.stringAt((m.current-3), 7, "LAUGHTO", ""))) || m.stringAt((m.current-4), 6, "DRAUGH", "") {
			m.metaphAdd("FT", "FT")
		} else {
			m.metaphAdd("T", "T")
		}
		m.current += 3
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_H_Part_Of_Other_Word() bool {
	// if the 'H' is the beginning of another word or syllable
	if m.stringAt((m.current + 1), 4, "HOUS", "HEAD", "HOLE", "HORN", "HARN", "") {
		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_GH() bool {
	//Parker's rule (with some further refinements) - e.g., 'hugh'
	if ((((m.current > 1) && m.stringAt((m.current-2), 1, "B", "H", "D", "G", "L", "")) ||
		//e.g., 'bough'
		((m.current > 2) && m.stringAt((m.current-3), 1, "B", "H", "D", "K", "W", "N", "P", "V", "") && !m.stringAt(0, 6, "ENOUGH", "")) ||
		//e.g., 'broughton'
		((m.current > 3) && m.stringAt((m.current-4), 1, "B", "H", "")) ||
		//'plough', 'slaugh'
		((m.current > 3) && m.stringAt((m.current-4), 2, "PL", "SL", "")) || ((m.current > 0) &&
		// 'sigh', 'light'
		((m.charAt(m.current-1) == 'I') || m.stringAt(0, 4, "PUGH", "") ||
			// e.g. 'MCDONAGH', 'MURTAGH', 'CREAGH'
			(m.stringAt((m.current-1), 3, "AGH", "") && ((m.current + 1) == m.last)) || m.stringAt((m.current-4), 6, "GERAGH", "DRAUGH", "") || (m.stringAt((m.current-3), 5, "GAUGH", "GEOGH", "MAUGH", "") && !m.stringAt(0, 9, "MCGAUGHEY", "")) ||
			// exceptions to 'tough', 'rough', 'lough'
			(m.stringAt((m.current-2), 4, "OUGH", "") && (m.current > 3) && !m.stringAt((m.current-4), 6, "CCOUGH", "ENOUGH", "TROUGH", "CLOUGH", ""))))) &&
		// suffixes starting w/ vowel where "-GH-" is usually silent
		(m.stringAt((m.current-3), 5, "VAUGH", "FEIGH", "LEIGH", "") ||
			m.stringAt((m.current-2), 4, "HIGH", "TIGH", "") || ((m.current + 1) == m.last) || (m.stringAt((m.current+2), 2, "IE", "EY", "ES", "ER", "ED", "TY", "") && ((m.current + 3) == m.last) && !m.stringAt((m.current-5), 9, "GALLAGHER", "")) || (m.stringAt((m.current+2), 1, "Y", "") && ((m.current + 2) == m.last)) || (m.stringAt((m.current+2), 3, "ING", "OUT", "") && ((m.current + 4) == m.last)) || (m.stringAt((m.current+2), 4, "ERTY", "") && ((m.current + 5) == m.last)) || (!isVowel(m.charAt(m.current+2)) || m.stringAt((m.current-3), 5, "GAUGH", "GEOGH", "MAUGH", "") || m.stringAt((m.current-4), 8, "BROUGHAM", "")))) &&
		// exceptions where '-g-' pronounced
		!(m.stringAt(0, 6, "BALOGH", "SABAGH", "") ||
			m.stringAt((m.current-2), 7, "BAGHDAD", "") || m.stringAt((m.current-3), 5, "WHIGH", "") || m.stringAt((m.current-5), 7, "SABBAGH", "AKHLAGH", "")) {
		// silent - do nothing
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_Special_Cases() bool {
	handled := false

	// special case: 'hiccough' == 'hiccup'
	if m.stringAt((m.current - 6), 8, "HICCOUGH", "") {
		m.metaphAdd("P", "P")
		handled = true
	} else
	// special case: 'lough' alt spelling for scots 'loch'
	if m.stringAt(0, 5, "LOUGH", "") {
		m.metaphAdd("K", "K")
		handled = true
	} else
	// hungarian
	if m.stringAt(0, 6, "BALOGH", "") {
		m.metaphAddExactApprox4("G", "", "K", "")
		handled = true
	} else
	// "maclaughlin"
	if m.stringAt((m.current - 3), 8, "LAUGHLIN", "COUGHLAN", "LOUGHLIN", "") {
		m.metaphAdd("K", "F")
		handled = true
	} else if m.stringAt((m.current-3), 5, "GOUGH", "") || m.stringAt((m.current-7), 9, "COLCLOUGH", "") {
		m.metaphAdd("", "F")
		handled = true
	}

	if handled {
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GH_To_F() bool {
	// the cases covered here would fall under
	// the GH_To_F rule below otherwise
	if m.encode_GH_Special_Cases() {
		return true
	} else {
		//e.g., 'laugh', 'cough', 'rough', 'tough'
		if (m.current > 2) && (m.charAt(m.current-1) == 'U') && isVowel(m.charAt(m.current-2)) && m.stringAt((m.current-3), 1, "C", "G", "L", "R", "T", "N", "S", "") && !m.stringAt((m.current-4), 8, "BREUGHEL", "FLAUGHER", "") {
			m.metaphAdd("F", "F")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode some contexts where "g" is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_G() bool {
	// e.g. "phlegm", "apothegm", "voigt"
	if (((m.current + 1) == m.last) && (m.stringAt((m.current-1), 3, "EGM", "IGM", "AGM", "") || m.stringAt(m.current, 2, "GT", ""))) || (m.stringAt(0, 5, "HUGES", "") && (m.length == 5)) {
		m.current++
		return true
	}

	// vietnamese names e.g. "Nguyen" but not "Ng"
	if m.stringAt(0, 2, "NG", "") && (m.current != m.last) {
		m.current++
		return true
	}

	return false
}

/**
 * ENcode "-GN-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GN() bool {
	if m.charAt(m.current+1) == 'N' {
		// 'align' 'sign', 'resign' but not 'resignation'
		// also 'impugn', 'impugnable', but not 'repugnant'
		if ((m.current > 1) && ((m.stringAt((m.current-1), 1, "I", "U", "E", "") || m.stringAt((m.current-3), 9, "LORGNETTE", "") || m.stringAt((m.current-2), 9, "LAGNIAPPE", "") || m.stringAt((m.current-2), 6, "COGNAC", "") || m.stringAt((m.current-3), 7, "CHAGNON", "") || m.stringAt((m.current-5), 9, "COMPAGNIE", "") || m.stringAt((m.current-4), 6, "BOLOGN", "")) &&
			// Exceptions: following are cases where 'G' is pronounced
			// in "assign" 'g' is silent, but not in "assignation"
			!(m.stringAt((m.current+2), 5, "ATION", "") || m.stringAt((m.current+2), 4, "ATOR", "") || m.stringAt((m.current+2), 3, "ATE", "ITY", "") ||
				// exception to exceptions, not pronounced:
				(m.stringAt((m.current+2), 2, "AN", "AC", "IA", "UM", "") && !(m.stringAt((m.current-3), 8, "POIGNANT", "") || m.stringAt((m.current-2), 6, "COGNAC", ""))) || m.stringAt(0, 7, "SPIGNER", "STEGNER", "") || (m.stringAt(0, 5, "SIGNE", "") && (m.length == 5)) || m.stringAt((m.current-2), 5, "LIGNI", "LIGNO", "REGNA", "DIGNI", "WEGNE",
				"TIGNE", "RIGNE", "REGNE", "TIGNO", "") || m.stringAt((m.current-2), 6, "SIGNAL", "SIGNIF", "SIGNAT", "") || m.stringAt((m.current-1), 5, "IGNIT", "")) && !m.stringAt((m.current-2), 6, "SIGNET", "LIGNEO", ""))) ||
			//not e.g. 'cagney', 'magna'
			(((m.current + 2) == m.last) && m.stringAt(m.current, 3, "GNE", "GNA", "") && !m.stringAt((m.current-2), 5, "SIGNA", "MAGNA", "SIGNE", "")) {
			m.metaphAddExactApprox4("N", "GN", "N", "KN")
		} else {
			m.metaphAddExactApprox("GN", "KN")
		}
		m.current += 2
		return true
	}
	return false
}

/**
 * Encode "-GL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GL() bool {
	//'tagliaro', 'puglia' BUT add K in alternative
	// since americans sometimes do this
	if m.stringAt((m.current+1), 3, "LIA", "LIO", "LIE", "") && isVowel(m.charAt(m.current-1)) {
		m.metaphAddExactApprox4("L", "GL", "L", "KL")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) initial_G_Soft() bool {
	if ((m.stringAt((m.current+1), 2, "EL", "EM", "EN", "EO", "ER", "ES", "IA", "IN", "IO", "IP", "IU", "YM", "YN", "YP", "YR", "EE", "") || m.stringAt((m.current+1), 3, "IRA", "IRO", "")) &&
		// except for smaller set of cases where => K, e.g. "gerber"
		!(m.stringAt((m.current+1), 3, "ELD", "ELT", "ERT", "INZ", "ERH", "ITE", "ERD", "ERL", "ERN",
			"INT", "EES", "EEK", "ELB", "EER", "") || m.stringAt((m.current+1), 4, "ERSH", "ERST", "INSB", "INGR", "EROW", "ERKE", "EREN", "") || m.stringAt((m.current+1), 5, "ELLER", "ERDIE", "ERBER", "ESUND", "ESNER", "INGKO", "INKGO",
			"IPPER", "ESELL", "IPSON", "EEZER", "ERSON", "ELMAN", "") || m.stringAt((m.current+1), 6, "ESTALT", "ESTAPO", "INGHAM", "ERRITY", "ERRISH", "ESSNER", "ENGLER", "") || m.stringAt((m.current+1), 7, "YNAECOL", "YNECOLO", "ENTHNER", "ERAGHTY", "") || m.stringAt((m.current+1), 8, "INGERICH", "EOGHEGAN", ""))) || (isVowel(m.charAt(m.current+1)) && (m.stringAt((m.current+1), 3, "EE ", "EEW", "") || (m.stringAt((m.current+1), 3, "IGI", "IRA", "IBE", "AOL", "IDE", "IGL", "") && !m.stringAt((m.current+1), 5, "IDEON", "")) || m.stringAt((m.current+1), 4, "ILES", "INGI", "ISEL", "") || (m.stringAt((m.current+1), 5, "INGER", "") && !m.stringAt((m.current+1), 8, "INGERICH", "")) || m.stringAt((m.current+1), 5, "IBBER", "IBBET", "IBLET", "IBRAN", "IGOLO", "IRARD", "IGANT", "") || m.stringAt((m.current+1), 6, "IRAFFE", "EEWHIZ", "") || m.stringAt((m.current+1), 7, "ILLETTE", "IBRALTA", ""))) {
		return true
	}

	return false
}

/**
 * Encode cases where 'G' is at start of word followed
 * by a "front" vowel e.g. 'E', 'I', 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_G_Front_Vowel() bool {
	// 'g' followed by vowel at beginning
	if (m.current == 0) && m.front_Vowel(m.current+1) {
		// special case "gila" as in "gila monster"
		if m.stringAt((m.current+1), 3, "ILA", "") && (m.length == 4) {
			m.metaphAdd("H", "H")
		} else if m.initial_G_Soft() {
			m.metaphAddExactApprox4("J", "G", "J", "K")
		} else {
			// only code alternate 'J' if front vowel
			if (m.charAt(m.current+1) == 'E') || (m.charAt(m.current+1) == 'I') {
				m.metaphAddExactApprox4("G", "J", "K", "J")
			} else {
				m.metaphAddExactApprox("G", "K")
			}
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-NGER-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_NGER() bool {
	if (m.current > 1) && m.stringAt((m.current-1), 4, "NGER", "") {
		// default 'G' => J  such as 'ranger', 'stranger', 'manger', 'messenger', 'orangery', 'granger'
		// 'boulanger', 'challenger', 'danger', 'changer', 'harbinger', 'lounger', 'ginger', 'passenger'
		// except for these the following
		if !(rootOrInflections(m.inWord, "ANGER") || rootOrInflections(m.inWord, "LINGER") || rootOrInflections(m.inWord, "MALINGER") || rootOrInflections(m.inWord, "FINGER") || (m.stringAt((m.current-3), 4, "HUNG", "FING", "BUNG", "WING", "RING", "DING", "ZENG",
			"ZING", "JUNG", "LONG", "PING", "CONG", "MONG", "BANG",
			"GANG", "HANG", "LANG", "SANG", "SING", "WANG", "ZANG", "") &&
			// exceptions to above where 'G' => J
			!(m.stringAt((m.current-6), 7, "BOULANG", "SLESING", "KISSING", "DERRING", "") || m.stringAt((m.current-8), 9, "SCHLESING", "") || m.stringAt((m.current-5), 6, "SALING", "BELANG", "") || m.stringAt((m.current-6), 7, "BARRING", "") || m.stringAt((m.current-6), 9, "PHALANGER", "") || m.stringAt((m.current-4), 5, "CHANG", ""))) || m.stringAt((m.current-4), 5, "STING", "YOUNG", "") || m.stringAt((m.current-5), 6, "STRONG", "") || m.stringAt(0, 3, "UNG", "ENG", "ING", "") || m.stringAt(m.current, 6, "GERICH", "") || m.stringAt(0, 6, "SENGER", "") || m.stringAt((m.current-3), 6, "WENGER", "MUNGER", "SONGER", "KINGER", "") || m.stringAt((m.current-4), 7, "FLINGER", "SLINGER", "STANGER", "STENGER", "KLINGER", "CLINGER", "") || m.stringAt((m.current-5), 8, "SPRINGER", "SPRENGER", "") || m.stringAt((m.current-3), 7, "LINGERF", "") || m.stringAt((m.current-2), 7, "ANGERLY", "ANGERBO", "INGERSO", "")) {
			m.metaphAddExactApprox4("J", "G", "J", "K")
		} else {
			m.metaphAddExactApprox4("G", "J", "K", "J")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-GER-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GER() bool {
	if (m.current > 0) && m.stringAt((m.current+1), 2, "ER", "") {
		// Exceptions to 'GE' where 'G' => K
		// e.g. "JAGER", "TIGER", "LIGER", "LAGER", "LUGER", "AUGER", "EAGER", "HAGER", "SAGER"
		if (((m.current == 2) && isVowel(m.charAt(m.current-1)) && !isVowel(m.charAt(m.current-2)) && !(m.stringAt((m.current-2), 5, "PAGER", "WAGER", "NIGER", "ROGER", "LEGER", "CAGER", "")) || m.stringAt((m.current-2), 5, "AUGER", "EAGER", "INGER", "YAGER", "")) || m.stringAt((m.current-3), 6, "SEEGER", "JAEGER", "GEIGER", "KRUGER", "SAUGER", "BURGER",
			"MEAGER", "MARGER", "RIEGER", "YAEGER", "STEGER", "PRAGER", "SWIGER",
			"YERGER", "TORGER", "FERGER", "HILGER", "ZEIGER", "YARGER",
			"COWGER", "CREGER", "KROGER", "KREGER", "GRAGER", "STIGER", "BERGER", "") ||
			// 'berger' but not 'bergerac'
			(m.stringAt((m.current-3), 6, "BERGER", "") && ((m.current + 2) == m.last)) || m.stringAt((m.current-4), 7, "KREIGER", "KRUEGER", "METZGER", "KRIEGER", "KROEGER", "STEIGER",
			"DRAEGER", "BUERGER", "BOERGER", "FIBIGER", "") ||
			// e.g. 'harshbarger', 'winebarger'
			(m.stringAt((m.current-3), 6, "BARGER", "") && (m.current > 4)) ||
			// e.g. 'weisgerber'
			(m.stringAt(m.current, 6, "GERBER", "") && (m.current > 0)) || m.stringAt((m.current-5), 8, "SCHWAGER", "LYBARGER", "SPRENGER", "GALLAGER", "WILLIGER", "") || m.stringAt(0, 4, "HARGER", "") || (m.stringAt(0, 4, "AGER", "EGER", "") && (m.length == 4)) || m.stringAt((m.current-1), 6, "YGERNE", "") || m.stringAt((m.current-6), 9, "SCHWEIGER", "")) && !(m.stringAt((m.current-5), 10, "BELLIGEREN", "") || m.stringAt(0, 7, "MARGERY", "") || m.stringAt((m.current-3), 8, "BERGERAC", "")) {
			if m.slavoGermanic() {
				m.metaphAddExactApprox("G", "K")
			} else {
				m.metaphAddExactApprox4("G", "J", "K", "J")
			}
		} else {
			m.metaphAddExactApprox4("J", "G", "J", "K")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * ENcode "-GEL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GEL() bool {
	// more likely to be "-GEL-" => JL
	if m.stringAt((m.current+1), 2, "EL", "") && (m.current > 0) {
		// except for
		// "BAGEL", "HEGEL", "HUGEL", "KUGEL", "NAGEL", "VOGEL", "FOGEL", "PAGEL"
		if ((m.length == 5) && isVowel(m.charAt(m.current-1)) && !isVowel(m.charAt(m.current-2)) && !m.stringAt((m.current-2), 5, "NIGEL", "RIGEL", "")) ||
			// or the following as combining forms
			m.stringAt((m.current-2), 5, "ENGEL", "HEGEL", "NAGEL", "VOGEL", "") || m.stringAt((m.current-3), 6, "MANGEL", "WEIGEL", "FLUGEL", "RANGEL", "HAUGEN", "RIEGEL", "VOEGEL", "") || m.stringAt((m.current-4), 7, "SPEIGEL", "STEIGEL", "WRANGEL", "SPIEGEL", "") || m.stringAt((m.current-4), 8, "DANEGELD", "") {
			if m.slavoGermanic() {
				m.metaphAddExactApprox("G", "K")
			} else {
				m.metaphAddExactApprox4("G", "J", "K", "J")
			}
		} else {
			m.metaphAddExactApprox4("J", "G", "J", "K")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-G-" followed by a vowel when non-initial leter.
 * Default for this is a 'J' sound, so check exceptions where
 * it is pronounced 'G'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Non_Initial_G_Front_Vowel() bool {
	// -gy-, gi-, ge-
	if m.stringAt((m.current + 1), 1, "E", "I", "Y", "") {
		// '-ge' at end
		// almost always 'j 'sound
		if m.stringAt(m.current, 2, "GE", "") && (m.current == (m.last - 1)) {
			if m.hard_GE_At_End() {
				if m.slavoGermanic() {
					m.metaphAddExactApprox("G", "K")
				} else {
					m.metaphAddExactApprox4("G", "J", "K", "J")
				}
			} else {
				m.metaphAdd("J", "J")
			}
		} else {
			if m.internal_Hard_G() {
				// don't encode KG or KK if e.g. "mcgill"
				if !((m.current == 2) && m.stringAt(0, 2, "MC", "")) || ((m.current == 3) && m.stringAt(0, 3, "MAC", "")) {
					if m.slavoGermanic() {
						m.metaphAddExactApprox("G", "K")
					} else {
						m.metaphAddExactApprox4("G", "J", "K", "J")
					}
				}
			} else {
				m.metaphAddExactApprox4("J", "G", "J", "K")
			}
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/*
 * Detect german names and other words that have
 * a 'hard' 'g' in the context of "-ge" at end
 *
 * @return true if encoding handled in this routine, false if not
 */
func (m *M3) hard_GE_At_End() bool {
	if m.stringAt(0, 6, "RENEGE", "STONGE", "STANGE", "PRANGE", "KRESGE", "") || m.stringAt(0, 5, "BYRGE", "BIRGE", "BERGE", "HAUGE", "") || m.stringAt(0, 4, "HAGE", "") || m.stringAt(0, 5, "LANGE", "SYNGE", "BENGE", "RUNGE", "HELGE", "") || m.stringAt(0, 4, "INGE", "LAGE", "") {
		return true
	}

	return false
}

/**
 * Exceptions to default encoding to 'J':
 * encode "-G-" to 'G' in "-g<frontvowel>-" words
 * where we are not at "-GE" at the end of the word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) internal_Hard_G() bool {
	// if not "-GE" at end
	if !(((m.current + 1) == m.last) && (m.charAt(m.current+1) == 'E')) && (m.internal_Hard_NG() || m.internal_Hard_GEN_GIN_GET_GIT() || m.internal_Hard_G_Open_Syllable() || m.internal_Hard_G_Other()) {
		return true
	}

	return false
}

/**
 * Detect words where "-ge-" or "-gi-" get a 'hard' 'g'
 * even though this is usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func (m *M3) internal_Hard_G_Other() bool {
	if (m.stringAt(m.current, 4, "GETH", "GEAR", "GEIS", "GIRL", "GIVI", "GIVE", "GIFT",
		"GIRD", "GIRT", "GILV", "GILD", "GELD", "") && !m.stringAt((m.current-3), 6, "GINGIV", "")) ||
		// "gish" but not "largish"
		(m.stringAt((m.current+1), 3, "ISH", "") && (m.current > 0) && !m.stringAt(0, 4, "LARG", "")) || (m.stringAt((m.current-2), 5, "MAGED", "MEGID", "") && !((m.current + 2) == m.last)) || m.stringAt(m.current, 3, "GEZ", "") || m.stringAt(0, 4, "WEGE", "HAGE", "") || (m.stringAt((m.current-2), 6, "ONGEST", "UNGEST", "") && ((m.current + 3) == m.last) && !m.stringAt((m.current-3), 7, "CONGEST", "")) || m.stringAt(0, 5, "VOEGE", "BERGE", "HELGE", "") || (m.stringAt(0, 4, "ENGE", "BOGY", "") && (m.length == 4)) || m.stringAt(m.current, 6, "GIBBON", "") || m.stringAt(0, 10, "CORREGIDOR", "") || m.stringAt(0, 8, "INGEBORG", "") || (m.stringAt(m.current, 4, "GILL", "") && (((m.current + 3) == m.last) || ((m.current + 4) == m.last)) && !m.stringAt(0, 8, "STURGILL", "")) {
		return true
	}

	return false
}

/**
 * Detect words where "-gy-", "-gie-", "-gee-",
 * or "-gio-" get a 'hard' 'g' even though this is
 * usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func (m *M3) internal_Hard_G_Open_Syllable() bool {
	if m.stringAt((m.current+1), 3, "EYE", "") || m.stringAt((m.current-2), 4, "FOGY", "POGY", "YOGI", "") || m.stringAt((m.current-2), 5, "MAGEE", "MCGEE", "HAGIO", "") || m.stringAt((m.current-1), 4, "RGEY", "OGEY", "") || m.stringAt((m.current-3), 5, "HOAGY", "STOGY", "PORGY", "") || m.stringAt((m.current-5), 8, "CARNEGIE", "") || (m.stringAt((m.current-1), 4, "OGEY", "OGIE", "") && ((m.current + 2) == m.last)) {
		return true
	}

	return false
}

/**
 * Detect a number of contexts, mostly german names, that
 * take a 'hard' 'g'.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func (m *M3) internal_Hard_GEN_GIN_GET_GIT() bool {
	if (m.stringAt((m.current-3), 6, "FORGET", "TARGET", "MARGIT", "MARGET", "TURGEN",
		"BERGEN", "MORGEN", "JORGEN", "HAUGEN", "JERGEN",
		"JURGEN", "LINGEN", "BORGEN", "LANGEN", "KLAGEN", "STIGER", "BERGER", "") && !m.stringAt(m.current, 7, "GENETIC", "GENESIS", "") && !m.stringAt((m.current-4), 8, "PLANGENT", "")) || (m.stringAt((m.current-3), 6, "BERGIN", "FEAGIN", "DURGIN", "") && ((m.current + 2) == m.last)) || (m.stringAt((m.current-2), 5, "ENGEN", "") && !m.stringAt((m.current+3), 3, "DER", "ETI", "ESI", "")) || m.stringAt((m.current-4), 7, "JUERGEN", "") || m.stringAt(0, 5, "NAGIN", "MAGIN", "HAGIN", "") || (m.stringAt(0, 5, "ENGIN", "DEGEN", "LAGEN", "MAGEN", "NAGIN", "") && (m.length == 5)) || (m.stringAt((m.current-2), 5, "BEGET", "BEGIN", "HAGEN", "FAGIN",
		"BOGEN", "WIGIN", "NTGEN", "EIGEN",
		"WEGEN", "WAGEN", "") && !m.stringAt((m.current-5), 8, "OSPHAGEN", "")) {
		return true
	}

	return false
}

/**
 * Detect a number of contexts of '-ng-' that will
 * take a 'hard' 'g' despite being followed by a
 * front vowel.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func (m *M3) internal_Hard_NG() bool {
	if (m.stringAt((m.current-3), 4, "DANG", "FANG", "SING", "") &&
		// exception to exception
		!m.stringAt((m.current-5), 8, "DISINGEN", "")) || m.stringAt(0, 5, "INGEB", "ENGEB", "") || (m.stringAt((m.current-3), 4, "RING", "WING", "HANG", "LONG", "") && !(m.stringAt((m.current-4), 5, "CRING", "FRING", "ORANG", "TWING", "CHANG", "PHANG", "") || m.stringAt((m.current-5), 6, "SYRING", "") || m.stringAt((m.current-3), 7, "RINGENC", "RINGENT", "LONGITU", "LONGEVI", "") ||
		// e.g. 'longino', 'mastrangelo'
		(m.stringAt(m.current, 4, "GELO", "GINO", "") && ((m.current + 3) == m.last)))) || (m.stringAt((m.current-1), 3, "NGY", "") &&
		// exceptions to exception
		!(m.stringAt((m.current-3), 5, "RANGY", "MANGY", "MINGY", "") || m.stringAt((m.current-4), 6, "SPONGY", "STINGY", ""))) {
		return true
	}

	return false
}

/**
 * Encode special case where "-GA-" => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_GA_To_J() bool {
	// 'margary', 'margarine'
	if (m.stringAt((m.current-3), 7, "MARGARY", "MARGARI", "") &&
		// but not in spanish forms such as "margatita"
		!m.stringAt((m.current-3), 8, "MARGARIT", "")) ||
		m.stringAt(0, 4, "GAOL", "") || m.stringAt((m.current-2), 5, "ALGAE", "") {
		m.metaphAddExactApprox4("J", "G", "J", "K")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'H'
 *
 *
 */
func (m *M3) encode_H() {
	if m.encode_Initial_Silent_H() || m.encode_Initial_HS() || m.encode_Initial_HU_HW() || m.encode_Non_Initial_Silent_H() {
		return
	}

	//only keep if first & before vowel or btw. 2 vowels
	if !m.encode_H_Pronounced() {
		//also takes care of 'HH'
		m.current++
	}
}

/**
 * Encode cases where initial 'H' is not pronounced (in American)
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_Silent_H() bool {
	//'hour', 'herb', 'heir', 'honor'
	if m.stringAt((m.current+1), 3, "OUR", "ERB", "EIR", "") || m.stringAt((m.current+1), 4, "ONOR", "") || m.stringAt((m.current+1), 5, "ONOUR", "ONEST", "") {
		// british pronounce H in this word
		// americans give it 'H' for the name,
		// no 'H' for the plant
		if (m.current == 0) && m.stringAt(m.current, 4, "HERB", "") {
			if m.encodeVowels {
				m.metaphAdd("HA", "A")
			} else {
				m.metaphAdd("H", "A")
			}
		} else if (m.current == 0) || m.encodeVowels {
			m.metaphAdd("A", "A")
		}

		m.current++
		// don't encode vowels twice
		m.current = m.skipVowels(m.current)
		return true
	}

	return false
}

/**
 * Encode "HS-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_HS() bool {
	// old chinese pinyin transliteration
	// e.g., 'HSIAO'
	if (m.current == 0) && m.stringAt(0, 2, "HS", "") {
		m.metaphAdd("X", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode cases where "HU-" is pronounced as part of a vowel dipthong
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_HU_HW() bool {
	// spanish spellings and chinese pinyin transliteration
	if m.stringAt(0, 3, "HUA", "HUE", "HWA", "") {
		if !m.stringAt(m.current, 4, "HUEY", "") {
			m.metaphAdd("A", "A")

			if !m.encodeVowels {
				m.current += 3
			} else {
				m.current++
				// don't encode vowels twice
				for isVowel(m.charAt(m.current)) || (m.charAt(m.current) == 'W') {
					m.current++
				}
			}
			return true
		}
	}

	return false
}

/**
 * Encode cases where 'H' is silent between vowels
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Non_Initial_Silent_H() bool {
	//exceptions - 'h' not pronounced
	// "PROHIB" BUT NOT "PROHIBIT"
	if m.stringAt((m.current-2), 5, "NIHIL", "VEHEM", "LOHEN", "NEHEM",
		"MAHON", "MAHAN", "COHEN", "GAHAN", "") || m.stringAt((m.current-3), 6, "GRAHAM", "PROHIB", "FRAHER",
		"TOOHEY", "TOUHEY", "") || m.stringAt((m.current-3), 5, "TOUHY", "") || m.stringAt(0, 9, "CHIHUAHUA", "") {
		if !m.encodeVowels {
			m.current += 2
		} else {
			m.current++
			// don't encode vowels twice
			m.current = m.skipVowels(m.current)
		}
		return true
	}

	return false
}

/**
 * Encode cases where 'H' is pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_H_Pronounced() bool {
	if (((m.current == 0) || isVowel(m.charAt(m.current-1)) || ((m.current > 0) && (m.charAt(m.current-1) == 'W'))) && isVowel(m.charAt(m.current+1))) ||
		// e.g. 'alWahhab'
		((m.charAt(m.current+1) == 'H') && isVowel(m.charAt(m.current+2))) {
		m.metaphAdd("H", "H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'J'
 *
 */
func (m *M3) encode_J() {
	if m.encode_Spanish_J() || m.encode_Spanish_OJ_UJ() {
		return
	}

	m.encode_Other_J()
}

/**
 * Encode cases where initial or medial "j" is in a spanish word or name
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Spanish_J() bool {
	//obvious spanish, e.g. "jose", "san jacinto"
	if (m.stringAt((m.current+1), 3, "UAN", "ACI", "ALI", "EFE", "ICA", "IME", "OAQ", "UAR", "") && !m.stringAt(m.current, 8, "JIMERSON", "JIMERSEN", "")) || (m.stringAt((m.current+1), 3, "OSE", "") && ((m.current + 3) == m.last)) || m.stringAt((m.current+1), 4, "EREZ", "UNTA", "AIME", "AVIE", "AVIA", "") || m.stringAt((m.current+1), 6, "IMINEZ", "ARAMIL", "") || (((m.current + 2) == m.last) && m.stringAt((m.current-2), 5, "MEJIA", "")) || m.stringAt((m.current-2), 5, "TEJED", "TEJAD", "LUJAN", "FAJAR", "BEJAR", "BOJOR", "CAJIG",
		"DEJAS", "DUJAR", "DUJAN", "MIJAR", "MEJOR", "NAJAR",
		"NOJOS", "RAJED", "RIJAL", "REJON", "TEJAN", "UIJAN", "") || m.stringAt((m.current-3), 8, "ALEJANDR", "GUAJARDO", "TRUJILLO", "") || (m.stringAt((m.current-2), 5, "RAJAS", "") && (m.current > 2)) || (m.stringAt((m.current-2), 5, "MEJIA", "") && !m.stringAt((m.current-2), 6, "MEJIAN", "")) || m.stringAt((m.current-1), 5, "OJEDA", "") || m.stringAt((m.current-3), 5, "LEIJA", "MINJA", "") || m.stringAt((m.current-3), 6, "VIAJES", "GRAJAL", "") || m.stringAt(m.current, 8, "JAUREGUI", "") || m.stringAt((m.current-4), 8, "HINOJOSA", "") || m.stringAt(0, 4, "SAN ", "") || (((m.current + 1) == m.last) && (m.charAt(m.current+1) == 'O') &&
		// exceptions
		!(m.stringAt(0, 4, "TOJO", "") || m.stringAt(0, 5, "BANJO", "") || m.stringAt(0, 6, "MARYJO", ""))) {
		// americans pronounce "juan" as 'wan'
		// and "marijuana" and "tijuana" also
		// do not get the 'H' as in spanish, so
		// just treat it like a vowel in these cases
		if !(m.stringAt(m.current, 4, "JUAN", "") || m.stringAt(m.current, 4, "JOAQ", "")) {
			m.metaphAdd("H", "H")
		} else {
			if m.current == 0 {
				m.metaphAdd("A", "A")
			}
		}
		m.advanceCounter(2, 1)
		return true
	}

	// Jorge gets 2nd HARHA. also JULIO, JESUS
	if m.stringAt((m.current+1), 4, "ORGE", "ULIO", "ESUS", "") && !m.stringAt(0, 6, "JORGEN", "") {
		// get both consonants for "jorge"
		if ((m.current + 4) == m.last) && m.stringAt((m.current+1), 4, "ORGE", "") {
			if m.encodeVowels {
				m.metaphAdd("JARJ", "HARHA")
			} else {
				m.metaphAdd("JRJ", "HRH")
			}
			m.advanceCounter(5, 5)
			return true
		}

		m.metaphAdd("J", "H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode cases where 'J' is clearly in a german word or name
 * that americans pronounce in the german fashion
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_German_J() bool {
	if m.stringAt((m.current+1), 2, "AH", "") || (m.stringAt((m.current+1), 5, "OHANN", "") && ((m.current + 5) == m.last)) || (m.stringAt((m.current+1), 3, "UNG", "") && !m.stringAt((m.current+1), 4, "UNGL", "")) || m.stringAt((m.current+1), 3, "UGO", "") {
		m.metaphAdd("A", "A")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-JOJ-" and "-JUJ-" as spanish words
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Spanish_OJ_UJ() bool {
	if m.stringAt((m.current + 1), 5, "OJOBA", "UJUY ", "") {
		if m.encodeVowels {
			m.metaphAdd("HAH", "HAH")
		} else {
			m.metaphAdd("HH", "HH")
		}

		m.advanceCounter(4, 3)
		return true
	}

	return false
}

/**
 * Encode 'J' => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_J_To_J() bool {
	if isVowel(m.charAt(m.current + 1)) {
		if (m.current == 0) && m.names_Beginning_With_J_That_Get_Alt_Y() {
			// 'Y' is a vowel so encode
			// is as 'A'
			if m.encodeVowels {
				m.metaphAdd("JA", "A")
			} else {
				m.metaphAdd("J", "A")
			}
		} else {
			if m.encodeVowels {
				m.metaphAdd("JA", "JA")
			} else {
				m.metaphAdd("J", "J")
			}
		}

		m.current++
		m.current = m.skipVowels(m.current)
		return false
	} else {
		m.metaphAdd("J", "J")
		m.current++
		return true
	}

	//		return false;
}

/**
 * Encode 'J' toward end in spanish words
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Spanish_J_2() bool {
	// spanish forms e.g. "brujo", "badajoz"
	if (((m.current - 2) == 0) && m.stringAt((m.current-2), 4, "BOJA", "BAJA", "BEJA", "BOJO", "MOJA", "MOJI", "MEJI", "")) || (((m.current - 3) == 0) && m.stringAt((m.current-3), 5, "FRIJO", "BRUJO", "BRUJA", "GRAJE", "GRIJA", "LEIJA", "QUIJA", "")) || (((m.current + 3) == m.last) && m.stringAt((m.current-1), 5, "AJARA", "")) || (((m.current + 2) == m.last) && m.stringAt((m.current-1), 4, "AJOS", "EJOS", "OJAS", "OJOS", "UJON", "AJOZ", "AJAL", "UJAR", "EJON", "EJAN", "")) || (((m.current + 1) == m.last) && (m.stringAt((m.current-1), 3, "OJA", "EJA", "") && !m.stringAt(0, 4, "DEJA", ""))) {
		m.metaphAdd("H", "H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'J' as vowel in some exception cases
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_J_As_Vowel() bool {
	if m.stringAt(m.current, 5, "JEWSK", "") {
		m.metaphAdd("J", "")
		return true
	}

	// e.g. "stijl", "sejm" - dutch, scandanavian, and eastern european spellings
	if (m.stringAt((m.current+1), 1, "L", "T", "K", "S", "N", "M", "") &&
		// except words from hindi and arabic
		!m.stringAt((m.current+2), 1, "A", "")) || m.stringAt(0, 9, "HALLELUJA", "LJUBLJANA", "") || m.stringAt(0, 4, "LJUB", "BJOR", "") || m.stringAt(0, 5, "HAJEK", "") || m.stringAt(0, 3, "WOJ", "") ||
		// e.g. 'fjord'
		m.stringAt(0, 2, "FJ", "") ||
		// e.g. 'rekjavik', 'blagojevic'
		m.stringAt(m.current, 5, "JAVIK", "JEVIC", "") || (((m.current + 1) == m.last) && m.stringAt(0, 5, "SONJA", "TANJA", "TONJA", "")) {
		return true
	}
	return false
}

/**
 * Call routines to encode 'J', in proper order
 *
 */
func (m *M3) encode_Other_J() {
	if m.current == 0 {
		if m.encode_German_J() {
			return
		} else {
			if m.encode_J_To_J() {
				return
			}
		}
	} else {
		if m.encode_Spanish_J_2() {
			return
		} else if !m.encode_J_As_Vowel() {
			m.metaphAdd("J", "J")
		}

		//it could happen! e.g. "hajj"
		// eat redundant 'J'
		if m.charAt(m.current+1) == 'J' {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Encode 'K'
 *
 *
 */
func (m *M3) encode_K() {
	if !m.encode_Silent_K() {
		m.metaphAdd("K", "K")

		// eat redundant 'K's and 'Q's
		if (m.charAt(m.current+1) == 'K') || (m.charAt(m.current+1) == 'Q') {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Encode cases where 'K' is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_K() bool {
	//skip this except for special cases
	if (m.current == 0) && m.stringAt(m.current, 2, "KN", "") {
		if !(m.stringAt((m.current+2), 5, "ESSET", "IEVEL", "") || m.stringAt((m.current+2), 3, "ISH", "")) {
			m.current += 1
			return true
		}
	}

	// e.g. "know", "knit", "knob"
	if (m.stringAt((m.current+1), 3, "NOW", "NIT", "NOT", "NOB", "") &&
		// exception, "slipknot" => SLPNT but "banknote" => PNKNT
		!m.stringAt(0, 8, "BANKNOTE", "")) || m.stringAt((m.current+1), 4, "NOCK", "NUCK", "NIFE", "NACK", "") || m.stringAt((m.current+1), 5, "NIGHT", "") {
		// N already encoded before
		// e.g. "penknife"
		if (m.current > 0) && m.charAt(m.current-1) == 'N' {
			m.current += 2
		} else {
			m.current++
		}

		return true
	}

	return false
}

/**
 * Encode 'L'
 *
 * Includes special vowel transposition
 * encoding, where 'LE' => AL
 *
 */
func (m *M3) encode_L() {
	// logic below needs to know this
	// after 'm.current' variable changed
	save_current := m.current

	m.interpolate_Vowel_When_Cons_L_At_End()

	if m.encode_LELY_To_L() || m.encode_COLONEL() || m.encode_French_AULT() || m.encode_French_EUIL() || m.encode_French_OULX() || m.encode_Silent_L_In_LM() || m.encode_Silent_L_In_LK_LV() || m.encode_Silent_L_In_OULD() {
		return
	}

	if m.encode_LL_As_Vowel_Cases() {
		return
	}

	m.encode_LE_Cases(save_current)
}

/**
 * Cases where an L follows D, G, or T at the
 * end have a schwa pronounced before the L
 *
 */
func (m *M3) interpolate_Vowel_When_Cons_L_At_End() {
	if m.encodeVowels == true {
		// e.g. "ertl", "vogl"
		if (m.current == m.last) && m.stringAt((m.current-1), 1, "D", "G", "T", "") {
			m.metaphAdd("A", "A")
		}
	}
}

/**
 * Catch cases where 'L' spelled twice but pronounced
 * once, e.g., 'DOCILELY' => TSL
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_LELY_To_L() bool {
	// e.g. "agilely", "docilely"
	if m.stringAt((m.current-1), 5, "ILELY", "") && ((m.current + 3) == m.last) {
		m.metaphAdd("L", "L")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode special case "colonel" => KRNL. Can somebody tell
 * me how this pronounciation came to be?
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_COLONEL() bool {
	if m.stringAt((m.current - 2), 7, "COLONEL", "") {
		m.metaphAdd("R", "R")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-AULT-", found in a french names
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_French_AULT() bool {
	// e.g. "renault" and "foucault", well known to americans, but not "fault"
	if (m.current > 3) && (m.stringAt((m.current-3), 5, "RAULT", "NAULT", "BAULT", "SAULT", "GAULT", "CAULT", "") || m.stringAt((m.current-4), 6, "REAULT", "RIAULT", "NEAULT", "BEAULT", "")) && !(rootOrInflections(m.inWord, "ASSAULT") || m.stringAt((m.current-8), 10, "SOMERSAULT", "") || m.stringAt((m.current-9), 11, "SUMMERSAULT", "")) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-EUIL-", always found in a french word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_French_EUIL() bool {
	// e.g. "auteuil"
	if m.stringAt((m.current-3), 4, "EUIL", "") && (m.current == m.last) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-OULX", always found in a french word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_French_OULX() bool {
	// e.g. "proulx"
	if m.stringAt((m.current-2), 4, "OULX", "") && ((m.current + 1) == m.last) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes contexts where 'L' is not pronounced in "-LM-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_L_In_LM() bool {
	if m.stringAt(m.current, 2, "LM", "LN", "") {
		// e.g. "lincoln", "holmes", "psalm", "salmon"
		if (m.stringAt((m.current-2), 4, "COLN", "CALM", "BALM", "MALM", "PALM", "") || (m.stringAt((m.current-1), 3, "OLM", "") && ((m.current + 1) == m.last)) || m.stringAt((m.current-3), 5, "PSALM", "QUALM", "") || m.stringAt((m.current-2), 6, "SALMON", "HOLMES", "") || m.stringAt((m.current-1), 6, "ALMOND", "") || ((m.current == 1) && m.stringAt((m.current-1), 4, "ALMS", ""))) && (!m.stringAt((m.current+2), 1, "A", "") && !m.stringAt((m.current-2), 5, "BALMO", "") && !m.stringAt((m.current-2), 6, "PALMER", "PALMOR", "BALMER", "") && !m.stringAt((m.current-3), 5, "THALM", "")) {
			m.current++
			return true
		} else {
			m.metaphAdd("L", "L")
			m.current++
			return true
		}
	}

	return false
}

/**
 * Encodes contexts where '-L-' is silent in 'LK', 'LV'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_L_In_LK_LV() bool {
	if (m.stringAt((m.current-2), 4, "WALK", "YOLK", "FOLK", "HALF", "TALK", "CALF", "BALK", "CALK", "") || (m.stringAt((m.current-2), 4, "POLK", "") && !m.stringAt((m.current-2), 5, "POLKA", "WALKO", "")) || (m.stringAt((m.current-2), 4, "HALV", "") && !m.stringAt((m.current-2), 5, "HALVA", "HALVO", "")) || (m.stringAt((m.current-3), 5, "CAULK", "CHALK", "BAULK", "FAULK", "") && !m.stringAt((m.current-4), 6, "SCHALK", "")) || (m.stringAt((m.current-2), 5, "SALVE", "CALVE", "") || m.stringAt((m.current-2), 6, "SOLDER", "")) &&
		// exceptions to above cases where 'L' is usually pronounced
		!m.stringAt((m.current-2), 6, "SALVER", "CALVER", "")) && !m.stringAt((m.current-5), 9, "GONSALVES", "GONCALVES", "") && !m.stringAt((m.current-2), 6, "BALKAN", "TALKAL", "") && !m.stringAt((m.current-3), 5, "PAULK", "CHALF", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode 'L' in contexts of "-OULD-" where it is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_L_In_OULD() bool {
	//'would', 'could'
	if m.stringAt((m.current-3), 5, "WOULD", "COULD", "") || (m.stringAt((m.current-4), 6, "SHOULD", "") && !m.stringAt((m.current-4), 8, "SHOULDER", "")) {
		m.metaphAddExactApprox("D", "T")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-ILLA-" and "-ILLE-" in spanish and french
 * contexts were americans know to pronounce it as a 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_LL_As_Vowel_Special_Cases() bool {
	if m.stringAt((m.current-5), 8, "TORTILLA", "") || m.stringAt((m.current-8), 11, "RATATOUILLE", "") ||
		// e.g. 'guillermo', "veillard"
		(m.stringAt(0, 5, "GUILL", "VEILL", "GAILL", "") &&
			// 'guillotine' usually has '-ll-' pronounced as 'L' in english
			!(m.stringAt((m.current-3), 7, "GUILLOT", "GUILLOR", "GUILLEN", "") || (m.stringAt(0, 5, "GUILL", "") && (m.length == 5)))) ||
		// e.g. "brouillard", "gremillion"
		m.stringAt(0, 7, "BROUILL", "GREMILL", "ROBILL", "") ||
		// e.g. 'mireille'
		(m.stringAt((m.current-2), 5, "EILLE", "") && ((m.current + 2) == m.last) &&
			// exception "reveille" usually pronounced as 're-vil-lee'
			!m.stringAt((m.current-5), 8, "REVEILLE", "")) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode other spanish cases where "-LL-" is pronounced as 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_LL_As_Vowel() bool {
	//spanish e.g. "cabrillo", "gallegos" but also "gorilla", "ballerina" -
	// give both pronounciations since an american might pronounce "cabrillo"
	// in the spanish or the american fashion.
	if (((m.current + 3) == m.length) && m.stringAt((m.current-1), 4, "ILLO", "ILLA", "ALLE", "")) || (((m.stringAt((m.last-1), 2, "AS", "OS", "") || m.stringAt(m.last, 2, "AS", "OS", "") || m.stringAt(m.last, 1, "A", "O", "")) && m.stringAt((m.current-1), 2, "AL", "IL", "")) && !m.stringAt((m.current-1), 4, "ALLA", "")) || m.stringAt(0, 5, "VILLE", "VILLA", "") || m.stringAt(0, 8, "GALLARDO", "VALLADAR", "MAGALLAN", "CAVALLAR", "BALLASTE", "") || m.stringAt(0, 3, "LLA", "") {
		m.metaphAdd("L", "")
		m.current += 2
		return true
	}
	return false
}

/**
 * Call routines to encode "-LL-", in proper order
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_LL_As_Vowel_Cases() bool {
	if m.charAt(m.current+1) == 'L' {
		if m.encode_LL_As_Vowel_Special_Cases() {
			return true
		} else if m.encode_LL_As_Vowel() {
			return true
		}
		m.current += 2

	} else {
		m.current++
	}

	return false
}

/**
 * Encode vowel-encoding cases where "-LE-" is pronounced "-EL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Vowel_LE_Transposition(save_current int) bool {
	// transposition of vowel sound and L occurs in many words,
	// e.g. "bristle", "dazzle", "goggle" => KAKAL
	if m.encodeVowels && (save_current > 1) && !isVowel(m.charAt(save_current-1)) && (m.charAt(save_current+1) == 'E') && (m.charAt(save_current-1) != 'L') && (m.charAt(save_current-1) != 'R') &&
		// lots of exceptions to this:
		!isVowel(m.charAt(save_current+2)) && !m.stringAt(0, 7, "ECCLESI", "COMPLEC", "COMPLEJ", "ROBLEDO", "") && !m.stringAt(0, 5, "MCCLE", "MCLEL", "") && !m.stringAt(0, 6, "EMBLEM", "KADLEC", "") && !(((save_current + 2) == m.last) && m.stringAt(save_current, 3, "LET", "")) && !m.stringAt(save_current, 7, "LETTING", "") && !m.stringAt(save_current, 6, "LETELY", "LETTER", "LETION", "LETIAN", "LETING", "LETORY", "") && !m.stringAt(save_current, 5, "LETUS", "LETIV", "") && !m.stringAt(save_current, 4, "LESS", "LESQ", "LECT", "LEDG", "LETE", "LETH", "LETS", "LETT", "") && !m.stringAt(save_current, 3, "LEG", "LER", "LEX", "") &&
		// e.g. "complement" !=> KAMPALMENT
		!(m.stringAt(save_current, 6, "LEMENT", "") && !(m.stringAt((m.current-5), 6, "BATTLE", "TANGLE", "PUZZLE", "RABBLE", "BABBLE", "") ||
			m.stringAt((m.current-4), 5, "TABLE", ""))) && !(((save_current + 2) == m.last) && m.stringAt((save_current-2), 5, "OCLES", "ACLES", "AKLES", "")) && !m.stringAt((save_current-3), 5, "LISLE", "AISLE", "") && !m.stringAt(0, 4, "ISLE", "") && !m.stringAt(0, 6, "ROBLES", "") && !m.stringAt((save_current-4), 7, "PROBLEM", "RESPLEN", "") && !m.stringAt((save_current-3), 6, "REPLEN", "") && !m.stringAt((save_current-2), 4, "SPLE", "") && (m.charAt(save_current-1) != 'H') && (m.charAt(save_current-1) != 'W') {
		m.metaphAdd("AL", "AL")
		m.flag_AL_inversion = true

		// eat redundant 'L'
		if m.charAt(save_current+2) == 'L' {
			m.current = save_current + 3
		}
		return true
	}

	return false
}

/**
 * Encode special vowel-encoding cases where 'E' is not
 * silent at the end of a word as is the usual case
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Vowel_Preserve_Vowel_After_L(save_current int) bool {
	// an example of where the vowel would NOT need to be preserved
	// would be, say, "hustled", where there is no vowel pronounced
	// between the 'l' and the 'd'
	if m.encodeVowels && !isVowel(m.charAt(save_current-1)) && (m.charAt(save_current+1) == 'E') && (save_current > 1) && ((save_current + 1) != m.last) && !(m.stringAt((save_current+1), 2, "ES", "ED", "") && ((save_current + 2) == m.last)) && !m.stringAt((save_current-1), 5, "RLEST", "") {
		m.metaphAdd("LA", "LA")
		m.current = m.skipVowels(m.current)
		return true
	}

	return false
}

/**
 * Call routines to encode "-LE-", in proper order
 *
 * @param save_current index of actual current letter
 *
 */
func (m *M3) encode_LE_Cases(save_current int) {
	if m.encode_Vowel_LE_Transposition(save_current) {
		return
	} else {
		if m.encode_Vowel_Preserve_Vowel_After_L(save_current) {
			return
		} else {
			m.metaphAdd("L", "L")
		}
	}
}

/**
 * Encode "-M-"
 *
 */
func (m *M3) encode_M() {
	if m.encode_Silent_M_At_Beginning() || m.encode_MR_And_MRS() || m.encode_MAC() || m.encode_MPT() {
		return
	}

	// Silent 'B' should really be handled
	// under 'B", not here under 'M'!
	m.encode_MB()

	m.metaphAdd("M", "M")
}

/**
 * Encode cases where 'M' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_M_At_Beginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(m.current, 2, "MN", "") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode special cases "Mr." and "Mrs."
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_MR_And_MRS() bool {
	if (m.current == 0) && m.stringAt(m.current, 2, "MR", "") {
		// exceptions for "mr." and "mrs."
		if (m.length == 2) && m.stringAt(m.current, 2, "MR", "") {
			if m.encodeVowels {
				m.metaphAdd("MASTAR", "MASTAR")
			} else {
				m.metaphAdd("MSTR", "MSTR")
			}
			m.current += 2
			return true
		} else if (m.length == 3) && m.stringAt(m.current, 3, "MRS", "") {
			if m.encodeVowels {
				m.metaphAdd("MASAS", "MASAS")
			} else {
				m.metaphAdd("MSS", "MSS")
			}
			m.current += 3
			return true
		}
	}

	return false
}

/**
 * Encode "Mac-" and "Mc-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_MAC() bool {
	// should only find irish and
	// scottish names e.g. 'macintosh'
	if (m.current == 0) && (m.stringAt(0, 7, "MACIVER", "MACEWEN", "") || m.stringAt(0, 8, "MACELROY", "MACILROY", "") || m.stringAt(0, 9, "MACINTOSH", "") || m.stringAt(0, 2, "MC", "")) {
		if m.encodeVowels {
			m.metaphAdd("MAK", "MAK")
		} else {
			m.metaphAdd("MK", "MK")
		}

		if m.stringAt(0, 2, "MC", "") {
			if m.stringAt((m.current+2), 1, "K", "G", "Q", "") &&
				// watch out for e.g. "McGeorge"
				!m.stringAt((m.current+2), 4, "GEOR", "") {
				m.current += 3
			} else {
				m.current += 2
			}
		} else {
			m.current += 3
		}

		return true
	}

	return false
}

/**
 * Encode silent 'M' in context of "-MPT-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_MPT() bool {
	if m.stringAt((m.current-2), 8, "COMPTROL", "") || m.stringAt((m.current-4), 7, "ACCOMPT", "") {
		m.metaphAdd("N", "N")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test if 'B' is silent in these contexts
 *
 * @return true if 'B' is silent in this context
 *
 */
func (m *M3) test_Silent_MB_1() bool {
	// e.g. "LAMB", "COMB", "LIMB", "DUMB", "BOMB"
	// Handle combining roots first
	if ((m.current == 3) && m.stringAt((m.current-3), 5, "THUMB", "")) || ((m.current == 2) && m.stringAt((m.current-2), 4, "DUMB", "BOMB", "DAMN", "LAMB", "NUMB", "TOMB", "")) {
		return true
	}

	return false
}

/**
 * Test if 'B' is pronounced in this context
 *
 * @return true if 'B' is pronounced in this context
 *
 */
func (m *M3) test_Pronounced_MB() bool {
	if m.stringAt((m.current-2), 6, "NUMBER", "") || (m.stringAt((m.current+2), 1, "A", "") && !m.stringAt((m.current-2), 7, "DUMBASS", "")) || m.stringAt((m.current+2), 1, "O", "") || m.stringAt((m.current-2), 6, "LAMBEN", "LAMBER", "LAMBET", "TOMBIG", "LAMBRE", "") {
		return true
	}

	return false
}

/**
 * Test whether "-B-" is silent in these contexts
 *
 * @return true if 'B' is silent in this context
 *
 */
func (m *M3) test_Silent_MB_2() bool {
	// 'M' is the current letter
	if (m.charAt(m.current+1) == 'B') && (m.current > 1) && (((m.current + 1) == m.last) ||
		// other situations where "-MB-" is at end of root
		// but not at end of word. The tests are for standard
		// noun suffixes.
		// e.g. "climbing" => KLMNK
		m.stringAt((m.current+2), 3, "ING", "ABL", "") || m.stringAt((m.current+2), 4, "LIKE", "") || ((m.charAt(m.current+2) == 'S') && ((m.current + 2) == m.last)) || m.stringAt((m.current-5), 7, "BUNCOMB", "") ||
		// e.g. "bomber",
		(m.stringAt((m.current+2), 2, "ED", "ER", "") && ((m.current + 3) == m.last) && (m.stringAt(0, 5, "CLIMB", "PLUMB", "") ||
			// e.g. "beachcomber"
			!m.stringAt((m.current-1), 5, "IMBER", "AMBER", "EMBER", "UMBER", "")) &&
			// exceptions
			!m.stringAt((m.current-2), 6, "CUMBER", "SOMBER", ""))) {
		return true
	}

	return false
}

/**
 * Test if 'B' is pronounced in these "-MB-" contexts
 *
 * @return true if "-B-" is pronounced in these contexts
 *
 */
func (m *M3) test_Pronounced_MB_2() bool {
	// e.g. "bombastic", "umbrage", "flamboyant"
	if m.stringAt((m.current-1), 5, "OMBAS", "OMBAD", "UMBRA", "") || m.stringAt((m.current-3), 4, "FLAM", "") {
		return true
	}

	return false
}

/**
 * Tests for contexts where "-N-" is silent when after "-M-"
 *
 * @return true if "-N-" is silent in these contexts
 *
 */
func (m *M3) test_MN() bool {

	if (m.charAt(m.current+1) == 'N') && (((m.current + 1) == m.last) ||
		// or at the end of a word but followed by suffixes
		(m.stringAt((m.current+2), 3, "ING", "EST", "") && ((m.current + 4) == m.last)) || ((m.charAt(m.current+2) == 'S') && ((m.current + 2) == m.last)) || (m.stringAt((m.current+2), 2, "LY", "ER", "ED", "") && ((m.current + 3) == m.last)) || m.stringAt((m.current-2), 9, "DAMNEDEST", "") || m.stringAt((m.current-5), 9, "GODDAMNIT", "")) {
		return true
	}

	return false
}

/**
 * Call routines to encode "-MB-", in proper order
 *
 */
func (m *M3) encode_MB() {
	if m.test_Silent_MB_1() {
		if m.test_Pronounced_MB() {
			m.current++
		} else {
			m.current += 2
		}
	} else if m.test_Silent_MB_2() {
		if m.test_Pronounced_MB_2() {
			m.current++
		} else {
			m.current += 2
		}
	} else if m.test_MN() {
		m.current += 2
	} else {
		// eat redundant 'M'
		if m.charAt(m.current+1) == 'M' {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Encode "-N-"
 *
 */
func (m *M3) encode_N() {
	if m.encode_NCE() {
		return
	}

	// eat redundant 'N'
	if m.charAt(m.current+1) == 'N' {
		m.current += 2
	} else {
		m.current++
	}

	if !m.stringAt((m.current-3), 8, "MONSIEUR", "") &&
		// e.g. "aloneness",
		!m.stringAt((m.current-3), 6, "NENESS", "") {
		m.metaphAdd("N", "N")
	}
}

/**
 * Encode "-NCE-" and "-NSE-"
 * "entrance" is pronounced exactly the same as "entrants"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_NCE() bool {
	//'acceptance', 'accountancy'
	if m.stringAt((m.current+1), 1, "C", "S", "") && m.stringAt((m.current+2), 1, "E", "Y", "I", "") && (((m.current + 2) == m.last) || ((m.current+3) == m.last) && (m.charAt(m.current+3) == 'S')) {
		m.metaphAdd("NTS", "NTS")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-P-"
 *
 */
func (m *M3) encode_P() {
	if m.encode_Silent_P_At_Beginning() || m.encode_PT() || m.encode_PH() || m.encode_PPH() || m.encode_RPS() || m.encode_COUP() || m.encode_PNEUM() || m.encode_PSYCH() || m.encode_PSALM() {
		return
	}

	m.encode_PB()

	m.metaphAdd("P", "P")
}

/**
 * Encode cases where "-P-" is silent at the start of a word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_P_At_Beginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(m.current, 2, "PN", "PF", "PS", "PT", "") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode cases where "-P-" is silent before "-T-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PT() bool {
	// 'pterodactyl', 'receipt', 'asymptote'
	if m.charAt(m.current+1) == 'T' {
		if ((m.current == 0) && m.stringAt(m.current, 5, "PTERO", "")) || m.stringAt((m.current-5), 7, "RECEIPT", "") || m.stringAt((m.current-4), 8, "ASYMPTOT", "") {
			m.metaphAdd("T", "T")
			m.current += 2
			return true
		}
	}
	return false
}

/**
 * Encode "-PH-", usually as F, with exceptions for
 * cases where it is silent, or where the 'P' and 'T'
 * are pronounced seperately because they belong to
 * two different words in a combining form
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PH() bool {
	if m.charAt(m.current+1) == 'H' {
		// 'PH' silent in these contexts
		if m.stringAt(m.current, 9, "PHTHALEIN", "") || ((m.current == 0) && m.stringAt(m.current, 4, "PHTH", "")) || m.stringAt((m.current-3), 10, "APOPHTHEGM", "") {
			m.metaphAdd("0", "0")
			m.current += 4
		} else
		// combining forms
		//'sheepherd', 'upheaval', 'cupholder'
		if (m.current > 0) && (m.stringAt((m.current+2), 3, "EAD", "OLE", "ELD", "ILL", "OLD", "EAP", "ERD",
			"ARD", "ANG", "ORN", "EAV", "ART", "") || m.stringAt((m.current+2), 4, "OUSE", "") || (m.stringAt((m.current+2), 2, "AM", "") && !m.stringAt((m.current-1), 5, "LPHAM", "")) || m.stringAt((m.current+2), 5, "AMMER", "AZARD", "UGGER", "") || m.stringAt((m.current+2), 6, "OLSTER", "")) && !m.stringAt((m.current-3), 5, "LYMPH", "NYMPH", "") {
			m.metaphAdd("P", "P")
			m.advanceCounter(3, 2)
		} else {
			m.metaphAdd("F", "F")
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode "-PPH-". I don't know why the greek poet's
 * name is transliterated this way...
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PPH() bool {
	// 'sappho'
	if (m.charAt(m.current+1) == 'P') && ((m.current + 2) < m.length) && (m.charAt(m.current+2) == 'H') {
		m.metaphAdd("F", "F")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-CORPS-" where "-PS-" not pronounced
 * since the cognate is here from the french
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_RPS() bool {
	//'-corps-', 'corpsman'
	if m.stringAt((m.current-3), 5, "CORPS", "") && !m.stringAt((m.current-3), 6, "CORPSE", "") {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-COUP-" where "-P-" is not pronounced
 * since the word is from the french
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_COUP() bool {
	//'coup'
	if (m.current == m.last) && m.stringAt((m.current-3), 4, "COUP", "") && !m.stringAt((m.current-5), 6, "RECOUP", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode 'P' in non-initial contexts of "-PNEUM-"
 * where is also silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PNEUM() bool {
	//'-pneum-'
	if m.stringAt((m.current + 1), 4, "NEUM", "") {
		m.metaphAdd("N", "N")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special case "-PSYCH-" where two encodings need to be
 * accounted for in one syllable, one for the 'PS' and one for
 * the 'CH'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PSYCH() bool {
	//'-psych-'
	if m.stringAt((m.current + 1), 4, "SYCH", "") {
		if m.encodeVowels {
			m.metaphAdd("SAK", "SAK")
		} else {
			m.metaphAdd("SK", "SK")
		}

		m.current += 5
		return true
	}

	return false
}

/**
 * Encode 'P' in context of "-PSALM-", where it has
 * become silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_PSALM() bool {
	//'-psalm-'
	if m.stringAt((m.current + 1), 4, "SALM", "") {
		// go ahead and encode entire word
		if m.encodeVowels {
			m.metaphAdd("SAM", "SAM")
		} else {
			m.metaphAdd("SM", "SM")
		}

		m.current += 5
		return true
	}

	return false
}

/**
 * Eat redundant 'B' or 'P'
 *
 */
func (m *M3) encode_PB() {
	// e.g. "campbell", "raspberry"
	// eat redundant 'P' or 'B'
	if m.stringAt((m.current + 1), 1, "P", "B", "") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode "-Q-"
 *
 */
func (m *M3) encode_Q() {
	// current pinyin
	if m.stringAt(m.current, 3, "QIN", "") {
		m.metaphAdd("X", "X")
		m.current++
		return
	}

	// eat redundant 'Q'
	if m.charAt(m.current+1) == 'Q' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("K", "K")
}

/**
 * Encode "-R-"
 *
 */
func (m *M3) encode_R() {
	if m.encode_RZ() {
		return
	}

	if !m.test_Silent_R() {
		if !m.encode_Vowel_RE_Transposition() {
			m.metaphAdd("R", "R")
		}
	}

	// eat redundant 'R'; also skip 'S' as well as 'R' in "poitiers"
	if (m.charAt(m.current+1) == 'R') || m.stringAt((m.current-6), 8, "POITIERS", "") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode "-RZ-" according
 * to american and polish pronunciations
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_RZ() bool {
	if m.stringAt((m.current-2), 4, "GARZ", "KURZ", "MARZ", "MERZ", "HERZ", "PERZ", "WARZ", "") || m.stringAt(m.current, 5, "RZANO", "RZOLA", "") || m.stringAt((m.current-1), 4, "ARZA", "ARZN", "") {
		return false
	}

	// 'yastrzemski' usually has 'z' silent in
	// united states, but should get 'X' in poland
	if m.stringAt((m.current - 4), 11, "YASTRZEMSKI", "") {
		m.metaphAdd("R", "X")
		m.current += 2
		return true
	}
	// 'BRZEZINSKI' gets two pronunciations
	// in the united states, neither of which
	// are authentically polish
	if m.stringAt((m.current - 1), 10, "BRZEZINSKI", "") {
		m.metaphAdd("RS", "RJ")
		// skip over 2nd 'Z'
		m.current += 4
		return true
	} else
	// 'z' in 'rz after voiceless consonant gets 'X'
	// in alternate polish style pronunciation
	if m.stringAt((m.current-1), 3, "TRZ", "PRZ", "KRZ", "") || (m.stringAt(m.current, 2, "RZ", "") && (isVowel(m.charAt(m.current-1)) || (m.current == 0))) {
		m.metaphAdd("RS", "X")
		m.current += 2
		return true
	} else
	// 'z' in 'rz after voiceled consonant, vowel, or at
	// beginning gets 'J' in alternate polish style pronunciation
	if m.stringAt((m.current - 1), 3, "BRZ", "DRZ", "GRZ", "") {
		m.metaphAdd("RS", "J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test whether 'R' is silent in this context
 *
 * @return true if 'R' is silent in this context
 *
 */
func (m *M3) test_Silent_R() bool {
	// test cases where 'R' is silent, either because the
	// word is from the french or because it is no longer pronounced.
	// e.g. "rogier", "monsieur", "surburban"
	if ((m.current == m.last) &&
		// reliably french word ending
		m.stringAt((m.current-2), 3, "IER", "") &&
		// e.g. "metier"
		(m.stringAt((m.current-5), 3, "MET", "VIV", "LUC", "") ||
			// e.g. "cartier", "bustier"
			m.stringAt((m.current-6), 4, "CART", "DOSS", "FOUR", "OLIV", "BUST", "DAUM", "ATEL",
				"SONN", "CORM", "MERC", "PELT", "POIR", "BERN", "FORT", "GREN",
				"SAUC", "GAGN", "GAUT", "GRAN", "FORC", "MESS", "LUSS", "MEUN",
				"POTH", "HOLL", "CHEN", "") ||
			// e.g. "croupier"
			m.stringAt((m.current-7), 5, "CROUP", "TORCH", "CLOUT", "FOURN", "GAUTH", "TROTT",
				"DEROS", "CHART", "") ||
			// e.g. "chevalier"
			m.stringAt((m.current-8), 6, "CHEVAL", "LAVOIS", "PELLET", "SOMMEL", "TREPAN", "LETELL", "COLOMB", "") || m.stringAt((m.current-9), 7, "CHARCUT", "") || m.stringAt((m.current-10), 8, "CHARPENT", ""))) || m.stringAt((m.current-2), 7, "SURBURB", "WORSTED", "") || m.stringAt((m.current-2), 9, "WORCESTER", "") || m.stringAt((m.current-7), 8, "MONSIEUR", "") || m.stringAt((m.current-6), 8, "POITIERS", "") {
		return true
	}

	return false
}

/**
 * Encode '-re-" as 'AR' in contexts
 * where this is the correct pronunciation
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Vowel_RE_Transposition() bool {
	// -re inversion is just like
	// -le inversion
	// e.g. "fibre" => FABAR or "centre" => SANTAR
	if (m.encodeVowels) && (m.charAt(m.current+1) == 'E') && (m.length > 3) && !m.stringAt(0, 5, "OUTRE", "LIBRE", "ANDRE", "") && !(m.stringAt(0, 4, "FRED", "TRES", "") && (m.length == 4)) && !m.stringAt((m.current-2), 5, "LDRED", "LFRED", "NDRED", "NFRED", "NDRES", "TRES", "IFRED", "") && !isVowel(m.charAt(m.current-1)) && (((m.current + 1) == m.last) || (((m.current + 2) == m.last) && m.stringAt((m.current+2), 1, "D", "S", ""))) {
		m.metaphAdd("AR", "AR")
		return true
	}

	return false
}

/**
 * Encode "-S-"
 *
 */
func (m *M3) encode_S() {
	if m.encode_SKJ() || m.encode_Special_SW() || m.encode_SJ() || m.encode_Silent_French_S_Final() || m.encode_Silent_French_S_Internal() || m.encode_ISL() || m.encode_STL() || m.encode_Christmas() || m.encode_STHM() || m.encode_ISTEN() || m.encode_Sugar() || m.encode_SH() || m.encode_SCH() || m.encode_SUR() || m.encode_SU() || m.encode_SSIO() || m.encode_SS() || m.encode_SIA() || m.encode_SIO() || m.encode_Anglicisations() || m.encode_SC() || m.encode_SEA_SUI_SIER() || m.encode_SEA() {
		return
	}

	m.metaphAdd("S", "S")

	if m.stringAt((m.current+1), 1, "S", "Z", "") && !m.stringAt((m.current+1), 2, "SH", "") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode a couple of contexts where scandinavian, slavic
 * or german names should get an alternate, native
 * pronunciation of 'SV' or 'XV'
 *
 * @return true if handled
 *
 */
func (m *M3) encode_Special_SW() bool {
	if m.current == 0 {
		//
		if m.names_Beginning_With_SW_That_Get_Alt_SV() {
			m.metaphAdd("S", "SV")
			m.current += 2
			return true
		}

		//
		if m.names_Beginning_With_SW_That_Get_Alt_XV() {
			m.metaphAdd("S", "XV")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode "-SKJ-" as X ("sh"), since americans pronounce
 * the name Dag Hammerskjold as "hammer-shold"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SKJ() bool {
	// scandinavian
	if m.stringAt(m.current, 4, "SKJO", "SKJU", "") && isVowel(m.charAt(m.current+3)) {
		m.metaphAdd("X", "X")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode initial swedish "SJ-" as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SJ() bool {
	if m.stringAt(0, 2, "SJ", "") {
		m.metaphAdd("X", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode final 'S' in words from the french, where they
 * are not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_French_S_Final() bool {
	// "louis" is an exception because it gets two pronuncuations
	if m.stringAt(0, 5, "LOUIS", "") && (m.current == m.last) {
		m.metaphAdd("S", "")
		m.current++
		return true
	}

	// french words familiar to americans where final s is silent
	if (m.current == m.last) && (m.stringAt(0, 4, "YVES", "") || (m.stringAt(0, 4, "HORS", "") && (m.current == 3)) || m.stringAt((m.current-4), 5, "CAMUS", "YPRES", "") || m.stringAt((m.current-5), 6, "MESNES", "DEBRIS", "BLANCS", "INGRES", "CANNES", "") || m.stringAt((m.current-6), 7, "CHABLIS", "APROPOS", "JACQUES", "ELYSEES", "OEUVRES",
		"GEORGES", "DESPRES", "") || m.stringAt(0, 8, "ARKANSAS", "FRANCAIS", "CRUDITES", "BRUYERES", "") || m.stringAt(0, 9, "DESCARTES", "DESCHUTES", "DESCHAMPS", "DESROCHES", "DESCHENES", "") || m.stringAt(0, 10, "RENDEZVOUS", "") || m.stringAt(0, 11, "CONTRETEMPS", "DESLAURIERS", "")) || ((m.current == m.last) && m.stringAt((m.current-2), 2, "AI", "OI", "UI", "") && !m.stringAt(0, 4, "LOIS", "LUIS", "")) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode non-final 'S' in words from the french where they
 * are not pronounced.
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_French_S_Internal() bool {
	// french words familiar to americans where internal s is silent
	if m.stringAt((m.current-2), 9, "DESCARTES", "") || m.stringAt((m.current-2), 7, "DESCHAM", "DESPRES", "DESROCH", "DESROSI", "DESJARD", "DESMARA",
		"DESCHEN", "DESHOTE", "DESLAUR", "") || m.stringAt((m.current-2), 6, "MESNES", "") || m.stringAt((m.current-5), 8, "DUQUESNE", "DUCHESNE", "") || m.stringAt((m.current-7), 10, "BEAUCHESNE", "") || m.stringAt((m.current-3), 7, "FRESNEL", "") || m.stringAt((m.current-3), 9, "GROSVENOR", "") || m.stringAt((m.current-4), 10, "LOUISVILLE", "") || m.stringAt((m.current-7), 10, "ILLINOISAN", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode silent 'S' in context of "-ISL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ISL() bool {
	//special cases 'island', 'isle', 'carlisle', 'carlysle'
	if (m.stringAt((m.current-2), 4, "LISL", "LYSL", "AISL", "") && !m.stringAt((m.current-3), 7, "PAISLEY", "BAISLEY", "ALISLAM", "ALISLAH", "ALISLAA", "")) || ((m.current == 1) && ((m.stringAt((m.current-1), 4, "ISLE", "") || m.stringAt((m.current-1), 5, "ISLAN", "")) && !m.stringAt((m.current-1), 5, "ISLEY", "ISLER", ""))) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-STL-" in contexts where the 'T' is silent. Also
 * encode "-USCLE-" in contexts where the 'C' is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_STL() bool {
	//'hustle', 'bustle', 'whistle'
	if m.stringAt(m.current, 4, "STLE", "STLI", "") && !m.stringAt((m.current+2), 4, "LESS", "LIKE", "LINE", "") || m.stringAt((m.current-3), 7, "THISTLY", "BRISTLY", "GRISTLY", "") ||
		// e.g. "corpuscle"
		m.stringAt((m.current-1), 5, "USCLE", "") {
		// KRISTEN, KRYSTLE, CRYSTLE, KRISTLE all pronounce the 't'
		// also, exceptions where "-LING" is a nominalizing suffix
		if m.stringAt(0, 7, "KRISTEN", "KRYSTLE", "CRYSTLE", "KRISTLE", "") || m.stringAt(0, 11, "CHRISTENSEN", "CHRISTENSON", "") || m.stringAt((m.current-3), 9, "FIRSTLING", "") || m.stringAt((m.current-2), 8, "NESTLING", "WESTLING", "") {
			m.metaphAdd("ST", "ST")
			m.current += 2
		} else {
			if m.encodeVowels && (m.charAt(m.current+3) == 'E') && (m.charAt(m.current+4) != 'R') && !m.stringAt((m.current+3), 4, "ETTE", "ETTA", "") && !m.stringAt((m.current+3), 2, "EY", "") {
				m.metaphAdd("SAL", "SAL")
				m.flag_AL_inversion = true
			} else {
				m.metaphAdd("SL", "SL")
			}
			m.current += 3
		}
		return true
	}

	return false
}

/**
 * Encode "christmas". Americans always pronounce this as "krissmuss"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Christmas() bool {
	//'christmas'
	if m.stringAt((m.current - 4), 8, "CHRISTMA", "") {
		m.metaphAdd("SM", "SM")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-STHM-" in contexts where the 'TH'
 * is silent.
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_STHM() bool {
	//'asthma', 'isthmus'
	if m.stringAt(m.current, 4, "STHM", "") {
		m.metaphAdd("SM", "SM")
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-ISTEN-" and "-STNT-" in contexts
 * where the 'T' is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ISTEN() bool {
	// 't' is silent in verb, pronounced in name
	if m.stringAt(0, 8, "CHRISTEN", "") {
		// the word itself
		if rootOrInflections(m.inWord, "CHRISTEN") || m.stringAt(0, 11, "CHRISTENDOM", "") {
			m.metaphAdd("S", "ST")
		} else {
			// e.g. 'christenson', 'christene'
			m.metaphAdd("ST", "ST")
		}
		m.current += 2
		return true
	}

	//e.g. 'glisten', 'listen'
	if m.stringAt((m.current-2), 6, "LISTEN", "RISTEN", "HASTEN", "FASTEN", "MUSTNT", "") || m.stringAt((m.current-3), 7, "MOISTEN", "") {
		m.metaphAdd("S", "S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special case "sugar"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Sugar() bool {
	//special case 'sugar-'
	if m.stringAt(m.current, 5, "SUGAR", "") {
		m.metaphAdd("X", "X")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-SH-" as X ("sh"), except in cases
 * where the 'S' and 'H' belong to different combining
 * roots and are therefore pronounced seperately
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SH() bool {
	if m.stringAt(m.current, 2, "SH", "") {
		// exception
		if m.stringAt((m.current - 2), 8, "CASHMERE", "") {
			m.metaphAdd("J", "J")
			m.current += 2
			return true
		}

		//combining forms, e.g. 'clotheshorse', 'woodshole'
		if (m.current > 0) &&
			// e.g. "mishap"
			((m.stringAt((m.current+1), 3, "HAP", "") && ((m.current + 3) == m.last)) ||
				// e.g. "hartsheim", "clothshorse"
				m.stringAt((m.current+1), 4, "HEIM", "HOEK", "HOLM", "HOLZ", "HOOD", "HEAD", "HEID",
					"HAAR", "HORS", "HOLE", "HUND", "HELM", "HAWK", "HILL", "") ||
				// e.g. "dishonor"
				m.stringAt((m.current+1), 5, "HEART", "HATCH", "HOUSE", "HOUND", "HONOR", "") ||
				// e.g. "mishear"
				(m.stringAt((m.current+2), 3, "EAR", "") && ((m.current + 4) == m.last)) ||
				// e.g. "hartshorn"
				(m.stringAt((m.current+2), 3, "ORN", "") && !m.stringAt((m.current-2), 7, "UNSHORN", "")) ||
				// e.g. "newshour" but not "bashour", "manshour"
				(m.stringAt((m.current+1), 4, "HOUR", "") && !(m.stringAt(0, 7, "BASHOUR", "") || m.stringAt(0, 8, "MANSHOUR", "") || m.stringAt(0, 6, "ASHOUR", ""))) ||
				// e.g. "dishonest", "grasshopper"
				m.stringAt((m.current+2), 5, "ARMON", "ONEST", "ALLOW", "OLDER", "OPPER", "EIMER", "ANDLE", "ONOUR", "") ||
				// e.g. "dishabille", "transhumance"
				m.stringAt((m.current+2), 6, "ABILLE", "UMANCE", "ABITUA", "")) {
			if !m.stringAt((m.current - 1), 1, "S", "") {
				m.metaphAdd("S", "S")
			}
		} else {
			m.metaphAdd("X", "X")
		}

		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-SCH-" in cases where the 'S' is pronounced
 * seperately from the "CH", in words from the dutch, italian,
 * and greek where it can be pronounced SK, and german words
 * where it is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SCH() bool {
	// these words were combining forms many centuries ago
	if m.stringAt((m.current + 1), 2, "CH", "") {
		if (m.current > 0) &&
			// e.g. "mischief", "escheat"
			(m.stringAt((m.current+3), 3, "IEF", "EAT", "") ||
				// e.g. "mischance"
				m.stringAt((m.current+3), 4, "ANCE", "ARGE", "") ||
				// e.g. "eschew"
				m.stringAt(0, 6, "ESCHEW", "")) {
			m.metaphAdd("S", "S")
			m.current++
			return true
		}

		//Schlesinger's rule
		//dutch, danish, italian, greek origin, e.g. "school", "schooner", "schiavone", "schiz-"
		if (m.stringAt((m.current+3), 2, "OO", "ER", "EN", "UY", "ED", "EM", "IA", "IZ", "IS", "OL", "") && !m.stringAt(m.current, 6, "SCHOLT", "SCHISL", "SCHERR", "")) || m.stringAt((m.current+3), 3, "ISZ", "") || (m.stringAt((m.current-1), 6, "ESCHAT", "ASCHIN", "ASCHAL", "ISCHAE", "ISCHIA", "") && !m.stringAt((m.current-2), 8, "FASCHING", "")) || (m.stringAt((m.current-1), 5, "ESCHI", "") && ((m.current + 3) == m.last)) || (m.charAt(m.current+3) == 'Y') {
			// e.g. "schermerhorn", "schenker", "schistose"
			if m.stringAt((m.current+3), 2, "ER", "EN", "IS", "") && (((m.current + 4) == m.last) || m.stringAt((m.current+3), 3, "ENK", "ENB", "IST", "")) {
				m.metaphAdd("X", "SK")
			} else {
				m.metaphAdd("SK", "SK")
			}
			m.current += 3
			return true
		} else {
			m.metaphAdd("X", "X")
			m.current += 3
			return true
		}
	}

	return false
}

/**
 * Encode "-SUR<E,A,Y>-" to J, unless it is at the beginning,
 * or preceeded by 'N', 'K', or "NO"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SUR() bool {
	// 'erasure', 'usury'
	if m.stringAt((m.current + 1), 3, "URE", "URA", "URY", "") {
		//'sure', 'ensure'
		if (m.current == 0) || m.stringAt((m.current-1), 1, "N", "K", "") || m.stringAt((m.current-2), 2, "NO", "") {
			m.metaphAdd("X", "X")
		} else {
			m.metaphAdd("J", "J")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-SU<O,A>-" to X ("sh") unless it is preceeded by
 * an 'R', in which case it is encoded to S, or it is
 * preceeded by a vowel, in which case it is encoded to J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SU() bool {
	//'sensuous', 'consensual'
	if m.stringAt((m.current+1), 2, "UO", "UA", "") && (m.current != 0) {
		// exceptions e.g. "persuade"
		if m.stringAt((m.current - 1), 4, "RSUA", "") {
			m.metaphAdd("S", "S")
		} else
		// exceptions e.g. "casual"
		if isVowel(m.charAt(m.current - 1)) {
			m.metaphAdd("J", "S")
		} else {
			m.metaphAdd("X", "S")
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encodes "-SSIO-" in contexts where it is pronounced
 * either J or X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SSIO() bool {
	if m.stringAt((m.current + 1), 4, "SION", "") {
		//"abcission"
		if m.stringAt((m.current - 2), 2, "CI", "") {
			m.metaphAdd("J", "J")
		} else
		//'mission'
		{
			if isVowel(m.charAt(m.current - 1)) {
				m.metaphAdd("X", "X")
			}
		}

		m.advanceCounter(4, 2)
		return true
	}

	return false
}

/**
 * Encode "-SS-" in contexts where it is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SS() bool {
	// e.g. "russian", "pressure"
	if m.stringAt((m.current-1), 5, "USSIA", "ESSUR", "ISSUR", "ISSUE", "") ||
		// e.g. "hessian", "assurance"
		m.stringAt((m.current-1), 6, "ESSIAN", "ASSURE", "ASSURA", "ISSUAB", "ISSUAN", "ASSIUS", "") {
		m.metaphAdd("X", "X")
		m.advanceCounter(3, 2)
		return true
	}

	return false
}

/**
 * Encodes "-SIA-" in contexts where it is pronounced
 * as X ("sh"), J, or S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SIA() bool {
	// e.g. "controversial", also "fuchsia", "ch" is silent
	if m.stringAt((m.current-2), 5, "CHSIA", "") || m.stringAt((m.current-1), 5, "RSIAL", "") {
		m.metaphAdd("X", "X")
		m.advanceCounter(3, 1)
		return true
	}

	// names generally get 'X' where terms, e.g. "aphasia" get 'J'
	if (m.stringAt(0, 6, "ALESIA", "ALYSIA", "ALISIA", "STASIA", "") && (m.current == 3) && !m.stringAt(0, 9, "ANASTASIA", "")) || m.stringAt((m.current-5), 9, "DIONYSIAN", "") || m.stringAt((m.current-5), 8, "THERESIA", "") {
		m.metaphAdd("X", "S")
		m.advanceCounter(3, 1)
		return true
	}

	if (m.stringAt(m.current, 3, "SIA", "") && ((m.current + 2) == m.last)) || (m.stringAt(m.current, 4, "SIAN", "") && ((m.current + 3) == m.last)) || m.stringAt((m.current-5), 9, "AMBROSIAL", "") {
		if (isVowel(m.charAt(m.current-1)) || m.stringAt((m.current-1), 1, "R", "")) &&
			// exclude compounds based on names, or french or greek words
			!(m.stringAt(0, 5, "JAMES", "NICOS", "PEGAS", "PEPYS", "") ||
				m.stringAt(0, 6, "HOBBES", "HOLMES", "JAQUES", "KEYNES", "") || m.stringAt(0, 7, "MALTHUS", "HOMOOUS", "") || m.stringAt(0, 8, "MAGLEMOS", "HOMOIOUS", "") || m.stringAt(0, 9, "LEVALLOIS", "TARDENOIS", "") || m.stringAt((m.current-4), 5, "ALGES", "")) {
			m.metaphAdd("J", "J")
		} else {
			m.metaphAdd("S", "S")
		}

		m.advanceCounter(2, 1)
		return true
	}
	return false
}

/**
 * Encodes "-SIO-" in contexts where it is pronounced
 * as J or X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SIO() bool {
	// special case, irish name
	if m.stringAt(0, 7, "SIOBHAN", "") {
		m.metaphAdd("X", "X")
		m.advanceCounter(3, 1)
		return true
	}

	if m.stringAt((m.current + 1), 3, "ION", "") {
		// e.g. "vision", "version"
		if isVowel(m.charAt(m.current-1)) || m.stringAt((m.current-2), 2, "ER", "UR", "") {
			m.metaphAdd("J", "J")
		} else { // e.g. "declension"
			m.metaphAdd("X", "X")
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases where "-S-" might well be from a german name
 * and add encoding of german pronounciation in alternate m.metaph
 * so that it can be found in a genealogical search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Anglicisations() bool {
	//german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
	//also, -sz- in slavic language altho in hungarian it is pronounced 's'
	if ((m.current == 0) && m.stringAt((m.current+1), 1, "M", "N", "L", "")) || m.stringAt((m.current+1), 1, "Z", "") {
		m.metaphAdd("S", "X")

		// eat redundant 'Z'
		if m.stringAt((m.current + 1), 1, "Z", "") {
			m.current += 2
		} else {
			m.current++
		}

		return true
	}

	return false
}

/**
 * Encode "-SC<vowel>-" in contexts where it is silent,
 * or pronounced as X ("sh"), S, or SK
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SC() bool {
	if m.stringAt(m.current, 2, "SC", "") {
		// exception 'viscount'
		if m.stringAt((m.current - 2), 8, "VISCOUNT", "") {
			m.current += 1
			return true
		}

		// encode "-SC<front vowel>-"
		if m.stringAt((m.current + 2), 1, "I", "E", "Y", "") {
			// e.g. "conscious"
			if m.stringAt((m.current+2), 4, "IOUS", "") ||
				// e.g. "prosciutto"
				m.stringAt((m.current+2), 3, "IUT", "") || m.stringAt((m.current-4), 9, "OMNISCIEN", "") ||
				// e.g. "conscious"
				m.stringAt((m.current-3), 8, "CONSCIEN", "CRESCEND", "CONSCION", "") || m.stringAt((m.current-2), 6, "FASCIS", "") {
				m.metaphAdd("X", "X")
			} else if m.stringAt(m.current, 7, "SCEPTIC", "SCEPSIS", "") || m.stringAt(m.current, 5, "SCIVV", "SCIRO", "") ||
				// commonly pronounced this way in u.s.
				m.stringAt(m.current, 6, "SCIPIO", "") || m.stringAt((m.current-2), 10, "PISCITELLI", "") {
				m.metaphAdd("SK", "SK")
			} else {
				m.metaphAdd("S", "S")
			}
			m.current += 2
			return true
		}

		m.metaphAdd("SK", "SK")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-S<EA,UI,IER>-" in contexts where it is pronounced
 * as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SEA_SUI_SIER() bool {
	// "nausea" by itself has => NJ as a more likely encoding. Other forms
	// using "nause-" (see encode_SEA()) have X or S as more familiar pronounciations
	if (m.stringAt((m.current-3), 6, "NAUSEA", "") && ((m.current + 2) == m.last)) ||
		// e.g. "casuistry", "frasier", "hoosier"
		m.stringAt((m.current-2), 5, "CASUI", "") || (m.stringAt((m.current-1), 5, "OSIER", "ASIER", "") && !(m.stringAt(0, 6, "EASIER", "") || m.stringAt(0, 5, "OSIER", "") || m.stringAt((m.current-2), 6, "ROSIER", "MOSIER", ""))) {
		m.metaphAdd("J", "X")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases where "-SE-" is pronounced as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_SEA() bool {
	if (m.stringAt(0, 4, "SEAN", "") && ((m.current + 3) == m.last)) || (m.stringAt((m.current-3), 6, "NAUSEO", "") && !m.stringAt((m.current-3), 7, "NAUSEAT", "")) {
		m.metaphAdd("X", "X")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-T-"
 *
 */
func (m *M3) encode_T() {
	if m.encode_T_Initial() || m.encode_TCH() || m.encode_Silent_French_T() || m.encode_TUN_TUL_TUA_TUO() || m.encode_TUE_TEU_TEOU_TUL_TIE() || m.encode_TUR_TIU_Suffixes() || m.encode_TI() || m.encode_TIENT() || m.encode_TSCH() || m.encode_TZSCH() || m.encode_TH_Pronounced_Separately() || m.encode_TTH() || m.encode_TH() {
		return
	}

	// eat redundant 'T' or 'D'
	if m.stringAt((m.current + 1), 1, "T", "D", "") {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("T", "T")
}

/**
 * Encode some exceptions for initial 'T'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_T_Initial() bool {
	if m.current == 0 {
		// americans usually pronounce "tzar" as "zar"
		if m.stringAt((m.current + 1), 3, "SAR", "ZAR", "") {
			m.current++
			return true
		}

		// old 'École française d'Extrême-Orient' chinese pinyin where 'ts-' => 'X'
		if ((m.length == 3) && m.stringAt((m.current+1), 2, "SO", "SA", "SU", "")) || ((m.length == 4) && m.stringAt((m.current+1), 3, "SAO", "SAI", "")) || ((m.length == 5) && m.stringAt((m.current+1), 4, "SING", "SANG", "")) {
			m.metaphAdd("X", "X")
			m.advanceCounter(3, 2)
			return true
		}

		// "TS<vowel>-" at start can be pronounced both with and without 'T'
		if m.stringAt((m.current+1), 1, "S", "") && isVowel(m.charAt(m.current+2)) {
			m.metaphAdd("TS", "S")
			m.advanceCounter(3, 2)
			return true
		}

		// e.g. "Tjaarda"
		if m.stringAt((m.current + 1), 1, "J", "") {
			m.metaphAdd("X", "X")
			m.advanceCounter(3, 2)
			return true
		}

		// cases where initial "TH-" is pronounced as T and not 0 ("th")
		if (m.stringAt((m.current+1), 2, "HU", "") && (m.length == 3)) || m.stringAt((m.current+1), 3, "HAI", "HUY", "HAO", "") || m.stringAt((m.current+1), 4, "HYME", "HYMY", "HANH", "") || m.stringAt((m.current+1), 5, "HERES", "") {
			m.metaphAdd("T", "T")
			m.advanceCounter(3, 2)
			return true
		}
	}

	return false
}

/**
 * Encode "-TCH-", reliably X ("sh", or in this case, "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TCH() bool {
	if m.stringAt((m.current + 1), 2, "CH", "") {
		m.metaphAdd("X", "X")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode the many cases where americans are aware that a certain word is
 * french and know to not pronounce the 'T'
 *
 * @return true if encoding handled in this routine, false if not
 * TOUCHET CHABOT BENOIT
 */
func (m *M3) encode_Silent_French_T() bool {
	// french silent T familiar to americans
	if ((m.current == m.last) && m.stringAt((m.current-4), 5, "MONET", "GENET", "CHAUT", "")) || m.stringAt((m.current-2), 9, "POTPOURRI", "") || m.stringAt((m.current-3), 9, "BOATSWAIN", "") || m.stringAt((m.current-3), 8, "MORTGAGE", "") || (m.stringAt((m.current-4), 5, "BERET", "BIDET", "FILET", "DEBUT", "DEPOT", "PINOT", "TAROT", "") || m.stringAt((m.current-5), 6, "BALLET", "BUFFET", "CACHET", "CHALET", "ESPRIT", "RAGOUT", "GOULET",
		"CHABOT", "BENOIT", "") || m.stringAt((m.current-6), 7, "GOURMET", "BOUQUET", "CROCHET", "CROQUET", "PARFAIT", "PINCHOT",
		"CABARET", "PARQUET", "RAPPORT", "TOUCHET", "COURBET", "DIDEROT", "") || m.stringAt((m.current-7), 8, "ENTREPOT", "CABERNET", "DUBONNET", "MASSENET", "MUSCADET", "RICOCHET", "ESCARGOT", "") || m.stringAt((m.current-8), 9, "SOBRIQUET", "CABRIOLET", "CASSOULET", "OUBRIQUET", "CAMEMBERT", "")) && !m.stringAt((m.current+1), 2, "AN", "RY", "IC", "OM", "IN", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-TU<N,L,A,O>-" in cases where it is pronounced
 * X ("sh", or in this case, "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TUN_TUL_TUA_TUO() bool {
	// e.g. "fortune", "fortunate"
	if m.stringAt((m.current-3), 6, "FORTUN", "") ||
		// e.g. "capitulate"
		(m.stringAt(m.current, 3, "TUL", "") && (isVowel(m.charAt(m.current-1)) && isVowel(m.charAt(m.current+3)))) ||
		// e.g. "obituary", "barbituate"
		m.stringAt((m.current-2), 5, "BITUA", "BITUE", "") ||
		// e.g. "actual"
		((m.current > 1) && m.stringAt(m.current, 3, "TUA", "TUO", "")) {
		m.metaphAdd("X", "T")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-T<vowel>-" forms where 'T' is pronounced as X
 * ("sh", or in this case "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TUE_TEU_TEOU_TUL_TIE() bool {
	// 'constituent', 'pasteur'
	if m.stringAt((m.current+1), 4, "UENT", "") || m.stringAt((m.current-4), 9, "RIGHTEOUS", "") || m.stringAt((m.current-3), 7, "STATUTE", "") || m.stringAt((m.current-3), 7, "AMATEUR", "") ||
		// e.g. "blastula", "pasteur"
		(m.stringAt((m.current - 1), 5, "NTULE", "NTULA", "STULE", "STULA", "STEUR", "")) ||
		// e.g. "statue"
		(((m.current + 2) == m.last) && m.stringAt(m.current, 3, "TUE", "")) ||
		// e.g. "constituency"
		m.stringAt(m.current, 5, "TUENC", "") ||
		// e.g. "statutory"
		m.stringAt((m.current-3), 8, "STATUTOR", "") ||
		// e.g. "patience"
		(((m.current + 5) == m.last) && m.stringAt(m.current, 6, "TIENCE", "")) {
		m.metaphAdd("X", "T")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-TU-" forms in suffixes where it is usually
 * pronounced as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TUR_TIU_Suffixes() bool {
	// 'adventure', 'musculature'
	if (m.current > 0) && m.stringAt((m.current+1), 3, "URE", "URA", "URI", "URY", "URO", "IUS", "") {
		// exceptions e.g. 'tessitura', mostly from romance languages
		if (m.stringAt((m.current+1), 3, "URA", "URO", "") &&
			//
			!m.stringAt((m.current+1), 4, "URIA", "") && ((m.current+3) == m.last)) && !m.stringAt((m.current-3), 7, "VENTURA", "") ||
			// e.g. "kachaturian", "hematuria"
			m.stringAt((m.current+1), 4, "URIA", "") {
			m.metaphAdd("T", "T")
		} else {
			m.metaphAdd("X", "T")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-TI<O,A,U>-" as X ("sh"), except
 * in cases where it is part of a combining form,
 * or as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TI() bool {
	// '-tio-', '-tia-', '-tiu-'
	// except combining forms where T already pronounced e.g 'rooseveltian'
	if (m.stringAt((m.current+1), 2, "IO", "") && !m.stringAt((m.current-1), 5, "ETIOL", "")) || m.stringAt((m.current+1), 3, "IAL", "") || m.stringAt((m.current-1), 5, "RTIUM", "ATIUM", "") || ((m.stringAt((m.current+1), 3, "IAN", "") && (m.current > 0)) && !(m.stringAt((m.current-4), 8, "FAUSTIAN", "") || m.stringAt((m.current-5), 9, "PROUSTIAN", "") || m.stringAt((m.current-2), 7, "TATIANA", "") || (m.stringAt((m.current-3), 7, "KANTIAN", "GENTIAN", "") || m.stringAt((m.current-8), 12, "ROOSEVELTIAN", ""))) || (((m.current + 2) == m.last) && m.stringAt(m.current, 3, "TIA", "") &&
		// exceptions to above rules where the pronounciation is usually X
		!(m.stringAt((m.current-3), 6, "HESTIA", "MASTIA", "") ||
			m.stringAt((m.current-2), 5, "OSTIA", "") || m.stringAt(0, 3, "TIA", "") || m.stringAt((m.current-5), 8, "IZVESTIA", ""))) || m.stringAt((m.current+1), 4, "IATE", "IATI", "IABL", "IATO", "IARY", "") || m.stringAt((m.current-5), 9, "CHRISTIAN", "")) {
		if ((m.current == 2) && m.stringAt(0, 4, "ANTI", "")) || m.stringAt(0, 5, "PATIO", "PITIA", "DUTIA", "") {
			m.metaphAdd("T", "T")
		} else if m.stringAt((m.current - 4), 8, "EQUATION", "") {
			m.metaphAdd("J", "J")
		} else {
			if m.stringAt(m.current, 4, "TION", "") {
				m.metaphAdd("X", "X")
			} else if m.stringAt(0, 5, "KATIA", "LATIA", "") {
				m.metaphAdd("T", "X")
			} else {
				m.metaphAdd("X", "T")
			}
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-TIENT-" where "TI" is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TIENT() bool {
	// e.g. 'patient'
	if m.stringAt((m.current + 1), 4, "IENT", "") {
		m.metaphAdd("X", "T")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-TSCH-" as X ("ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TSCH() bool {
	//'deutsch'
	if m.stringAt(m.current, 4, "TSCH", "") &&
		// combining forms in german where the 'T' is pronounced seperately
		!m.stringAt((m.current-3), 4, "WELT", "KLAT", "FEST", "") {
		// pronounced the same as "ch" in "chit" => X
		m.metaphAdd("X", "X")
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-TZSCH-" as X ("ch")
 *
 * "Neitzsche is peachy"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TZSCH() bool {
	//'neitzsche'
	if m.stringAt(m.current, 5, "TZSCH", "") {
		m.metaphAdd("X", "X")
		m.current += 5
		return true
	}

	return false
}

/**
 * Encodes cases where the 'H' in "-TH-" is the beginning of
 * another word in a combining form, special cases where it is
 * usually pronounced as 'T', and a special case where it has
 * become pronounced as X ("sh", in this case "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TH_Pronounced_Separately() bool {
	//'adulthood', 'bithead', 'apartheid'
	if ((m.current > 0) && m.stringAt((m.current+1), 4, "HOOD", "HEAD", "HEID", "HAND", "HILL", "HOLD",
		"HAWK", "HEAP", "HERD", "HOLE", "HOOK", "HUNT",
		"HUMO", "HAUS", "HOFF", "HARD", "") && !m.stringAt((m.current-3), 5, "SOUTH", "NORTH", "")) || m.stringAt((m.current+1), 5, "HOUSE", "HEART", "HASTE", "HYPNO", "HEQUE", "") ||
		// watch out for greek root "-thallic"
		(m.stringAt((m.current+1), 4, "HALL", "") && ((m.current + 4) == m.last) && !m.stringAt((m.current-3), 5, "SOUTH", "NORTH", "")) || (m.stringAt((m.current+1), 3, "HAM", "") && ((m.current + 3) == m.last) && !(m.stringAt(0, 6, "GOTHAM", "WITHAM", "LATHAM", "") || m.stringAt(0, 7, "BENTHAM", "WALTHAM", "WORTHAM", "") || m.stringAt(0, 8, "GRANTHAM", ""))) || (m.stringAt((m.current+1), 5, "HATCH", "") && !((m.current == 0) || m.stringAt((m.current-2), 8, "UNTHATCH", ""))) || m.stringAt((m.current-3), 7, "WARTHOG", "") ||
		// and some special cases where "-TH-" is usually pronounced 'T'
		m.stringAt((m.current-2), 6, "ESTHER", "") || m.stringAt((m.current-3), 6, "GOETHE", "") || m.stringAt((m.current-2), 8, "NATHALIE", "") {
		// special case
		if m.stringAt((m.current - 3), 7, "POSTHUM", "") {
			m.metaphAdd("X", "X")
		} else {
			m.metaphAdd("T", "T")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode the "-TTH-" in "matthew", eating the redundant 'T'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TTH() bool {
	// 'matthew' vs. 'outthink'
	if m.stringAt(m.current, 3, "TTH", "") {
		if m.stringAt((m.current - 2), 5, "MATTH", "") {
			m.metaphAdd("0", "0")
		} else {
			m.metaphAdd("T0", "T0")
		}
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-TH-". 0 (zero) is used in Metaphone to encode this sound
 * when it is pronounced as a dipthong, either voiced or unvoiced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_TH() bool {
	if m.stringAt(m.current, 2, "TH", "") {
		//'-clothes-'
		if m.stringAt((m.current - 3), 7, "CLOTHES", "") {
			// vowel already encoded so skip right to S
			m.current += 3
			return true
		}

		//special case "thomas", "thames", "beethoven" or germanic words
		if m.stringAt((m.current+2), 4, "OMAS", "OMPS", "OMPK", "OMSO", "OMSE",
			"AMES", "OVEN", "OFEN", "ILDA", "ILDE", "") || (m.stringAt(0, 4, "THOM", "") && (m.length == 4)) || (m.stringAt(0, 5, "THOMS", "") && (m.length == 5)) || m.stringAt(0, 4, "VAN ", "VON ", "") || m.stringAt(0, 3, "SCH", "") {
			m.metaphAdd("T", "T")

		} else {
			// give an 'etymological' 2nd
			// encoding for "smith"
			if m.stringAt(0, 2, "SM", "") {
				m.metaphAdd("0", "T")
			} else {
				m.metaphAdd("0", "0")
			}
		}

		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-V-"
 *
 */
func (m *M3) encode_V() {
	// eat redundant 'V'
	if m.charAt(m.current+1) == 'V' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAddExactApprox("V", "F")
}

/**
 * Encode "-W-"
 *
 */
func (m *M3) encode_W() {
	if m.encode_Silent_W_At_Beginning() || m.encode_WITZ_WICZ() || m.encode_WR() || m.encode_Initial_W_Vowel() || m.encode_WH() || m.encode_Eastern_European_W() {
		return
	}

	// e.g. 'zimbabwe'
	if m.encodeVowels && m.stringAt(m.current, 2, "WE", "") && ((m.current + 1) == m.last) {
		m.metaphAdd("A", "A")
	}

	//else skip it
	m.current++

}

/**
 * Encode cases where 'W' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Silent_W_At_Beginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(m.current, 2, "WR", "") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode polish patronymic suffix, mapping
 * alternate spellings to the same encoding,
 * and including easern european pronounciation
 * to the american so that both forms can
 * be found in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_WITZ_WICZ() bool {
	//polish e.g. 'filipowicz'
	if ((m.current + 3) == m.last) && m.stringAt(m.current, 4, "WICZ", "WITZ", "") {
		if m.encodeVowels {
			if (m.primary.Len() > 0) && m.primary.String()[m.primary.Len()-1] == 'A' {
				m.metaphAdd("TS", "FAX")
			} else {
				m.metaphAdd("ATS", "FAX")
			}
		} else {
			m.metaphAdd("TS", "FX")
		}
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-WR-" as R ('W' always effectively silent)
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_WR() bool {
	//can also be in middle of word
	if m.stringAt(m.current, 2, "WR", "") {
		m.metaphAdd("R", "R")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "W-", adding central and eastern european
 * pronounciations so that both forms can be found
 * in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_W_Vowel() bool {
	if (m.current == 0) && isVowel(m.charAt(m.current+1)) {
		//Witter should match Vitter
		if m.germanic_Or_Slavic_Name_Beginning_With_W() {
			if m.encodeVowels {
				m.metaphAddExactApprox4("A", "VA", "A", "FA")
			} else {
				m.metaphAddExactApprox4("A", "V", "A", "F")
			}
		} else {
			m.metaphAdd("A", "A")
		}

		m.current++
		// don't encode vowels twice
		m.current = m.skipVowels(m.current)
		return true
	}

	return false
}

/**
 * Encode "-WH-" either as H, or close enough to 'U' to be
 * considered a vowel
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_WH() bool {
	if m.stringAt(m.current, 2, "WH", "") {
		// cases where it is pronounced as H
		// e.g. 'who', 'whole'
		if (m.charAt(m.current+2) == 'O') &&
			// exclude cases where it is pronounced like a vowel
			!(m.stringAt((m.current+2), 4, "OOSH", "") || m.stringAt((m.current+2), 3, "OOP", "OMP", "ORL", "ORT", "") || m.stringAt((m.current+2), 2, "OA", "OP", "")) {
			m.metaphAdd("H", "H")
			m.advanceCounter(3, 2)
			return true
		} else {
			// combining forms, e.g. 'hollowhearted', 'rawhide'
			if m.stringAt((m.current+2), 3, "IDE", "ARD", "EAD", "AWK", "ERD",
				"OOK", "AND", "OLE", "OOD", "") || m.stringAt((m.current+2), 4, "EART", "OUSE", "OUND", "") || m.stringAt((m.current+2), 5, "AMMER", "") {
				m.metaphAdd("H", "H")
				m.current += 2
				return true
			} else if m.current == 0 {
				m.metaphAdd("A", "A")
				m.current += 2
				// don't encode vowels twice
				m.current = m.skipVowels(m.current)
				return true
			}
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-W-" when in eastern european names, adding
 * the eastern european pronounciation to the american so
 * that both forms can be found in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Eastern_European_W() bool {
	//Arnow should match Arnoff
	if ((m.current == m.last) && isVowel(m.charAt(m.current-1))) || m.stringAt((m.current-1), 5, "EWSKI", "EWSKY", "OWSKI", "OWSKY", "") || (m.stringAt(m.current, 5, "WICKI", "WACKI", "") && ((m.current + 4) == m.last)) || m.stringAt(m.current, 4, "WIAK", "") && ((m.current+3) == m.last) || m.stringAt(0, 3, "SCH", "") {
		m.metaphAddExactApprox4("", "V", "", "F")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-X-"
 *
 */
func (m *M3) encode_X() {
	if m.encode_Initial_X() || m.encode_Greek_X() || m.encode_X_Special_Cases() || m.encode_X_To_H() || m.encode_X_Vowel() || m.encode_French_X_Final() {
		return
	}

	// eat redundant 'X' or other redundant cases
	if m.stringAt((m.current+1), 1, "X", "Z", "S", "") ||
		// e.g. "excite", "exceed"
		m.stringAt((m.current+1), 2, "CI", "CE", "") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode initial X where it is usually pronounced as S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Initial_X() bool {
	// current chinese pinyin spelling
	if m.stringAt(0, 3, "XIA", "XIO", "XIE", "") || m.stringAt(0, 2, "XU", "") {
		m.metaphAdd("X", "X")
		m.current++
		return true
	}

	// else
	if m.current == 0 {
		m.metaphAdd("S", "S")
		m.current++
		return true
	}

	return false
}

/**
 * Encode X when from greek roots where it is usually pronounced as S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_Greek_X() bool {
	// 'xylophone', xylem', 'xanthoma', 'xeno-'
	if m.stringAt((m.current+1), 3, "YLO", "YLE", "ENO", "") || m.stringAt((m.current+1), 4, "ANTH", "") {
		m.metaphAdd("S", "S")
		m.current++
		return true
	}

	return false
}

/**
 * Encode special cases, "LUXUR-", "Texeira"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_X_Special_Cases() bool {
	// 'luxury'
	if m.stringAt((m.current - 2), 5, "LUXUR", "") {
		m.metaphAddExactApprox("GJ", "KJ")
		m.current++
		return true
	}

	// 'texeira' portuguese/galician name
	if m.stringAt(0, 7, "TEXEIRA", "") || m.stringAt(0, 8, "TEIXEIRA", "") {
		m.metaphAdd("X", "X")
		m.current++
		return true
	}

	return false
}

/**
 * Encode special case where americans know the
 * proper mexican indian pronounciation of this name
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_X_To_H() bool {
	// TODO: look for other mexican indian words
	// where 'X' is usually pronounced this way
	if m.stringAt((m.current-2), 6, "OAXACA", "") || m.stringAt((m.current-3), 7, "QUIXOTE", "") {
		m.metaphAdd("H", "H")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-X-" in vowel contexts where it is usually
 * pronounced KX ("ksh")
 * account also for BBC pronounciation of => KS
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_X_Vowel() bool {
	// e.g. "sexual", "connexion" (british), "noxious"
	if m.stringAt((m.current + 1), 3, "UAL", "ION", "IOU", "") {
		m.metaphAdd("KX", "KS")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases of "-X", encoding as silent when part
 * of a french word where it is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_French_X_Final() bool {
	//french e.g. "breaux", "paix"
	if !((m.current == m.last) && (m.stringAt((m.current-3), 3, "IAU", "EAU", "IEU", "") || m.stringAt((m.current-2), 2, "AI", "AU", "OU", "OI", "EU", ""))) {
		m.metaphAdd("KS", "KS")
	}

	return false
}

/**
 * Encode "-Z-"
 *
 */
func (m *M3) encode_Z() {
	if m.encode_ZZ() || m.encode_ZU_ZIER_ZS() || m.encode_French_EZ() || m.encode_German_Z() {
		return
	}

	if m.encode_ZH() {
		return
	} else {
		m.metaphAdd("S", "S")
	}

	// eat redundant 'Z'
	if m.charAt(m.current+1) == 'Z' {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode cases of "-ZZ-" where it is obviously part
 * of an italian word where "-ZZ-" is pronounced as TS
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ZZ() bool {
	// "abruzzi", 'pizza'
	if (m.charAt(m.current+1) == 'Z') && ((m.stringAt((m.current+2), 1, "I", "O", "A", "") && ((m.current + 2) == m.last)) || m.stringAt((m.current-2), 9, "MOZZARELL", "PIZZICATO", "PUZZONLAN", "")) {
		m.metaphAdd("TS", "S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special cases where "-Z-" is pronounced as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ZU_ZIER_ZS() bool {
	if ((m.current == 1) && m.stringAt((m.current-1), 4, "AZUR", "")) || (m.stringAt(m.current, 4, "ZIER", "") && !m.stringAt((m.current-2), 6, "VIZIER", "")) || m.stringAt(m.current, 3, "ZSA", "") {
		m.metaphAdd("J", "S")

		if m.stringAt(m.current, 3, "ZSA", "") {
			m.current += 2
		} else {
			m.current++
		}
		return true
	}

	return false
}

/**
 * Encode cases where americans recognize "-EZ" as part
 * of a french word where Z not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_French_EZ() bool {
	if ((m.current == 3) && m.stringAt((m.current-3), 4, "CHEZ", "")) || m.stringAt((m.current-5), 6, "RENDEZ", "") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode cases where "-Z-" is in a german word
 * where Z => TS in german
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_German_Z() bool {
	if ((m.current == 2) && ((m.current + 1) == m.last) && m.stringAt((m.current-2), 4, "NAZI", "")) || m.stringAt((m.current-2), 6, "NAZIFY", "MOZART", "") || m.stringAt((m.current-3), 4, "HOLZ", "HERZ", "MERZ", "FITZ", "") || (m.stringAt((m.current-3), 4, "GANZ", "") && !isVowel(m.charAt(m.current+1))) || m.stringAt((m.current-4), 5, "STOLZ", "PRINZ", "") || m.stringAt((m.current-4), 7, "VENEZIA", "") || m.stringAt((m.current-3), 6, "HERZOG", "") ||
		// german words beginning with "sch-" but not schlimazel, schmooze
		(strings.Contains(m.inWord, "SCH") && !(m.stringAt((m.last - 2), 3, "IZE", "OZE", "ZEL", ""))) || ((m.current > 0) && m.stringAt(m.current, 4, "ZEIT", "")) || m.stringAt((m.current-3), 4, "WEIZ", "") {
		if (m.current > 0) && m.charAt(m.current-1) == 'T' {
			m.metaphAdd("S", "S")
		} else {
			m.metaphAdd("TS", "TS")
		}
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-ZH-" as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *M3) encode_ZH() bool {
	//chinese pinyin e.g. 'zhao', also english "phonetic spelling"
	if m.charAt(m.current+1) == 'H' {
		m.metaphAdd("J", "J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test for names derived from the swedish,
 * dutch, or slavic that should get an alternate
 * pronunciation of 'SV' to match the native
 * version
 *
 * @return true if swedish, dutch, or slavic derived name
 */
func (m *M3) names_Beginning_With_SW_That_Get_Alt_SV() bool {
	if m.stringAt(0, 7, "SWANSON", "SWENSON", "SWINSON", "SWENSEN",
		"SWOBODA", "") || m.stringAt(0, 9, "SWIDERSKI", "SWARTHOUT", "") || m.stringAt(0, 10, "SWEARENGIN", "") {
		return true
	}

	return false
}

/**
 * Test for names derived from the german
 * that should get an alternate pronunciation
 * of 'XV' to match the german version spelled
 * "schw-"
 *
 * @return true if german derived name
 */
func (m *M3) names_Beginning_With_SW_That_Get_Alt_XV() bool {
	if m.stringAt(0, 5, "SWART", "") || m.stringAt(0, 6, "SWARTZ", "SWARTS", "SWIGER", "") || m.stringAt(0, 7, "SWITZER", "SWANGER", "SWIGERT",
		"SWIGART", "SWIHART", "") || m.stringAt(0, 8, "SWEITZER", "SWATZELL", "SWINDLER", "") || m.stringAt(0, 9, "SWINEHART", "") || m.stringAt(0, 10, "SWEARINGEN", "") {
		return true
	}

	return false
}

/**
 * Test whether the word in question
 * is a name of germanic or slavic origin, for
 * the purpose of determining whether to add an
 * alternate encoding of 'V'
 *
 * @return true if germanic or slavic name
 */
func (m *M3) germanic_Or_Slavic_Name_Beginning_With_W() bool {
	if m.stringAt(0, 3, "WEE", "WIX", "WAX", "") || m.stringAt(0, 4, "WOLF", "WEIS", "WAHL", "WALZ", "WEIL", "WERT",
		"WINE", "WILK", "WALT", "WOLL", "WADA", "WULF",
		"WEHR", "WURM", "WYSE", "WENZ", "WIRT", "WOLK",
		"WEIN", "WYSS", "WASS", "WANN", "WINT", "WINK",
		"WILE", "WIKE", "WIER", "WELK", "WISE", "") || m.stringAt(0, 5, "WIRTH", "WIESE", "WITTE", "WENTZ", "WOLFF", "WENDT",
		"WERTZ", "WILKE", "WALTZ", "WEISE", "WOOLF", "WERTH",
		"WEESE", "WURTH", "WINES", "WARGO", "WIMER", "WISER",
		"WAGER", "WILLE", "WILDS", "WAGAR", "WERTS", "WITTY",
		"WIENS", "WIEBE", "WIRTZ", "WYMER", "WULFF", "WIBLE",
		"WINER", "WIEST", "WALKO", "WALLA", "WEBRE", "WEYER",
		"WYBLE", "WOMAC", "WILTZ", "WURST", "WOLAK", "WELKE",
		"WEDEL", "WEIST", "WYGAN", "WUEST", "WEISZ", "WALCK",
		"WEITZ", "WYDRA", "WANDA", "WILMA", "WEBER", "") || m.stringAt(0, 6, "WETZEL", "WEINER", "WENZEL", "WESTER", "WALLEN", "WENGER",
		"WALLIN", "WEILER", "WIMMER", "WEIMER", "WYRICK", "WEGNER",
		"WINNER", "WESSEL", "WILKIE", "WEIGEL", "WOJCIK", "WENDEL",
		"WITTER", "WIENER", "WEISER", "WEXLER", "WACKER", "WISNER",
		"WITMER", "WINKLE", "WELTER", "WIDMER", "WITTEN", "WINDLE",
		"WASHER", "WOLTER", "WILKEY", "WIDNER", "WARMAN", "WEYANT",
		"WEIBEL", "WANNER", "WILKEN", "WILTSE", "WARNKE", "WALSER",
		"WEIKEL", "WESNER", "WITZEL", "WROBEL", "WAGNON", "WINANS",
		"WENNER", "WOLKEN", "WILNER", "WYSONG", "WYCOFF", "WUNDER",
		"WINKEL", "WIDMAN", "WELSCH", "WEHNER", "WEIGLE", "WETTER",
		"WUNSCH", "WHITTY", "WAXMAN", "WILKER", "WILHAM", "WITTIG",
		"WITMAN", "WESTRA", "WEHRLE", "WASSER", "WILLER", "WEGMAN",
		"WARFEL", "WYNTER", "WERNER", "WAGNER", "WISSER", "") || m.stringAt(0, 7, "WISEMAN", "WINKLER", "WILHELM", "WELLMAN", "WAMPLER", "WACHTER",
		"WALTHER", "WYCKOFF", "WEIDNER", "WOZNIAK", "WEILAND", "WILFONG",
		"WIEGAND", "WILCHER", "WIELAND", "WILDMAN", "WALDMAN", "WORTMAN",
		"WYSOCKI", "WEIDMAN", "WITTMAN", "WIDENER", "WOLFSON", "WENDELL",
		"WEITZEL", "WILLMAN", "WALDRUP", "WALTMAN", "WALCZAK", "WEIGAND",
		"WESSELS", "WIDEMAN", "WOLTERS", "WIREMAN", "WILHOIT", "WEGENER",
		"WOTRING", "WINGERT", "WIESNER", "WAYMIRE", "WHETZEL", "WENTZEL",
		"WINEGAR", "WESTMAN", "WYNKOOP", "WALLICK", "WURSTER", "WINBUSH",
		"WILBERT", "WALLACH", "WYNKOOP", "WALLICK", "WURSTER", "WINBUSH",
		"WILBERT", "WALLACH", "WEISSER", "WEISNER", "WINDERS", "WILLMON",
		"WILLEMS", "WIERSMA", "WACHTEL", "WARNICK", "WEIDLER", "WALTRIP",
		"WHETSEL", "WHELESS", "WELCHER", "WALBORN", "WILLSEY", "WEINMAN",
		"WAGAMAN", "WOMMACK", "WINGLER", "WINKLES", "WIEDMAN", "WHITNER",
		"WOLFRAM", "WARLICK", "WEEDMAN", "WHISMAN", "WINLAND", "WEESNER",
		"WARTHEN", "WETZLER", "WENDLER", "WALLNER", "WOLBERT", "WITTMER",
		"WISHART", "WILLIAM", "") || m.stringAt(0, 8, "WESTPHAL", "WICKLUND", "WEISSMAN", "WESTLUND", "WOLFGANG", "WILLHITE",
		"WEISBERG", "WALRAVEN", "WOLFGRAM", "WILHOITE", "WECHSLER", "WENDLING",
		"WESTBERG", "WENDLAND", "WININGER", "WHISNANT", "WESTRICK", "WESTLING",
		"WESTBURY", "WEITZMAN", "WEHMEYER", "WEINMANN", "WISNESKI", "WHELCHEL",
		"WEISHAAR", "WAGGENER", "WALDROUP", "WESTHOFF", "WIEDEMAN", "WASINGER",
		"WINBORNE", "") || m.stringAt(0, 9, "WHISENANT", "WEINSTEIN", "WESTERMAN", "WASSERMAN", "WITKOWSKI", "WEINTRAUB",
		"WINKELMAN", "WINKFIELD", "WANAMAKER", "WIECZOREK", "WIECHMANN", "WOJTOWICZ",
		"WALKOWIAK", "WEINSTOCK", "WILLEFORD", "WARKENTIN", "WEISINGER", "WINKLEMAN",
		"WILHEMINA", "") || m.stringAt(0, 10, "WISNIEWSKI", "WUNDERLICH", "WHISENHUNT", "WEINBERGER", "WROBLEWSKI",
		"WAGUESPACK", "WEISGERBER", "WESTERVELT", "WESTERLUND", "WASILEWSKI",
		"WILDERMUTH", "WESTENDORF", "WESOLOWSKI", "WEINGARTEN", "WINEBARGER",
		"WESTERBERG", "WANNAMAKER", "WEISSINGER", "") || m.stringAt(0, 11, "WALDSCHMIDT", "WEINGARTNER", "WINEBRENNER", "") || m.stringAt(0, 12, "WOLFENBARGER", "") || m.stringAt(0, 13, "WOJCIECHOWSKI", "") {
		return true
	}

	return false
}

/**
 * Test whether the word in question
 * is a name starting with 'J' that should
 * match names starting with a 'Y' sound.
 * All forms of 'John', 'Jane', etc, get
 * and alt to match e.g. 'Ian', 'Yana'. Joelle
 * should match 'Yael', 'Joseph' should match
 * 'Yusef'. German and slavic last names are
 * also included.
 *
 * @return true if name starting with 'J' that
 * should get an alternate encoding as a vowel
 */
func (m *M3) names_Beginning_With_J_That_Get_Alt_Y() bool {
	if m.stringAt(0, 3, "JAN", "JON", "JAN", "JIN", "JEN", "") || m.stringAt(0, 4, "JUHL", "JULY", "JOEL", "JOHN", "JOSH",
		"JUDE", "JUNE", "JONI", "JULI", "JENA",
		"JUNG", "JINA", "JANA", "JENI", "JOEL",
		"JANN", "JONA", "JENE", "JULE", "JANI",
		"JONG", "JOHN", "JEAN", "JUNG", "JONE",
		"JARA", "JUST", "JOST", "JAHN", "JACO",
		"JANG", "JUDE", "JONE", "") || m.stringAt(0, 5, "JOANN", "JANEY", "JANAE", "JOANA", "JUTTA",
		"JULEE", "JANAY", "JANEE", "JETTA", "JOHNA",
		"JOANE", "JAYNA", "JANES", "JONAS", "JONIE",
		"JUSTA", "JUNIE", "JUNKO", "JENAE", "JULIO",
		"JINNY", "JOHNS", "JACOB", "JETER", "JAFFE",
		"JESKE", "JANKE", "JAGER", "JANIK", "JANDA",
		"JOSHI", "JULES", "JANTZ", "JEANS", "JUDAH",
		"JANUS", "JENNY", "JENEE", "JONAH", "JONAS",
		"JACOB", "JOSUE", "JOSEF", "JULES", "JULIE",
		"JULIA", "JANIE", "JANIS", "JENNA", "JANNA",
		"JEANA", "JENNI", "JEANE", "JONNA", "") || m.stringAt(0, 6, "JORDAN", "JORDON", "JOSEPH", "JOSHUA", "JOSIAH",
		"JOSPEH", "JUDSON", "JULIAN", "JULIUS", "JUNIOR",
		"JUDITH", "JOESPH", "JOHNIE", "JOANNE", "JEANNE",
		"JOANNA", "JOSEFA", "JULIET", "JANNIE", "JANELL",
		"JASMIN", "JANINE", "JOHNNY", "JEANIE", "JEANNA",
		"JOHNNA", "JOELLE", "JOVITA", "JOSEPH", "JONNIE",
		"JANEEN", "JANINA", "JOANIE", "JAZMIN", "JOHNIE",
		"JANENE", "JOHNNY", "JONELL", "JENELL", "JANETT",
		"JANETH", "JENINE", "JOELLA", "JOEANN", "JULIAN",
		"JOHANA", "JENICE", "JANNET", "JANISE", "JULENE",
		"JOSHUA", "JANEAN", "JAIMEE", "JOETTE", "JANYCE",
		"JENEVA", "JORDAN", "JACOBS", "JENSEN", "JOSEPH",
		"JANSEN", "JORDON", "JULIAN", "JAEGER", "JACOBY",
		"JENSON", "JARMAN", "JOSLIN", "JESSEN", "JAHNKE",
		"JACOBO", "JULIEN", "JOSHUA", "JEPSON", "JULIUS",
		"JANSON", "JACOBI", "JUDSON", "JARBOE", "JOHSON",
		"JANZEN", "JETTON", "JUNKER", "JONSON", "JAROSZ",
		"JENNER", "JAGGER", "JASMIN", "JEPSEN", "JORDEN",
		"JANNEY", "JUHASZ", "JERGEN", "JAKOB", "") || m.stringAt(0, 7, "JOHNSON", "JOHNNIE", "JASMINE", "JEANNIE", "JOHANNA",
		"JANELLE", "JANETTE", "JULIANA", "JUSTINA", "JOSETTE",
		"JOELLEN", "JENELLE", "JULIETA", "JULIANN", "JULISSA",
		"JENETTE", "JANETTA", "JOSELYN", "JONELLE", "JESENIA",
		"JANESSA", "JAZMINE", "JEANENE", "JOANNIE", "JADWIGA",
		"JOLANDA", "JULIANE", "JANUARY", "JEANICE", "JANELLA",
		"JEANETT", "JENNINE", "JOHANNE", "JOHNSIE", "JANIECE",
		"JOHNSON", "JENNELL", "JAMISON", "JANSSEN", "JOHNSEN",
		"JARDINE", "JAGGERS", "JURGENS", "JOURDAN", "JULIANO",
		"JOSEPHS", "JHONSON", "JOZWIAK", "JANICKI", "JELINEK",
		"JANSSON", "JOACHIM", "JANELLE", "JACOBUS", "JENNING",
		"JANTZEN", "JOHNNIE", "") || m.stringAt(0, 8, "JOSEFINA", "JEANNINE", "JULIANNE", "JULIANNA", "JONATHAN",
		"JONATHON", "JEANETTE", "JANNETTE", "JEANETTA", "JOHNETTA",
		"JENNEFER", "JULIENNE", "JOSPHINE", "JEANELLE", "JOHNETTE",
		"JULIEANN", "JOSEFINE", "JULIETTA", "JOHNSTON", "JACOBSON",
		"JACOBSEN", "JOHANSEN", "JOHANSON", "JAWORSKI", "JENNETTE",
		"JELLISON", "JOHANNES", "JASINSKI", "JUERGENS", "JARNAGIN",
		"JEREMIAH", "JEPPESEN", "JARNIGAN", "JANOUSEK", "") || m.stringAt(0, 9, "JOHNATHAN", "JOHNATHON", "JORGENSEN", "JEANMARIE", "JOSEPHINA",
		"JEANNETTE", "JOSEPHINE", "JEANNETTA", "JORGENSON", "JANKOWSKI",
		"JOHNSTONE", "JABLONSKI", "JOSEPHSON", "JOHANNSEN", "JURGENSEN",
		"JIMMERSON", "JOHANSSON", "") || m.stringAt(0, 10, "JAKUBOWSKI", "") {
		return true
	}

	return false
}
