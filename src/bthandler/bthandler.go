package bthandler

import (
	//	"clean_pathinfo"
	//	"createfirstgz"
	//	"createpage"
	//	"domains"
	//	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	//	"keywords_and_phrases"
	"log/syslog"
	"net/http"
	//	"somekeywords"
	//	"somephrases"
	"menu"
	"strings"

	//	"redis_hdl"
	//	"sitemaphandler"
	//	"robots_txt"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, bot string, startparameters []string, blocksite bool, variant string) {


	c_local, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}

	if strings.HasSuffix(pathinfo, "menu") {

		bkeyword_phrasearr := menu.GetMenu(golog, c_local, startparameters, locale, themes, site)

		golog.Info("bkeyword_phrasearr "+string(bkeyword_phrasearr))

		resp.Header().Add("Content-type", "application/javascript")

		resp.Header().Add("Access-Control-Allow-Origin", "*")
		jsonBytes := []byte(fmt.Sprintf("%s", bkeyword_phrasearr))

		resp.Write(jsonBytes)

	}

}
