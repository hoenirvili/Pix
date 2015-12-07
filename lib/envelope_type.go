// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

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

// FDG PARSER ENVELOPE SERVICE
type Envelope2 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:web,attr"`
	CreateBody Body2    `xml:"soapenv:Body"`
}

type Body2 struct {
	CreateText Text2 `xml:"web:parseText"`
}

type Text2 struct {
	TypeRow []byte `xml:"txt"`
}

//////////////////////////////////////////////////

// NAME ENTITY RECOGNIZER SERVICE
type Envelope3 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:rec,attr"`
	CreateBody Body3    `xml:"soapenv:Body"`
}

type Body3 struct {
	CreateText Text3 `xml:"rec:recognizeEntities"`
}

type Text3 struct {
	TypeRow []byte `xml:"text"`
	Lang    []byte `xml:"language"`
}

//////////////////////////////////////////////////

//ANAPHORA RESOLUTION SERVICE

type Envelope4 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:res,attr"`
	CreateBody Body4    `xml:"soapenv:Body"`
}

type Body4 struct {
	CreateText Text4 `xml:"res:solveLinks"`
}

type Text4 struct {
	TypeRow []byte `xml:"text"`
	Lang    []byte `xml:"language"`
}

///////////////////////////////////////////////////

// CLAUSE SPLITTER SERVICE
type Envelope5 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:spl,attr"`
	CreateBody Body5    `xml:"soapenv:Body"`
}

type Body5 struct {
	CreateText Text5 `xml:"spl:split"`
}

type Text5 struct {
	TypeRow []byte `xml:"text"`
	Lang    []byte `xml:"language"`
}

///////////////////////////////////////////////////////

//DISCOURSE PARSE SERVICE
type Envelope6 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:par,attr"`
	CreateBody Body6    `xml:"soapenv:Body"`
}

type Body6 struct {
	CreateText Text6 `xml:"par:parse"`
}

type Text6 struct {
	TypeRow []byte `xml:"text"`
	Lang    []byte `xml:"language"`
}

///////////////////////////////////////////////////////////

// BERMUDA
type Envelope7 struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:ws,attr"`
	CreateBody Body7    `xml:"soapenv:Body"`
}

type Body7 struct {
	CreateText Text7 `xml:"ws:ProcessText"`
}

type Text7 struct {
	Model []byte `xml:"model"`
	Test  []byte `xml:"text"`
}
