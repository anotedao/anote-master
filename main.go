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

	// val := "%s%s__3AMz6wgnriNTfUtqfV3Mv27K6MXTNHTgX14__3AGtguyZWGoVLLUU2hxyKrucGLdTihBit1j"
	// dataTransaction("64.227.11.38", &val, nil, nil)

	// lease("3AMz6wgnriNTfUtqfV3Mv27K6MXTNHTgX14")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
