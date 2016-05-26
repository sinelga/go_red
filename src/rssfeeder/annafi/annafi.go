package main

import (
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"net/http"
	"golang.org/x/net/html/atom"	
)

func main() {

	resp, err := http.Get("http://anna.fi/horoskoopit/paivahoroskoopit/paivahoroskooppi-torstai-26-5/")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)

	}
				
	articles := scrape.FindAll(root, scrape.ByClass("article__body"))

	for _, article := range articles {

		prs :=scrape.FindAll(article,scrape.ByTag(atom.H2))
		
		for _,pr :=range prs {
			
			fmt.Println(scrape.Text(pr))
			
			next,ok :=	scrape.Find(pr.NextSibling.NextSibling.NextSibling.NextSibling,	scrape.ByTag(atom.P))
			
			if ok {
				fmt.Println(scrape.Text(next))
			
			}													
			
		}

		
	}

}
