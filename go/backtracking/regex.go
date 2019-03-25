package backtracking

func regexMatch(text, pattern string, textIndex, patternIndex int, matched *bool) {
	if *matched {
		return
	}

	textLength := len(text)
	patternLength := len(pattern)

	if patternIndex == patternLength && textIndex == textLength {
		*matched = true
		return
	}

	if patternIndex >= patternLength {
		return
	}

	if pattern[patternIndex] == '*' {
		for k := 0; k <= textLength-textIndex; k++ {
			regexMatch(text, pattern, textIndex+k, patternIndex+1, matched)
		}
	} else if pattern[patternIndex] == '?' {
		regexMatch(text, pattern, textIndex, patternIndex+1, matched)
		regexMatch(text, pattern, textIndex+1, patternIndex+1, matched)
	} else if textIndex < textLength && text[textIndex] == pattern[patternIndex] {
		regexMatch(text, pattern, textIndex+1, patternIndex+1, matched)
	}
}
