package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func sendEnvelopeRequest(env *Envelope, url string) []byte {
	fmt.Fprintf(os.Stdout, "%s ===> %s\n", blue("Sending request"), green(url))

	xml, err := getEnvelope(*env)

	if err != nil && len(xml) == 0 {
		ErrNow("Can't get the envlope content in your byte buffer")
	}

	xml = sanitizeEnvelope(xml)
	// make the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xml))
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")

	// Create new client that will do the request
	client := &http.Client{}
	// Do the request
	response, err := client.Do(req)
	if err != nil {
		ErrNow("Client can't do the specific request")
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
