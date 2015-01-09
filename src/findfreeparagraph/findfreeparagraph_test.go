package findfreeparagraph

import (
	"log"
	"log/syslog"
	"testing"
	"fmt"
)

func TestFindFromQ(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	unmarPar :=FindFromQ(*golog,"fi_FI","porno")
	
//	fmt.Println(unmarPar)
	fmt.Println(unmarPar.Pushsite)

}
