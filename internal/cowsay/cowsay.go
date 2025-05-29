package cowsay

import (
	"fmt"
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

func RenderCowsay(message string) string {

	// TODO: move to argument, no magic numbers
	bubbleMaxWidth := 25

	// Slice to hold the message lines
	messageLines := lineBreak(message, bubbleMaxWidth - 4)
	
	// Test
	fmt.Printf("%v\n", messageLines)

	// The string to return
	output := ""

	// Speech bubble width, in characters
	bubbleWidth := 22
	
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

	// Bottom
	builder.WriteRune(' ')
	for i := 0; i < bubbleWidth - 2; i++ {
		builder.WriteRune(bottomBorder)
	}
	builder.WriteRune(' ')
	builder.WriteRune('\n')

	output += builder.String()

	output += "\n"
	// Add cow
	output += fmt.Sprintf("The cow says: %v\n", message)
	
	return output
}
