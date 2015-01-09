package findfreeparagraph

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"strconv"
)


func FindFromQ(golog syslog.Writer, locale string, themes string,site string, bot string,startparameters []string) domains.Paragraph {

	redisprotocol := startparameters[0]
	redishost := startparameters[1]
	blockhitlimitstr := startparameters[2]

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}

	queuename := locale + ":" + themes

	var unmarPar domains.Paragraph

	if quan_prs, err := redis.Int(c.Do("LLEN", queuename)); err != nil {

		golog.Crit(err.Error())

	} else {

		if quan_prs > 1 {

			bparagraph, _ := redis.Bytes(c.Do("LPOP", queuename))

			err := json.Unmarshal(bparagraph, &unmarPar)
			if err != nil {

				golog.Crit(err.Error())

			}
			
//			looks like g don't like it

			if pushsite, err := redis.Strings(c.Do("ZRANGEBYSCORE", "pushdomains", "-inf", "+inf", "LIMIT", 0, 1)); err != nil {

				golog.Crit("FindFromQ: " + err.Error())

			} else {

				if len(pushsite) > 0  {

					unmarPar.Pushsite = pushsite[0]

					_, err = c.Do("ZINCRBY", "pushdomains", 1, pushsite[0])
					if err != nil {

						golog.Crit(err.Error())

					}

				}
				
//				if bot=="google" {
										
					if hits,err := redis.Int(c.Do("ZINCRBY",bot , "1", site)); err != nil {
						
						golog.Crit("FindFromQ: " + err.Error())
					} else {
												
						
						blockhitlimit_int,_ :=strconv.Atoi(blockhitlimitstr)	
						
						if hits > blockhitlimit_int {
							
							golog.Info(bot +" hit??? > "+site+" "+strconv.Itoa(hits)+" limit "+blockhitlimitstr)
							
							unmarPar.Plocallink=""
							
						}
												
						
					}				
					
//				}
																				
			}

		} else {

			golog.Crit("need more free paragraphs!!!!")

		}

	}
	return unmarPar
}
