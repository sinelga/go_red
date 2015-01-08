package somephrases

import (
	"log/syslog"
	"math/rand"
	"time"
)

func GetSome(golog syslog.Writer, phrases []string, quant int) []string {

	outarrmap := make(map[string]struct{})
	var outarr []string
	min := int(0)
	max := len(phrases)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < quant; i++ {
//		rand.Seed(time.Now().UTC().UnixNano())
		rndint := rand.Intn(max-min) + min
		outarrmap[phrases[rndint]] = struct{}{}

	}

	for key, _ := range outarrmap {

		outarr = append(outarr, key)

	}

	return outarr

}
