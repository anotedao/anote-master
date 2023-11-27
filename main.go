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

	// val := "%s%s__3AYAhn6tKGssduMsCvnjEPa2MMSduRPW2gr__3AY3UN4EsZK6MMYxJUptoeX8uTGS6hr4QNy"
	// dataTransaction("157.230.229.87", &val, nil, nil)

	// lease("3AYAhn6tKGssduMsCvnjEPa2MMSduRPW2gr")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
