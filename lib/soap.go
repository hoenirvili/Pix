package lib

import "encoding/xml"

type Envelope struct {
	CreateBody Body `xml:"Body"`
}

type Body struct {
	CreateText Text `xml:"parseText_XML"`
}

type Text struct {
	TextRow string `xml:"rawTextInput"`
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
	e.CreateBody.CreateText.TextRow = string(cnt)
}
