package main

import (
	"log"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	// val := "%s%s__3AUZKgXG36FNzvSuW4Q2Qhdv4zafUabZWwj__3AD2cHaZ2jFRgTyBHqcK67wZpjqCurZb7mn"
	// dataTransaction("64.227.176.28", &val, nil, nil)

	// lease("3AUZKgXG36FNzvSuW4Q2Qhdv4zafUabZWwj")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
