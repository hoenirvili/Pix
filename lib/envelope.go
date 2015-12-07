package lib

import "encoding/xml"

func newEnvelope(key int, path string, buffer []byte) (interface{}, error) {
	switch key {
	case 0:
		var env Envelope
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TextRow = buffer
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://webPosRo.uaic/"
		return env, err
	case 1:
		var env Envelope1
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://webNpChunkerRo.uaic/"
		return env, err
	case 2:
		var env Envelope2
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://webFdgRo.uaic/"
		return env, err
	case 3:
		var env Envelope3
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		//DEFAULT RO
		env.CreateBody.CreateText.Lang = []byte("ro")
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://recognizer.entity.named.uaic/"
		return env, err
	case 4:
		var env Envelope4
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		//DEFAULT RO
		env.CreateBody.CreateText.Lang = []byte("ro")
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://resolution.anaphora.uaic/"
		return env, err
	case 5:
		var env Envelope5
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		//DEFAULT RO
		env.CreateBody.CreateText.Lang = []byte("ro")
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://splitter.clause.uaic/"
		return env, err
	case 6:
		var env Envelope6
		buffSliceEnv := fileContent(path)
		err := xml.Unmarshal(buffSliceEnv, &env)
		env.CreateBody.CreateText.TypeRow = buffer
		//DEFAULT RO
		env.CreateBody.CreateText.Lang = []byte("ro")
		env.Val1 = "http://schemas.xmlsoap.org/soap/envelope/"
		env.Val2 = "http://parser.discourse.uaic/"
		return env, err
	}

	return nil, nil
}

// get xmlEnvelope
func getEnvelope(env interface{}) ([]byte, error) {
	buff, err := xml.MarshalIndent(env, "", "   ")

	if err != nil {
		return nil, err
	}

	return buff, nil
}
