package cowsay

import (
	"fmt"
)

func RenderCowsay(message string) string {
	output := fmt.Sprintf("The cow says: %v\n", message)
	return output
}
