package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Envelope struct {
	XMLName    xml.Name `xml:soapenv:Envelope"`
	Key1       string   `xml:"xmlns:soapenv,attr"`
	Value1     string   `xml:",chardata"`
	Key2       string   `xml:"xmlns:web,attr"`
	Value2     string   `xml:",chardata"`
	CreateBody Body     `xml:"soapenv:Body"`
}
type Body struct {
	CreateText Text `xml:"soapenv:parseText_XML"`
}

type Text struct {
	TextRow []byte `xml:"web:rawTextInput"`
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
	e.Key1 = "http://schemas.xmlsoap.org/soap/envelope/"
	e.Key2 = "http://webPosRo.uaic/"
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

	replace := bytes.Replace(buff, wrong, open, -1)

	newReplace := bytes.Replace(replace, eofWrong, close, -1)

	return newReplace

}

func sendEnvelopeRequest(env *Envelope, url string) {
	fmt.Printf("Sending request to %s this urs !\n", url)

	xmlEnvelope, err := getEnvelope(*env)

	if err != nil && len(xmlEnvelope) == 0 {
		fmt.Println("Can't get the envlope content in your byte buffer")
		panic(err)
	}

	xmlEnvelope = sanitizeXml(xmlEnvelope)
	fmt.Println(string(xmlEnvelope))
}
