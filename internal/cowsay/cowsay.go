package cowsay

import (
	"strings"
)

func lineBreak(s string, max int) []string {
	words := strings.Fields(s)

	// First deal with possible words that are longer than
	// max length and need to be broken up.
	tmpWords := []string{}
	for _, w := range words {
		rs := []rune(w)
		for len(rs) > max {
			tmpWords = append(tmpWords, string(rs[:max]))
			rs = rs[max:]
		}
		if len(rs) > 0 {
			tmpWords = append(tmpWords, string(rs))
		}
	}
	words = tmpWords
	
	// Now all words are less than or equal to max line length.
	lines := []string{}
	currentLine := ""
	for _, w := range words {
		currentWordLength := len([]rune(w)) // 3
		currentSpaceLeft := max - len([]rune(currentLine)) // 0
		// See if there is room for another word
		// (and a space)
		if currentWordLength <= currentSpaceLeft-1 {
			if len(currentLine) != 0 {
				currentLine += " "
			}
			currentLine += w
		} else {
			// Not enough space to add more words.
			// Push to lines, and reset currentLine
			// to only contain the current word.
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = w
		}
	}
	// For the last unfinished line, push to lines
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func RenderCowsay(message string, width int) string {

	// Slice of string to hold the message lines, 4 runes reserved 
	// for the bubble left and right borders
	messageLines := lineBreak(message, width)
	
	// Update/shrink width if the longest line is shorter
	// than width
	longestLine := 0
	for _, l := range messageLines {
		if len([]rune(l)) > longestLine {
			longestLine = len([]rune(l))
		}
	}
	width = longestLine
	
	// Pad shorter lines with trailing spaces so all lines are
	// the same length
	for i, l := range messageLines {
		spacesToAdd := longestLine - len([]rune(l))
		messageLines[i] += strings.Repeat(" ", spacesToAdd)
	}

	// Start assembling the string to return
	output := ""

	// Speech bubble elements
	topBorder := '_'
	//leftBorder := '<'
	//rightBorder := '>'
	bottomBorder := '-'
	
	// Build the speech bubble with text
	var builder strings.Builder

	// Top
	builder.WriteRune(' ')
	for i := 0; i < width + 2; i++ {
		builder.WriteRune(topBorder)
	}
	builder.WriteRune(' ')
	builder.WriteRune('\n')
	
	// Lines of text
	for _, l := range messageLines {
		builder.WriteString("< ")
		builder.WriteString(l)
		builder.WriteString(" >")
		builder.WriteRune('\n')
	}
	
	// Bottom
	builder.WriteRune(' ')
	for i := 0; i < width + 2; i++ {
		builder.WriteRune(bottomBorder)
	}
	builder.WriteRune(' ')

	output += builder.String()

	// Add cow, TODO: configurable with variables or cowfiles?
	output += `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
		
	return output
}
