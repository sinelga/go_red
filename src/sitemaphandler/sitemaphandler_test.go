package sitemaphandler

import (
    "testing"
    "log/syslog"
    "log"
    "github.com/garyburd/redigo/redis"
)

func TestCreate(t *testing.T) {
	
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	c_local, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c_local.Close()
	
	locale :="fi_FI"
	themes :="porno"
	
	startparameters := []string{"tcp",":6379"}

	Create(*golog,c_local,locale,themes,"www.test.com",startparameters,"10")

}

