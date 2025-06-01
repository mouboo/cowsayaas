// spec declares a type that specifies what the user requests in terms of
// cow appearance. It is used by the handler and the renderer. It has a 
// constructor with some defaults.
package spec

type CowSpec struct {
	Text string
	Width int
	Eyes string
	Tongue string
	Mode string
	Think bool
}

func NewCowSpec() CowSpec {
	return CowSpec{
		Text:	"",
		Width:	40,
		Eyes:	"",
		Tongue:	"",
		Mode:	"default",
		Think:	false,
	}
}
