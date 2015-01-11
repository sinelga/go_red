package sitemapcreator

import (
	"bytes"
	"domains"
	"encoding/xml"
	"log/syslog"
	"sitemappathes"
	"time"
)

func Createsitemap(golog syslog.Writer, keyword_phrasearr []domains.Keyword_phrase, site string) []byte {

	var buffer bytes.Buffer

	pathsarr := sitemappathes.CreatePathes(golog, keyword_phrasearr)

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	d := time.Now()

	for i, path := range pathsarr {

		pubdate := time.Date(d.Year(), d.Month(), d.Day()-(i+1), 0, 0, 0, 0, time.UTC).Local().Format(time.RFC3339)

		doc := new(domains.Page)

//		doc.Loc = "http://" + site + "/#!/q/" + path
		doc.Loc = "http://" + site + "/q/" + path
		doc.Lastmod = pubdate
		//	doc.Name = "The Example Times"
		//	doc.Language = "en"
		//	doc.Title = "Companies A, B in Merger Talks"
		//	doc.Keywords = "business, merger, acquisition, A, B"
		//	doc.Image = "http://www.google.com/spacer.gif"
		docList.Pages = append(docList.Pages, doc)

	}

	resultXml, err := xml.MarshalIndent(docList, "", "  ")
	if err != nil {

		golog.Crit(err.Error())
	}

	buffer.WriteString(xml.Header)
	buffer.Write(resultXml)

	return buffer.Bytes()

}
