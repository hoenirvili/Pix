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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func sendEnvelopeRequest(env interface{}, url string) []byte {
	xml, err := getEnvelope(env)

	if err != nil && len(xml) == 0 {
		ErrNow("Can't get the envlope content in your byte buffer")
	}

	xml = sanitizeEnvelope(xml)

	// make the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xml))
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")

	timeout := time.Duration(10 * time.Second)
	// Create new client that will do the request and after 10 seconds of waiting just timeout
	client := &http.Client{
		Timeout: timeout,
	}

	fmt.Fprintf(os.Stdout, "%s ===> %s\n", blue("REQUEST"), green(url))

	// Do the request
	response, err := client.Do(req)

	if err != nil || response.StatusCode >= http.StatusInternalServerError {
		fmt.Fprintf(os.Stderr, "%s ===> %s\n", yellow("TIMEOUT"), green(url))
		os.Exit(0)
	}

	defer response.Body.Close()

	fmt.Println(green("Response Status:"), blue(response.Status))
	fmt.Println(green("Reponse Headers"), blue(response.Header))

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ErrNow("Can't read body response ")
	}

	//sanitize the new body
	body = replaceAllEntities(body)

	return body
}
