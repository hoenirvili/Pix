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

var s = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<ns2:parseText_XMLResponse xmlns:ns2="http://webPosRo.uaic/">
			<return>
				<![CDATA[
						<POS_Output>
							 <S id="1" offset="0">
								<W LEMMA="”" MSD="DBLQ" id="1.1" offset="0">„</W>
								<W LEMMA="într-adevăr" MSD="Rg" POS="ADVERB" id="1.2" offset="1">Într-adevăr</W>
								<W LEMMA="," MSD="COMMA" id="1.3" offset="12">,</W>
								<W LEMMA="acum" MSD="Rg" POS="ADVERB" id="1.4" offset="14">acum</W>
							</S>
						</POS_Output>
				]]>
			</return>
		</ns2:parseText_XMLResponse>
	</S:Body>
</S:Envelope>`

func parseXMLResponse(responseBuffer string) {
	v := XMLResponse{}

	var responseByteBuffer = []byte(responseBuffer)

	//SERIALIZING
	responseByteBuffer = bytes.Replace(responseByteBuffer, []byte("<![CDATA["), []byte(""), -1)
	responseByteBuffer = bytes.Replace(responseByteBuffer, []byte("]]>"), []byte(""), -1)

	fmt.Println()
	fmt.Printf("%s\n", responseByteBuffer)
	fmt.Println()

	err := xml.Unmarshal([]byte(responseBuffer), &v)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("%+v\n", v.XMLName)
	fmt.Printf("%+v\n", v.RespBody)
	fmt.Printf("%+v\n", v.RespBody.StrgRet)
	fmt.Printf("%+v\n", v.RespBody.StrgRet.Ret)
	fmt.Printf(" DATA = %+v\n", v.RespBody.StrgRet.Ret.POS_Output)
	fmt.Println()

}

func TestParseXMLResponse(t *testing.T) {
	parseXMLResponse(s)
}
