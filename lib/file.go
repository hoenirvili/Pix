// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lib

import (
	"io/ioutil"
	"os"
)

// Gets the content of the file specified by
// a path var and return the content uint8 type
func fileContent(path string) []byte {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		ErrNow("Can't read from file")
	}

	return buffer
}

// Creates the file on path  and fill the file with
// the specific content
func saveFile(path string, buffer string) {

	// create file
	fd, err := os.Create(path)

	if err != nil {
		ErrNow("Can't create file here")
	}

	defer fd.Close()

	if n, err := fd.WriteString(buffer); err != nil && n <= 0 {
		ErrNow("Can't write to file")
	}
}
