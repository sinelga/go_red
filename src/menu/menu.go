package menu

import (
	"github.com/garyburd/redigo/redis"
	"keywords_and_phrases"
	"log/syslog"
	"somekeywords"
	"somephrases"
	//	"strings"
	"domains"
	"encoding/json"
)

func GetMenu(golog syslog.Writer, c redis.Conn, startparameters []string, locale string, themes string, site string) []byte {

	var bkeyword_phrasearr []byte

	queuename := locale + ":" + themes + ":menu:" + site

	if exist, err := redis.Int(c.Do("EXISTS", queuename)); err != nil {

		golog.Err("menu:exist " + err.Error())

	} else {

		if exist == 1 {

			if menu, err := redis.String(c.Do("GET", queuename)); err != nil {

				golog.Err("menu " + err.Error())
			} else {

				bkeyword_phrasearr = []byte(menu)

			}

		} else if exist == 0 {
			

			keywords, phrases := keywords_and_phrases.GetAll(golog, locale, themes, startparameters)

			quantint := 50

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

				if _, err := redis.String(c.Do("SET", queuename, string(bkeyword_phrasearr),"EX",60)); err != nil {

				}

			}			
			
		}

	}


	return bkeyword_phrasearr

}
