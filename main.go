package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"flag"
)

func main(){
	sPort := flag.Int("sport", 8080, "redirect from port (default:8080)")
	dPort := flag.Int("dport", 80, "redirect to port (default:80)")
	rCode := flag.Int("rcode", 301, "responce code (default:301)")
	Template := flag.String("template", "http://[%HOST%]:[%PORT%]/[%PATH%]", "template to redirect")
	flag.Parse()

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		host := strings.Split(r.Host,":")[0]
		path := r.URL.Path[1:]
		url := *Template
		url = strings.Replace(url, "[%HOST%]", host, -1)
		url = strings.Replace(url, "[%PORT%]", fmt.Sprintf("%d", *dPort), -1)
		url = strings.Replace(url, "[%PATH%]", path, -1)
		http.Redirect(w, r, url, *rCode)
		log.Printf("%v %v => %v", r.RemoteAddr, r.RequestURI, url)
	})

	log.Printf("Go redirect listening %d...", *sPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *sPort), nil)
}
