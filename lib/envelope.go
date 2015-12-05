package lib

import "encoding/xml"

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

func newEnvelope(path string) (*Envelope, error) {
	// create new envelope
	var env Envelope

	buffSliceEnv := fileContent(path)

	err := xml.Unmarshal(buffSliceEnv, &env)

	return &env, err
}

func (e *Envelope) addContent(buffer []byte) {
	e.CreateBody.CreateText.TextRow = buffer
	e.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
	e.Val2 = "http://webPosRo.uaic/"
}

// get xmlEnvelope
func getEnvelope(env Envelope) ([]byte, error) {
	buff, err := xml.MarshalIndent(env, "", "   ")

	if err != nil {
		return nil, err
	}

	return buff, nil
}
