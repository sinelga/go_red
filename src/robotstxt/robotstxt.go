package robotstxt

import (
"bytes"
	"log/syslog"

)

func Createrobotstxt(golog syslog.Writer,site string) []byte{
	
	var buffer bytes.Buffer
	
	buffer.WriteString("User-agent: *\n")
	buffer.WriteString("Allow: /\n")
	buffer.WriteString("Sitemap: http://"+site+"/sitemap.xml\n")
	
	return buffer.Bytes() 
	
}

