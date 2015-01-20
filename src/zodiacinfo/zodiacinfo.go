package zodiacinfo

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"strings"
)

func GetInfo(golog syslog.Writer, c redis.Conn, locale string, themes string, site string, path string) []byte {

	var tryget bool = false

	var bzodiacinfo []byte

	var redlink string = locale + ":" + themes + ":" + site + ":" + path

	if exist, err := redis.Int(c.Do("EXISTS", redlink)); err != nil {

		golog.Err("zodiacinfo:exist " + err.Error())

	} else {

		if exist == 0 {

			golog.Info("site " + site)

			splitsite := strings.Split(site, ".")

			if len(splitsite) > 1 {

				splitsite := strings.Split(site, ".")
				onlydomain := splitsite[len(splitsite)-2] + "." + splitsite[len(splitsite)-1]
				redlink = locale + ":" + themes + ":" + onlydomain + ":" + path

			}

			if exist, err := redis.Int(c.Do("EXISTS", redlink)); err != nil {

				golog.Err("zodiacinfo:exist " + err.Error())

			} else {

				if exist == 1 {
					
					tryget = true

				}

			}

		} else if exist == 1 {

			tryget = true

		}

		if tryget {

			if resualt, err := redis.Bytes(c.Do("GET", redlink)); err != nil {

				golog.Err("zodiacinfo " + err.Error())
			} else {

				bzodiacinfo = resualt

			}
		}

	}

	return bzodiacinfo

}
