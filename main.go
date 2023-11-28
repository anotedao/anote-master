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

	// val := "%s%s__3AArU5bWZ7xZBR2ygoB9nKiYQzvGFiq3PXF__3AGtguyZWGoVLLUU2hxyKrucGLdTihBit1j"
	// dataTransaction("165.232.130.25", &val, nil, nil)

	// lease("3AArU5bWZ7xZBR2ygoB9nKiYQzvGFiq3PXF")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	startMonitor()
}
