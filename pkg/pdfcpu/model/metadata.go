package model

import (
	"encoding/xml"
	"time"
)

type UserDate time.Time

const userDateFormatNoTimeZone = "2006-01-02T15:04:05Z"
const userDateFormatNegTimeZone = "2006-01-02T15:04:05-07:00"
const userDateFormatPosTimeZone = "2006-01-02T15:04:05+07:00"

func (ud *UserDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

type Alt struct {
	Entries []string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# li"`
}

type Seq struct {
	Entries []string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# li"`
}

type Title struct {
	Alt Alt `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Alt"`
}

type Desc struct {
	Alt Alt `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Alt"`
}

type Creator struct {
	Seq Seq `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Seq"`
}

type Description struct {
	Title        Title    `xml:"http://purl.org/dc/elements/1.1/ title"`
	Author       Creator  `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Subject      Desc     `xml:"http://purl.org/dc/elements/1.1/ description"`
	Creator      string   `xml:"http://ns.adobe.com/xap/1.0/ CreatorTool"`
	CreationDate UserDate `xml:"http://ns.adobe.com/xap/1.0/ CreateDate"`
	ModDate      UserDate `xml:"http://ns.adobe.com/xap/1.0/ ModifyDate"`
	Producer     string   `xml:"http://ns.adobe.com/pdf/1.3/ Producer"`
	Trapped      bool     `xml:"http://ns.adobe.com/pdf/1.3/ Trapped"`
	Keywords     string   `xml:"http://ns.adobe.com/pdf/1.3/ Keywords"`
}

type RDF struct {
	XMLName     xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	Description Description
}

type XMPMeta struct {
	XMLName xml.Name `xml:"adobe:ns:meta/ xmpmeta"`
	RDF     RDF
}

func removeTag(s, kw string) string { _ = "STUB: not implemented"; return "" }

func RemoveKeywords(metadata *[]byte) error { _ = "STUB: not implemented"; return nil }
