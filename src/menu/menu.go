package menu

import (
	"github.com/garyburd/redigo/redis"
	"keywords_and_phrases"
	"log/syslog"
	"somekeywords"
	"somephrases"
	"domains"
	"encoding/json"
	"sitemaphandler"
	"strconv"
)

func GetMenu(golog syslog.Writer, c redis.Conn, startparameters []string, locale string, themes string, site string, quant string) []byte {

	var bkeyword_phrasearr []byte

	quantint, _ := strconv.Atoi(quant)

	menu_queuename := locale + ":" + themes + ":menu:" + site

	if exist, err := redis.Int(c.Do("EXISTS", menu_queuename)); err != nil {

		golog.Err("menu:exist " + err.Error())

	} else {

		if exist == 1 {

			if menu, err := redis.String(c.Do("GET", menu_queuename)); err != nil {

				golog.Err("menu " + err.Error())
			} else {

				bkeyword_phrasearr = []byte(menu)

			}

		} else if exist == 0 {

			keywords, phrases := keywords_and_phrases.GetAll(golog, locale, themes, startparameters)

			somekeywordsres := somekeywords.GetSome(golog, keywords, quantint)
			somephrasesres := somephrases.GetSome(golog, phrases, quantint)

			var keyword_phrasearr []domains.Keyword_phrase

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

				if _, err := redis.String(c.Do("SET", menu_queuename, string(bkeyword_phrasearr), "EX", 172800)); err != nil {

					golog.Err(err.Error())

				}

			}

			sitemaphandler.Create(golog, c, locale, themes, site, startparameters, quant, keyword_phrasearr)

		}

	}

	return bkeyword_phrasearr

}
