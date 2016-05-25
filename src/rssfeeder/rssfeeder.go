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
	"net/http"
	"os"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
//	"golang.org/x/net/html/atom"
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

		for i, item := range items {

			if i == 0 {

				fmt.Println(item.Link, item.Content)

				resp, err := http.Get(item.Link)
				if err != nil {
					panic(err)
				}
				root, err := html.Parse(resp.Body)
				if err != nil {
					panic(err)

				}

//				matcher := func(n *html.Node) bool {
					// must check for nil values
//					fmt.Println(n)
//					if n.DataAtom == atom.A && n.Parent != nil  {
//						fmt.Println(scrape.ByClass(n,"body"))
////						fmt.Println(scrape.Attr(n, "class"))						
////						return scrape.Attr(n.Parent.Parent, "class") == "body"
//						return scrape.Attr(n.Parent, "class") =="title"
//					}
//					return false
//				}

//				articles := scrape.FindAll(root, matcher)
				articles := scrape.FindAll(root,scrape.ByClass("body"))
				

				for _, article := range articles {
//					fmt.Printf("%2d %s \n", i, scrape.Text(article)))
					fmt.Println(scrape.Text(article))
//					breaks := scrape.FindAll(article,scrape.ByTag(atom.Br))
//					for _,br := range breaks {
//						
//						fmt.Println(br.)
//						
//					}
					
					
				}

			}

		}
	}

}
