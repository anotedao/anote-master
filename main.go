package main

import (
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	conf = initConfig()

	// initWaves()

	// val := "%s%s__3AVEnVqk9RpNj89Kbid1o68mXw4rR3SrsKw__3AEBPBgPx59KfgjyFvhS6mR6APSPfae38J6"
	// dataTransaction("161.35.146.50", &val, nil, nil)

	// lease("3AVEnVqk9RpNj89Kbid1o68mXw4rR3SrsKw")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
