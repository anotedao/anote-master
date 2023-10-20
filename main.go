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

	// val := "%s%s__3AHtdF53rn9JAXWZ7z62NC8PB8GYdj2YinH__3AEBPBgPx59KfgjyFvhS6mR6APSPfae38J6"
	// dataTransaction("144.24.145.97", &val, nil, nil)

	// lease("3AHtdF53rn9JAXWZ7z62NC8PB8GYdj2YinH")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
