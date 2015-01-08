package startones

import (
	"log"
	"log/syslog"
	"testing"
	

)

func TestStart(t *testing.T) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	parameters,sitestoblock :=Start(*golog)
	
	golog.Info(parameters[0])
	golog.Info(parameters[1])
	golog.Info(parameters[2])
	
	
	for sitetoblock :=range sitestoblock {
		
		golog.Info(sitetoblock)
		
	}
	
	
	
	
}