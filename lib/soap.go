package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Envelope struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:web,attr"`
	CreateBody Body     `xml:"soapenv:Body"`
}
type Body struct {
	CreateText Text `xml:"web:parseText_XML"`
}

type Text struct {
	TextRow []byte `xml:"rawTextInput"`
}

func setEnvelope(path string) (*Envelope, error) {
	// create new envelope
	var env Envelope

	buffSliceEnv := retGetFileCnt(path)

	err := xml.Unmarshal(buffSliceEnv, &env)

	if err != nil {
		return &env, err
	}

	return &env, nil
}

func (e *Envelope) addCntToEnv(cnt []byte) {
	e.CreateBody.CreateText.TextRow = cnt
	e.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
	e.Val2 = "http://webPosRo.uaic/"
}

// get xmlEnvelope
func getEnvelope(env Envelope) ([]byte, error) {
	buff, err := xml.MarshalIndent(env, "", "   ")

	if err != nil {
		return []byte(""), err
	}

	return buff, nil
}

func sanitizeXml(buff []byte) []byte {
	var wrong = []byte("<Envelope")
	var eofWrong = []byte("</Envelope>")

	var open = []byte("<soapenv:Envelope ")
	var close = []byte("</soapenv:Envelope>")

	buff = bytes.Replace(buff, wrong, open, -1)

	buff = bytes.Replace(buff, eofWrong, close, -1)

	return buff

}

func sendEnvelopeRequest(env *Envelope, url string) ([]byte, []rune) {
	fmt.Printf("Sending request to %s this urs !\n", url)

	xmlEnvelope, err := getEnvelope(*env)

	if err != nil && len(xmlEnvelope) == 0 {
		fmt.Println("Can't get the envlope content in your byte buffer")
		panic(err)
	}

	xmlEnvelope = sanitizeXml(xmlEnvelope)

	fmt.Println(string(xmlEnvelope))

	resp, err := http.Post(url, "text/xml;charset=UTF-8", bytes.NewReader(xmlEnvelope))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server not responding\n")
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newRune := bytes.Runes(b)
	return b, newRune

}
