package somekeywords

import (
	"log/syslog"
	"math/rand"
	"time"
)

func GetSome(golog syslog.Writer, keywords []string, quant int) []string {

	outarrmap := make(map[string]struct{})
	var outarr []string
	min := int(0)
	max := len(keywords)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < quant; i++ {
//		rand.Seed(time.Now().UTC().UnixNano())
		rndint := rand.Intn(max-min) + min
		outarrmap[keywords[rndint]] = struct{}{}

	}

	for key, _ := range outarrmap {

		outarr = append(outarr, key)

	}

	return outarr

}
