// spec declares a type that specifies what the user requests in terms of
// cow appearance. It is used by the handler and the renderer. It has a
// constructor with some defaults.
package spec

type CowSpec struct {
	Text   string
	Width  int
	Think  bool
	File   string
	Mode   string
	Eyes   string
	Tongue string
}

func NewCowSpec() CowSpec {
	return CowSpec{
		Text:   "",
		Width:  40,
		Think:  false,
		File:   "default",
		Mode:   "",
		Eyes:   "",
		Tongue: "",
	}
}
