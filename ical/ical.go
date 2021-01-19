package ical

import "unicode"

// Internal fold content line using recursive strategy
func _foldContentLine(line string, maxLen int) []string {
	folded := []string{}
	lineLen := len(line)

	if lineLen > maxLen {
		lineHalfLen := lineLen / 2
		folded = append(folded, _foldContentLine(line[:lineHalfLen], maxLen)...)
		folded = append(folded, _foldContentLine(line[lineHalfLen:], maxLen)...)
	} else {
		folded = append(folded, line)
	}

	return folded
}

// FoldContentLine folds a line if it is greater than max length given (https://tools.ietf.org/html/rfc2445)
func FoldContentLine(line string, maxLen int) []string {
	substrs := _foldContentLine(line, maxLen)

	// Append fold signature to all but last line
	for i, s := range substrs[:len(substrs)-1] {
		substrs[i] = s + "\r\n\t"
	}

	return substrs
}

// UnfoldContentLines unfolds content lines into one string (https://tools.ietf.org/html/rfc2445)
func UnfoldContentLines(lines []string) string {
	unfolded := ""

	for _, line := range lines {
		lineLength := len(line)

		// Check for line fold signature CRLF + whitespace
		if lineLength >= 3 && line[lineLength-3] == '\r' && line[lineLength-2] == '\n' && unicode.IsSpace(rune(line[lineLength-1])) {
			unfolded += line[:lineLength-3]
		} else {
			unfolded += line
		}
	}

	return unfolded
}
