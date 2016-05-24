package main

import (
	//	"blogfeeder/addlink"
	"domains"
	"encoding/csv"
	//	"encoding/json"
	//	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	//	"github.com/gosimple/slug"
	"gopkg.in/gcfg.v1"
	"log"
	//	"path/filepath"
	//	"dbhandler"
	"gopkg.in/mgo.v2"
	"os"
	//	"time"
)

var rootdir = ""
var backendrootdir = ""
var locale = ""
var themes = ""
var rssresorsesfile = ""

var resorses []domains.Rssresors

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes
		backendrootdir = cfg.Dirs.Backendrootdir
		rssresorsesfile = cfg.Dirs.Rssresorsesfile

	}

	csvfile, err := os.Open(rssresorsesfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {

		fmt.Println(err)
		return
	}

	for _, record := range records {

		res := domains.Rssresors{record[0], record[1]}
		resorses = append(resorses, res)
	}

}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//	linksdir := filepath.Join(rootdir, "links")

	//	uniqstitle := dbhandler.GetAllStitle(*session, locale, themes)

	for _, res := range resorses {

		//		now := time.Now()

		topic := res.Topic
		//		stopic := slug.Make(topic)
		fmt.Println(topic)

		feed, err := rss.Fetch(res.Link)
		if err != nil {
			// handle error.
			panic(err.Error())
		}

		items := feed.Items

		for _, item := range items {

			title := item.Title
			//			stitle := slug.Make(title)

			//			if _, ok := uniqstitle[stitle]; !ok {

			contents := item.Summary
			
			fmt.Println(title,contents)

			//				site := addlink.AddLinktoAllfiles(linksdir, stopic, stitle)

			//				fmt.Println("site", site)
			//				item := domains.BlogItem{stopic, topic, stitle, title, contents, now, now}

			//				dbhandler.InsertRecord(*session, locale, themes, site, "blog", stopic, topic, item)

			//			}

		}
	}

}
