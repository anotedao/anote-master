package main

import (
	"log"
	// _ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	conf = initConfig()

	// initWaves()

	// val := "%s%s__3A9i3q423TrD1GrV5NUF7fKPtanWms58dH5__3AEvVZB8ecyLkYZt8QFoPxy5q1zKoQoifJP"
	// dataTransaction("165.22.90.138", &val, nil, nil)

	// lease("3A9i3q423TrD1GrV5NUF7fKPtanWms58dH5")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
