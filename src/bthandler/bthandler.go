package bthandler

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"menu"
	"net/http"
	"paragraph"
	"sitemaphandler"
	"strings"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, bot string, startparameters []string, blocksite bool, variant string,menupath string,quant string) {

	c_local, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	var jsonBytes []byte
	var bres []byte

	if strings.HasPrefix(pathinfo, "/menu") {

		bres = menu.GetMenu(golog, c_local, startparameters, locale, themes, site,quant)

	} else if strings.HasPrefix(pathinfo, "/paragraph") {

		bres = paragraph.GetParagrph(golog, c_local, startparameters, locale, themes, site,pathinfo, menupath)

	} else if strings.HasPrefix(pathinfo, "/sitemap.xml") {

		bres = sitemaphandler.Create(golog,c_local,locale,themes,site,startparameters,"15")

	}
	
	
	//	golog.Info("bres " + string(bres))
	if strings.HasPrefix(pathinfo, "/sitemap.xml") {
		resp.Header().Add("Content-type", "application/xml")
		
	} else {
	
	resp.Header().Add("Content-type", "application/javascript")
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	
	}
	jsonBytes = []byte(fmt.Sprintf("%s", bres))

	resp.Write(jsonBytes)

}