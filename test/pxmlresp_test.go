// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package Pix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

type Return struct {
	POS_Output string `xml:"POS_Output"`
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

var s = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<ns2:parseText_XMLResponse xmlns:ns2="http://webPosRo.uaic/">
			<return>
				<![CDATA[
					<POS_Output>testuleeee</POS_Output>
				]]>
			</return>
		</ns2:parseText_XMLResponse>
	</S:Body>
</S:Envelope>`)

func parseXMLResponse(responseBuffer []byte) {
	v := XMLResponse{}

	//var responseByteBuffer = []byte(responseBuffer)

	//SERIALIZING
	responseBuffer = bytes.Replace(responseBuffer, []byte("<![CDATA["), []byte(""), -1)
	responseBuffer = bytes.Replace(responseBuffer, []byte("]]>"), []byte(""), -1)

	err := xml.Unmarshal(responseBuffer, &v)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("%v\n", v.XMLName)
	fmt.Printf("%v\n", v.RespBody)
	fmt.Println()

}

func TestParseXMLResponse(t *testing.T) {
	parseXMLResponse(s)
}
