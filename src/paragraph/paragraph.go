package paragraph

import (
	"github.com/garyburd/redigo/redis"

	"encoding/json"
	"findfreeparagraph"
	"log/syslog"
)

func GetParagrph(golog syslog.Writer, c redis.Conn, startparameters []string, locale string, themes string, site string, path string,menupath string) []byte {

	var bparagraph []byte

	queuename := locale + ":" + themes + ":paragraph:" + site + ":" + menupath
	if exist, err := redis.Int(c.Do("EXISTS", queuename)); err != nil {

		golog.Err("menu:exist " + err.Error())

	} else {

		if exist == 1 {

			if paragraph, err := redis.String(c.Do("GET", queuename)); err != nil {

				golog.Err("menu " + err.Error())
			} else {

				bparagraph = []byte(paragraph)

			}

		} else if exist == 0 {

			paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, "test.com", "google", startparameters)


			if paragraph, err := json.Marshal(paragraph); err != nil {

				golog.Err(err.Error())

			} else {
				
				bparagraph = []byte(paragraph)
				
				
				if _, err := redis.String(c.Do("SET", queuename, string(bparagraph),"EX",120)); err != nil {
					
					golog.Err(err.Error())

				}
								

			}


		}

	}

	return bparagraph
}
