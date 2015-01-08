package bthandler

import (
    "testing"
    "log"
	"log/syslog"
//	"github.com/garyburd/redigo/redis"
	"net/http"
	"net/http/httptest"
)

func TestBTrequestHandler(t *testing.T) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
				
	}
	
	startparameters :=[]string{"tcp",":6379","5000"}
	
	
	req, err := http.NewRequest("GET", "http://www.example.com/menu", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
		
	BTrequestHandler(*golog, resp, req , "fi_FI", "porno", "www.test.com", "/menu", "google", startparameters, false,"variant")
		

}

