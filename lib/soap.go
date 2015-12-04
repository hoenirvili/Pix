package lib

import (
	"encoding/xml"
	"fmt"
	"os"
)

var soapRequestBuffer = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://webPosRo.uaic/">
	<soapenv:Body>
		<web:parseText_XML>
				<rawTextInput>
				</rawTextInput>
		</web:parseText_XML>
		</soapenv:Body>
 </soapenv:Envelope>`

type Return struct {
	TextInput string `xml:"rowTextInput"`
}
type Action struct {
	Ret Return `xml:"return"`
}

type Body struct {
	ActionSet Action `xml:"parseText_XML"`
}

// check
type XMLResponse struct {
	XMLName  xml.Name `xml:"Envelope"`
	RespBody Body     `xml:"Body"`
}

func unloadXml() *XMLResponse {
	v := XMLResponse{}

	err := xml.Unmarshal([]byte(soapRequestBuffer), &v)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", v)
	return &v
}

func loadXml(buffer []byte) string {
	v := unloadXml()

	//load content of file into buffer
	v.RespBody.ActionSet.Ret.TextInput = string(buffer)

	//fmt.Printf("%+v\n", v.RespBody.ActionSet.Ret.TextInput)

	buffSlice, err := xml.Marshal(v)

	if err != nil {
		fmt.Println(err)
	}

	return rune.(buffSlice)
}

// send soap request and return an xml identity
func sendSoapRequest(bodyRequest []byte, wsldUrl string) string {

	fmt.Fprintf(os.Stdout, "%s : %s\n", "Starting the request on", wsldUrl)

	buffer := loadXml(bodyRequest)

	fmt.Println(buffer)
	//curlCommand := fmt.Sprintf("curl -H \"Content-Type; text/xml;charset=UTF-8\" -H \"parseText_XML\" --data-row %s -X POST %s", buffer, wsldUrl)

	return ""
}
