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

import "bytes"

func replaceAllEntities(buff []byte) []byte {

	var entities = map[string]string{
		" ":  "&nbsp;",
		"\"": "&quot;",
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
	}
	for i := 0; i < 2; i++ {
		for key, val := range entities {
			buff = bytes.Replace(buff, []byte(val), []byte(key), -1)
		}
	}
	return buff
}

func sanitizeEnvelope(buffer []byte) []byte {

	var (
		StartF = []byte("<Envelope")
		FinalF = []byte("</Envelope>")
		StartT = []byte("<soapenv:Envelope ")
		FinalT = []byte("</soapenv:Envelope>")
	)

	// Check all the bytes  equal to StartF and FinalF
	// And replace all with StartT and FinalT
	buffer = bytes.Replace(buffer, StartF, StartT, -1)
	buffer = bytes.Replace(buffer, FinalF, FinalT, -1)

	// return the new sanitize envelope buffer
	return buffer
}
