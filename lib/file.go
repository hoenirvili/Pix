//   Copyright 8-11-2015 Hoenir
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Gets the content of the file specified by
// a path var and return the content uint8 type
func fileContent(path string) []byte {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		ErrNow(fmt.Sprintf("%s %s", "Can't read from ", path))
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
