package cowsay

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode/utf8"
)

// RenderCowsay generates "ascii" art of a cow with a speech bubble based
// on a given CowSpec.
func RenderCowsay(c *CowConfig) (string, error) {

	// 1. Prepare the text by wordwrapping it within a maximum width,
	// and pad the other lines with spaces to have a uniform width.
	messageLines, width := formatText(c.Text, c.Width)


	// 2. Build the speech bubble with text. Borders depend on the number of lines.
	//  ________        _______        _________
	// < 1 line >      / two   \      / three   \
	//  --------       \ lines /      | lines   |
	//                  -------       \ or more /
	//                                 ---------
	var b strings.Builder

	// Top
	b.WriteRune(' ')
	b.WriteString(strings.Repeat("_", width+2))
	b.WriteString(" \n")

	// Middle (sides and text)           
	numLines := len(messageLines)
	leftBorder := "< "
	rightBorder := " >\n"
	for i, text := range messageLines {
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
		b.WriteString(text)
		b.WriteString(rightBorder)
	}

	// Bottom
	b.WriteRune(' ')
	b.WriteString(strings.Repeat("-", width+2))
	b.WriteString(" \n")

	speechbubble := b.String()

	// 3. Add cow, with possible template modifications
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
	tmplBytes, err := os.ReadFile(filepath.Join(cowsdir, c.File+".cow"))
	if err != nil {
		return "", fmt.Errorf("Failed to read cowfile: %w", err)
	}
	
	// Parse the template
	tmpl, err := template.New("cow").Parse(string(tmplBytes))
	if err != nil {
		return "", fmt.Errorf("Failed to parse template: %w", err)
	}

	// Create buffer for the rendered cow
	var cowBuf bytes.Buffer

	// Execute the template with the CowConfig as input data
	err = tmpl.Execute(&cowBuf, c)
	if err != nil {
		return "", fmt.Errorf("Failed to execute template: %w", err)
	}

	// Add cow to output
	cow := cowBuf.String()

	// Put it all together
	output := speechbubble + cow

	return output, nil
}

// formatText() takes a string and an int. It splits the string into a slice of
// strings where each string fits in max length.
func formatText(text string, width int) ([]string, int) {
	// Split the string into words. If a single word is longer than max,
	// break it into max sized chunks.
	var words []string
	for _, word := range strings.Fields(text) {
		runes := []rune(word)
		for len(runes) > width {
			words = append(words, string(runes[:width]))
			runes = runes[width:]
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
		// if there's enough room in the current line, add a word
		lineLen := utf8.RuneCountInString(currentLine)
		wordLen := utf8.RuneCountInString(word)
		if lineLen + wordLen + spaceNeeded <= width {
			// insert either 0 or 1 spaces
			currentLine += strings.Repeat(" ", spaceNeeded)
			// insert the current word
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
	
	// Find length of the longest line after wordwrapping
	longestLine := 0
	for _, l := range lines {
		if utf8.RuneCountInString(l) > longestLine {
			longestLine = utf8.RuneCountInString(l)
		}
	}
	
	// Pad the shorter lines with spaces so all lines are equal length
	for i, l := range lines {
		spacesToAdd := longestLine - utf8.RuneCountInString(l)
		lines[i] += strings.Repeat(" ", spacesToAdd)
	}
	
	return lines, longestLine
}
