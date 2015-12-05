package lib

import "bytes"

var entities = map[string]string{
	" ":  "&nbsp;",
	"&":  "&amp;",
	"<":  "&lt;",
	">":  "&gt;",
	"\"": "&quot;",
}

func replaceAllEntities(buff []byte) []byte {

	for key, val := range entities {

		buff = bytes.Replace(buff, []byte(val), []byte(key), -1)

	}

	return buff
}
