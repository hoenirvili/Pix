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
		//TODO
		break
	case 3:
		//TODO
		break
	case 4:
		//TODO
		break
	case 5:
		//TODO
		break
	case 6:
		//TODO
		break
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
