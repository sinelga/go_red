package main

import (
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"startones"
	"sync"
	"bthandler"
)

var startOnce sync.Once
var startparameters []string

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	//	themes := req.Header.Get("X-THEMES")
	//	locale := req.Header.Get("X-LOCALE")
	//	variant := req.Header.Get("X-VARIANT")
		site := req.Header.Get("X-DOMAIN")
		pathinfo := req.Header.Get("X-PATHINFO")
	//	bot := req.Header.Get("X-BOT")

	startOnce.Do(func() {
		startparameters = startones.Start(*golog)
	})
	
	
	bthandler.BTrequestHandler(*golog, resp, req, "fi_FI", "porno", site,pathinfo , "google", startparameters,false,"variant")

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
