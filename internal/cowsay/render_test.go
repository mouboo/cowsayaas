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
			width: 6,
			wantedLines: []string{
				"hello",
				"world",
			},
			wantedWidth: 5,
		},
		{
			text: "helloworldloremipsumdolorsitamet",
			width: 5,
			wantedLines: []string{
				"hello",
				"world",
				"lorem",
				"ipsum",
				"dolor",
				"sitam",
				"et   ",
			},
			wantedWidth: 5,
		},
		{
			text: "",
			width: 8,
			wantedLines: []string{
				"",
			},
			wantedWidth: 0,
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

func TestRenderCowsay(t *testing.T) {
	CowfileDir = "../../data/cows"
	c := CowConfig{
		Text:   "Moo!",
		Width:  40,
		Think:  false,
		File:   "default",
		Mode:   "",
		Eyes:   "",
		Tongue: "",
	}
	want := ` ______ 
< Moo! >
 ------ 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
	got, err := RenderCowsay(&c)
	if err != nil {
		t.Fatalf("RenderCowsay returned error: %v", err)
	}
	if want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}
