// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

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
