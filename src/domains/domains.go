package domains

import (
	"encoding/xml"
//	"time"
	)

type ServerConfig struct {
	Main struct {
		Locale     string
		Themes     string
		Changefreq string
	}
	Dirs struct {
		Rootdir        string
		Rootdirm	string
		Backendrootdir string		
		Dbdir          string
		Rssresorsesfile string
	}
}

type Rssresors struct {
	Topic string
	Link  string
}

type Fortune_feed_links struct {
	
	Locale string
	Themes string
	Path string
	Qdomain string
	Qpath string
	
	
}


type FortuneZodiac struct {
	Redlink string
	Zodiacinfo []byte		
}


type Fortuneteller struct {
	Id int
	Name string
	Phone string
	Location string
	Moto string
	Desc string
	
	
}

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

type Pages struct {
//	Version string   `xml:"version,attr"`			
	XMLName    xml.Name `xml:"urlset"`
	XmlNS      string   `xml:"xmlns,attr"`
//	XmlImageNS string   `xml:"xmlns:image,attr"`
//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages      []*Page  `xml:"url"`
}

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

type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string `xml:"lastmod"`
//	Name       string   `xml:"news:news>news:publication>news:name"`
//	Language   string   `xml:"news:news>news:publication>news:language"`
//	Title      string   `xml:"news:news>news:title"`
//	Keywords   string   `xml:"news:news>news:keywords"`
//	Image      string   `xml:"image:image>image:loc"`
}
