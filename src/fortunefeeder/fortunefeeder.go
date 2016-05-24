package main

import (
	"encoding/csv"
	//	"encoding/json"
	"domains"
	"flag"
	"fmt"
//	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"os"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var siteFlag *string = flag.String("site", "", "Site like www.test.com")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

//	site := *siteFlag

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	csvFile, err := os.Open("/home/juno/git/go_red/feedlinks.csv")

	if err != nil {
		golog.Crit(err.Error())
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		golog.Crit(err.Error())
	}

	var links []domains.Fortune_feed_links
	var link domains.Fortune_feed_links

	for _, each := range csvData {

		link.Locale = each[0]
		link.Themes = each[1]
		link.Path = each[2]
		link.Qdomain = each[3]
		link.Qpath = each[4]

		links = append(links, link)

	}

//	var allFortuneZodiac []domains.FortuneZodiac
//	var fortuneZodiac domains.FortuneZodiac

	for _, link := range links {

		response, err := http.Get(link.Qdomain + "/" + link.Qpath)
		if err != nil {

			golog.Crit(err.Error())
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {

				golog.Crit(err.Error())
				os.Exit(1)
			}
			fmt.Println(string(contents))

//			fortuneZodiac.Redlink = link.Locale + ":" + link.Themes + ":" + site + ":" + link.Path
//			fortuneZodiac.Zodiacinfo = contents
//
//			allFortuneZodiac = append(allFortuneZodiac, fortuneZodiac)

		}

	}

//	c, err := redis.Dial("tcp", ":6379")
//	if err != nil {
//
//		golog.Crit(err.Error())
//
//	}
//
//	for _, fortunezodiac := range allFortuneZodiac {
//
//		if _, err := c.Do("SET", fortunezodiac.Redlink, fortunezodiac.Zodiacinfo); err != nil {
//
//			golog.Err("fortunefeeder " + err.Error())
//		}
//
//	}

}
