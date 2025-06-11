package cowsay

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// RenderCowsay generates "ascii" art of a cow with a speech bubble based
// on a given CowSpec.
func RenderCowsay(c CowConfig) (string, error) {

	// Top part - the speech bubble
	// Slice of string to hold the message lines
	messageLines := lineBreak(c.Text, c.Width)

	// Update/shrink width if the longest line is shorter
	// than width
	longestLine := 0
	for _, l := range messageLines {
		if len([]rune(l)) > longestLine {
			longestLine = len([]rune(l))
		}
	}
	c.Width = longestLine

	// Pad shorter lines with trailing spaces so all lines are
	// the same length
	for i, l := range messageLines {
		spacesToAdd := longestLine - len([]rune(l))
		messageLines[i] += strings.Repeat(" ", spacesToAdd)
	}

	// Build the speech bubble with text
	var b strings.Builder

	// Top
	topBorder := "_"
	b.WriteRune(' ')
	b.WriteString(strings.Repeat(topBorder, c.Width+2))
	b.WriteString(" \n")

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
			} else if i == numLines-1 {
				leftBorder = "\\ "
				rightBorder = " /\n"
			} else {
				leftBorder = "| "
				rightBorder = " |\n"
			}
		}
		b.WriteString(leftBorder)
		b.WriteString(l)
		b.WriteString(rightBorder)
	}

	// Bottom
	bottomBorder := "-"
	b.WriteRune(' ')
	b.WriteString(strings.Repeat(bottomBorder, c.Width+2))
	b.WriteString(" \n")

	output := b.String()

	// Add cow, with possible template modifications
	//        \   ^__^
	//         \  (oo)\_______
	//            (__)\       )\/\
	//                ||----w |
	//                ||     ||

	// Modify cowspec fields based on stated mode
	switch c.Mode {
	case "borg":
		if c.Eyes == "" {
			c.Eyes = "=="
		}
	case "dead":
		if c.Eyes == "" {
			c.Eyes = "xx"
		}
		if c.Tongue == "" {
			c.Tongue = "U"
		}
	case "greedy":
		if c.Eyes == "" {
			c.Eyes = "$$"
		}
	case "paranoia":
		if c.Eyes == "" {
			c.Eyes = "@@"
		}
	case "stoned":
		if c.Eyes == "" {
			c.Eyes = "**"
		}
		if c.Tongue == "" {
			c.Tongue = "U"
		}
	case "tired":
		if c.Eyes == "" {
			c.Eyes = "--"
		}
	case "wired":
		if c.Eyes == "" {
			c.Eyes = "OO"
		}
	case "youthful":
		if c.Eyes == "" {
			c.Eyes = ".."
		}
	}

	// Load cow template file named in c.File, defaults to "default"
	cowsdir := "./data/cows"
	templateBytes, err := os.ReadFile(filepath.Join(cowsdir, c.File+".cow"))
	if err != nil {
		return "", fmt.Errorf("Failed to read cowfile: %w", err)
	}

	// Parse the template
	template, err := template.New("cow").Parse(string(templateBytes))
	if err != nil {
		return "", fmt.Errorf("Failed to parse template: %w", err)
	}

	// Create buffer for the rendered cow
	var cowBuf bytes.Buffer

	// Execute the template with the CowSpec as input data
	err = template.Execute(&cowBuf, c)
	if err != nil {
		return "", fmt.Errorf("Failed to execute template: %w", err)
	}

	// Add cow to output
	output += cowBuf.String()

	return output, nil
}

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
		if len([]rune(currentLine))+len([]rune(word))+spaceNeeded <= max {
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
