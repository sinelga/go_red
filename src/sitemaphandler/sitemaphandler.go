package sitemaphandler

import (
	"domains"
	"encoding/json"
	//	"encoding/xml"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	//	"time"
	//	"bytes"
	"keywords_and_phrases"
	"strconv"
	//	"sitemappathes"
	"somekeywords"
	"somephrases"
)

func Create(golog syslog.Writer, c redis.Conn, locale string, themes string, site string, startparameters []string, quant string, keyword_phrase_from_menu []domains.Keyword_phrase) []domains.Keyword_phrase {


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

			if keyword_phrase_from_menu != nil {

				for _, keyword_phrase := range keyword_phrase_from_menu {

					keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

				}

				putIntoRedis(golog, c, keyword_phrasearr, queuename)

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

			if keyword_phrase_from_menu != nil {

				for _, keyword_phrase := range keyword_phrase_from_menu {

					keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

				}

				putIntoRedis(golog, c, keyword_phrasearr, queuename)

			}

			putIntoRedis(golog, c, keyword_phrasearr, queuename)

		}

	}

	return keyword_phrasearr

}

func putIntoRedis(golog syslog.Writer, c redis.Conn, keyword_phrasearr []domains.Keyword_phrase, queuename string) {

	if bkeyword_phrasearr, err := json.Marshal(keyword_phrasearr); err != nil {

		golog.Err(err.Error())

	} else {

		if _, err := redis.String(c.Do("SET", queuename, string(bkeyword_phrasearr), "EX", 864000)); err != nil {

			golog.Err(err.Error())

		}

	}

}
