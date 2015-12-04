package lib

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Return struct {
	MainBuffer string
}
type StorageReturn struct {
	Ret Return `xml:"return"`
}

type Body struct {
	StrgRet StorageReturn `xml:"parseText_XMLResponse"`
}

// check
type XMLResponse struct {
	XMLName  xml.Name `xml:"Envelope"`
	RespBody Body     `xml:"Body"`
}

func parseXMLResponse() {
	responseBuffer := ``
	v := XMLResponse{}

	err := xml.Unmarshal([]byte(responseBuffer), &v)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", v)
}

//
//func loadXML(string xmlPath) string {

//}

// send soap request and return an xml identity
func sendSoapRequest(bodyRequest, wsldUrl string) string {
	fmt.Fprintf(os.Stdout, "%s : %s\n", "Starting the request on", wsldUrl)
	//	var xmlStr =
	//client := http.Client{}
	return ""
}
