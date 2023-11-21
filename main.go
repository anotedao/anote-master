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

	// val := "%s%s__3AJXGdrAgPYt218fNNwx15BW4EdyHeuo1hq__3AGtguyZWGoVLLUU2hxyKrucGLdTihBit1j"
	// dataTransaction("157.245.253.62", &val, nil, nil)

	// lease("3AJXGdrAgPYt218fNNwx15BW4EdyHeuo1hq")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
