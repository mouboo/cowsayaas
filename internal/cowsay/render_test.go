package cowsay

import (
	"testing"
)

func TestFormatText(t *testing.T) {
	tests := []struct {
		text string
		width int
		wantedLines []string
		wantedWidth int
	}{
		{
			text: "hello world",
			width: 5,
			wantedLines: []string{
				"hello",
				"world",
			},
			wantedWidth: 5,
		},
	}
	
	for _, tt := range tests {
		lines, w := formatText(tt.text, tt.width)
		if w != tt.wantedWidth {
			t.Errorf("wanted width %v, got %v", tt.wantedWidth, w)
		}
		for i := range lines {
			if lines[i] != tt.wantedLines[i] {
				t.Errorf("line %v: wanted %v, got %v", i, tt.wantedLines[i], lines[i])
			}
		}
	}
}
