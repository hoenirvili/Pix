package lib

import (
	"fmt"
	"os"
)

// Checks if the content is empty
func isEmpty(buffer []byte) {
	if buffer == nil {
		fmt.Fprintf(os.Stderr, "Error : %s\n", red("Empty stream of data"))
		os.Exit(1)
	}
}
