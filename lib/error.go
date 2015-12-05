package lib

import (
	"fmt"
	"os"
)

func ErrNow(message string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", red(message))
}
