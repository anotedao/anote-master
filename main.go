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

	// val := "%s%s__3APqq58wCCkfkm9QJNKZBbF6xJR46zpPLes__3APCiVkxSRNzJsdSbUAMQnoZBkhSE3nyd9U"
	// dataTransaction("24.144.85.235", &val, nil, nil)

	// lease("3APqq58wCCkfkm9QJNKZBbF6xJR46zpPLes")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
