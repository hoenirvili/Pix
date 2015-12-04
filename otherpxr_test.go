package Pix

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type ReturnO struct {
	RequestId         int    `xml:"requestId"`
	DataCenterId      string `xml:"dataCenterId"`
	DataCenterVersion int    `xml:"dataCenterVersion"`
	StorageId         string `xml:"storageId"`
}

type StorageReturnO struct {
	Ret ReturnO `xml:"return"`
}

type BodyO struct {
	StrgRet StorageReturnO `xml:"createStorageReturn"`
}

type StorageResponse struct {
	XmlName  xml.Name `xml:"Envelope"`
	RespBody BodyO    `xml:"Body"`
}

var respX = `<?xml version='1.0' encoding='UTF-8'?>
	<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
		<S:Body>
			<ns2:createStorageReturn xmlns:ns2="http://ws.api.mysite.com/">
				<return>
					<requestId>16660663</requestId>
					<dataCenterId>ssrr-444tt-yy-99</dataCenterId>
					<dataCenterVersion>12</dataCenterVersion>
					<storageId>towrrt24903FR55405</storageId>
				</return>
			</ns2:createStorageReturn>
		</S:Body>
	</S:Envelope>`

func pxr(buff string) {
	v := StorageResponse{}

	err := xml.Unmarshal([]byte(buff), &v)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Printf("%+v\n", v)
	fmt.Printf("%+v\n", v.XmlName)
	fmt.Printf("%+v\n", v.RespBody)
	fmt.Printf("%+v\n", v.RespBody.StrgRet)
	fmt.Printf("%+v\n", v.RespBody.StrgRet.Ret)
	fmt.Println()
}
func TestOthertParseXMLResponse(t *testing.T) {
	pxr(respX)
}
