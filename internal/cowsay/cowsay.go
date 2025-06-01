package cowsay

import (
	"strings"
	
	//"github.com/mouboo/cowsayaas/internal/spec"
)

// lineBreak() takes a string and an int. It splits the string into a slice of
// string where each string fits in max length
func lineBreak(s string, max int) []string {
	// Split the string into words. If a single word is longer than max, 
	// break it into max sized chunks.
	var words []string
	for _, word := range strings.Fields(s) {
		runes := []rune(word)
		for len(runes) > max {
			words = append(words, string(runes[:max]))
			runes = runes[max:]
		}
		if len(runes) > 0 {
			words = append(words, string(runes))
		}
	}
	
	// Build up lines of allowed length and append to a new slice of string.
	var lines []string
	var currentLine string
	for _, word := range words {
		// words are appended with a space in front, unless it's the first word
		// of a line
		spaceNeeded := 1
		if currentLine == "" {
			spaceNeeded = 0
		}
		// if there's space in the current line, add a word
		if len([]rune(currentLine)) + len([]rune(word)) + spaceNeeded <= max {
			currentLine += strings.Repeat(" ", spaceNeeded)
			currentLine += word
		} else {
			// if a new line is needed
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}
	// add remaining words in underfull last line
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func RenderCowsay(message string, width int) string {

	// Slice of string to hold the message lines
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
	
	// Build the speech bubble with text
	var builder strings.Builder

	// Top
	topBorder := "_"
	builder.WriteRune(' ')
	builder.WriteString(strings.Repeat(topBorder, width + 2))
	builder.WriteString(" \n")
	
	// Lines with text. Borders depend on the number of lines.
	//  ________        _______        _________
	// < 1 line >      / two   \      / three   \
	//  --------       \ lines /      | lines   |
	//                  -------       \ or more /
	//                                 --------- 
	numLines := len(messageLines)
	leftBorder := "< "
	rightBorder := " >\n"
	for i, l := range messageLines {
		if numLines > 1 {
			if i == 0 {
				leftBorder = "/ "
				rightBorder = " \\\n"
			} else if i == numLines - 1 {
				leftBorder = "\\ "
				rightBorder = " /\n"
			} else {
				leftBorder = "| "
				rightBorder = " |\n"
			}
		}
		builder.WriteString(leftBorder)
		builder.WriteString(l)
		builder.WriteString(rightBorder)
	}
	
	// Bottom
	bottomBorder := "-"
	builder.WriteRune(' ')
	builder.WriteString(strings.Repeat(bottomBorder, width + 2))
	builder.WriteString(" ")

	output := builder.String()

	// Add cow, TODO: configurable with variables or cowfiles?
	output += `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
// Notes about original cowsay(1) options:
// width (default 40)
// eyes (user input"
// tongue (user input)
// think (boolean)
// cowfile
// mode: borg, dead, greedy, paranoia, stoned, tired, wired, youthful
// 		 (if mode is supplied, eyes and tongue parameters are ignored)
	return output
}
