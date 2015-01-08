package startones

import (
//	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log/syslog"
//	"os"
	"strings"
//	"fmt"
)

//func Start(golog syslog.Writer) ([]string,map[string]struct{}) {
  func Start(golog syslog.Writer) ([]string) {
	
//	sitestoblock := make(map[string]struct{})
	

	content, err := ioutil.ReadFile("/home/juno/git/go_red/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	cleanparameters := []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

//	c, err := redis.Dial("tcp", ":6379")
//	defer c.Close()
//	if err != nil {
//
//		golog.Err("startones: " + err.Error())
//		os.Exit(1)
//	} else {
//
//		if limitsites, err := redis.Strings(c.Do("ZRANGEBYSCORE", "limitsites", "("+cleanparameters[2], "+inf")); err != nil {
//
//			golog.Crit("FindFromQ: " + err.Error())
//
//		} else {
//
//			
//			for _,sitetoblock :=range limitsites {
//				
//				sitestoblock[sitetoblock] = struct{}{} 
//				
//			}
//			
//
//
//		}
//
//	}
//	c.Close()

	return cleanparameters

}
