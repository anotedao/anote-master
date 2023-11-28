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

	// val := "%s%s__3AKZGfBftW9py18vUX8cuNT5jM3aDd4N1xM__3AY3UN4EsZK6MMYxJUptoeX8uTGS6hr4QNy"
	// dataTransaction("164.92.80.82", &val, nil, nil)

	// lease("3AKZGfBftW9py18vUX8cuNT5jM3aDd4N1xM")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
