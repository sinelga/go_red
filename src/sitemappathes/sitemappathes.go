package sitemappathes

import (
//	"fmt"
	"domains"
	"log/syslog"
	"net/url"
//	"math/rand"
//	"time"
)

func CreatePathes(golog syslog.Writer,keyword_phrasearr []domains.Keyword_phrase) []string{
	
	var pathesarr []string 
	
	for _,keyword_phrase :=range keyword_phrasearr {
		
//		golog.Info(keyword_phrase.Phrase)
		
		urlv := url.URL{}
		
//		urlv.Scheme="http"
//		urlv.Host= site
		urlv.Path =keyword_phrase.Keyword+"/"+keyword_phrase.Phrase

		
		pathesarr = append(pathesarr,urlv.String())
		
		
		
	}
	
	return pathesarr 


}
