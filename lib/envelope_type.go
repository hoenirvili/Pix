package lib

import "encoding/xml"

// POS TAGGER ENVELOPE SERVICE

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

///////////////////////////////////////////////

//CHUNK TAGGER ENVELOPE SERVICE
type Envelope1 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:web,attr"`
	CreateBody Body1    `xml:"soapenv:Body"`
}

type Body1 struct {
	CreateText Text1 `xml:"web:chunkText"`
}

type Text1 struct {
	TypeRow []byte `xml:"inputText"`
}

//////////////////////////////////////////////////
