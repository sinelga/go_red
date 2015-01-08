package redis_hdl

import (
	//		"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func GetRedis(golog syslog.Writer,c redis.Conn, command string,queuename string,params []string) {



	
//	c, err := redis.Dial(redisprotocol, redishost)
//	if err != nil {
//
//		golog.Crit(err.Error())
//
//	}

}
