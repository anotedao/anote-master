package main

import (
	"log"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	// val := "%s%s__3AVJ1R8481NeiPiaVn8HRXCU6Mp1GE6q8rV__3AWQfMMTNWYaPsqs4vmYb1SMhh3zy4RD5Vr"
	// dataTransaction("146.190.145.57", &val, nil, nil)

	// lease("3AVJ1R8481NeiPiaVn8HRXCU6Mp1GE6q8rV")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
