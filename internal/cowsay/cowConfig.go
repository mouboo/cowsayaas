// spec declares a type that specifies what the user requests in terms of
// cow appearance. It is used by the handler and the renderer. It has a
// constructor with some defaults.
package cowsay

type CowConfig struct {
	Text   string	`json:"text"`
	Width  int		`json:"width"`
	Think  bool		`json:"think"`
	File   string	`json:"file"`
	Mode   string	`json:"mode"`
	Eyes   string	`json:"eyes"`
	Tongue string	`json:"tongue"`
}

func NewCowConfig() CowConfig {
	return CowConfig{
		Text:   "Moo!",
		Width:  40,
		Think:  false,
		File:   "default",
		Mode:   "",
		Eyes:   "",
		Tongue: "",
	}
}
