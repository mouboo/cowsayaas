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
		currentWordLength := len([]rune(w))
		currentSpaceLeft := max - len([]rune(currentLine))
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
			lines = append(lines, currentLine)
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

	// (Maximum) speech bubble width, in runes
	bubbleWidth := width

	// Slice of string to hold the message lines, 4 runes reserved 
	// for the bubble left and right borders
	messageLines := lineBreak(message, bubbleWidth - 4)
	
	// Update/shrink bubbleWidth if the longest line is shorter
	// than bubbleWidth
	longestLine := 0
	for _, l := range messageLines {
		if len([]rune(l)) > longestLine {
			longestLine = len([]rune(l))
		}
	}
	bubbleWidth = longestLine + 4
	
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
	for i := 0; i < bubbleWidth - 2; i++ {
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
	for i := 0; i < bubbleWidth - 2; i++ {
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
