package lib

import "bytes"

func replaceAllEntities(buff []byte) []byte {

	var entities = map[string]string{
		" ":  "&nbsp;",
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
	}

	for key, val := range entities {
		buff = bytes.Replace(buff, []byte(val), []byte(key), -1)
	}

	return buff
}

func sanitizeEnvelope(buffer []byte) []byte {

	var (
		StartF = []byte("<Envelope")
		FinalF = []byte("</Envelope>")
		StartT = []byte("<soapenv:Envelope ")
		FinalT = []byte("</soapenv:Envelope>")
	)

	// Check all the bytes that is equal to StartF and FinalF
	// And replace all with StartT and FinalT
	buffer = bytes.Replace(buffer, StartF, StartT, -1)
	buffer = bytes.Replace(buffer, FinalF, FinalT, -1)

	// return the new sanitize envelope buffer
	return buffer
}
