package keywords_and_phrases

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func GetAll(golog syslog.Writer, locale string, themes string, startparameters []string) ([]string, []string) {

	redisprotocol := startparameters[0]
	redishost := startparameters[1]

	queuename := locale + ":" + themes + ":keywords"

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}

	var keywords []string

	if keywords, err = redis.Strings(c.Do("ZREVRANGEBYSCORE", queuename, "+inf", "-inf", "LIMIT", 0, 1000)); err != nil {

		golog.Crit("findkeywords: " + err.Error())

	}

	queuename = locale + ":" + themes + ":phrases"

	var phrases []string

	if phrases, err = redis.Strings(c.Do("ZREVRANGEBYSCORE", queuename, "+inf", "-inf", "LIMIT", 0, 1000)); err != nil {

		golog.Crit("keywords_and_phrases: " + err.Error())

	}

	return keywords, phrases

}
