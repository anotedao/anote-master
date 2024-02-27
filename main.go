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

	// startMonitor()

	// dataTransaction("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM", nil, nil, nil)
}
