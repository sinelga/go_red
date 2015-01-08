package domains

import (
//	"encoding/xml"
//	"time"
	)

//type Site struct {
//	Pathinfo   string
//	Created    int64
//	Updated    int64
//	Hits       int
//	Paragraphs []Paragraph
//}

type Keyword struct {
	
	Keyword string
	Hits int	
	
}

type Phrase struct {
	
	Phrase string
	Hits int	
	
}

type Keyword_phrase struct {
	
	Keyword string
	Phrase string
	
}


type Paragraph struct {
	Ptitle     string
	Pphrase    string
	Plocallink string
	Phost      string
	Sentences  []string
	Pushsite   string
}

//type Htmlpage struct {
//	Locale string
//	Themes string
//	Variant string
//	Created string
//	Updated string
//	Paragraphs []Paragraph
//	Pushsite string
//	
//}

//type Pages struct {
////	Version string   `xml:"version,attr"`			
//	XMLName    xml.Name `xml:"urlset"`
//	XmlNS      string   `xml:"xmlns,attr"`
////	XmlImageNS string   `xml:"xmlns:image,attr"`
////	XmlNewsNS  string   `xml:"xmlns:news,attr"`
//	Pages      []*Page  `xml:"url"`
//}

//type Page struct {
//	XMLName    xml.Name `xml:"url"`
//	Loc        string   `xml:"loc"`
//	Lastmod    string `xml:"lastmod"`
////	Name       string   `xml:"news:news>news:publication>news:name"`
////	Language   string   `xml:"news:news>news:publication>news:language"`
////	Title      string   `xml:"news:news>news:title"`
////	Keywords   string   `xml:"news:news>news:keywords"`
////	Image      string   `xml:"image:image>image:loc"`
//}


