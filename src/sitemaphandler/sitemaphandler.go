package sitemaphandler

import (
	"domains"
	"encoding/json"
	"encoding/xml"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"time"
	"bytes"
	"keywords_and_phrases"
	"strconv"
	"sitemappathes"
	"somekeywords"
	"somephrases"
)

func Create(golog syslog.Writer, c redis.Conn, locale string, themes string, site string, startparameters []string, quant string) []byte {

	var buffer bytes.Buffer

	var bkeyword_phrasearr []byte

	quantint, _ := strconv.Atoi(quant)

	queuename := locale + ":" + themes + ":sitemap:" + site
	var keyword_phrasearr []domains.Keyword_phrase

	if exist, err := redis.Int(c.Do("EXISTS", queuename)); err != nil {

		golog.Err("sitemap:exist " + err.Error())

	} else {

		if exist == 1 {

			if sitemap, err := redis.String(c.Do("GET", queuename)); err != nil {

				golog.Err("sitemap" + err.Error())

			} else {

				bsitemap := []byte(sitemap)

				if err = json.Unmarshal(bsitemap, &keyword_phrasearr); err != nil {

					golog.Err(err.Error())

				}

			}

		} else if exist == 0 {

			keywords, phrases := keywords_and_phrases.GetAll(golog, locale, themes, startparameters)

			somekeywordsres := somekeywords.GetSome(golog, keywords, quantint)
			somephrasesres := somephrases.GetSome(golog, phrases, quantint)

			if len(somekeywordsres) <= len(somephrasesres) {

				for i, keyword := range somekeywordsres {

					keyword_phrase := domains.Keyword_phrase{keyword, somephrasesres[i]}
					keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

				}

			} else {

				for i, phrase := range somephrasesres {

					keyword_phrase := domains.Keyword_phrase{somekeywordsres[i], phrase}
					keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

				}

			}

			if bkeyword_phrasearr, err = json.Marshal(keyword_phrasearr); err != nil {

				golog.Err(err.Error())

			} else {

				if _, err := redis.String(c.Do("SET", queuename, string(bkeyword_phrasearr), "EX", 864000)); err != nil {

					golog.Err(err.Error())

				}

			}

		}

	}

	pathsarr := sitemappathes.CreatePathes(golog, keyword_phrasearr)

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	
	d := time.Now()

	for i, path := range pathsarr {

		pubdate := time.Date(d.Year(), d.Month(), d.Day()-(i+1), 0, 0, 0, 0, time.UTC).Local().Format(time.RFC3339)

		doc := new(domains.Page)

		doc.Loc = "http://" + site + "/!#/q/" + path
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
