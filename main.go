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

	// val := "%s%s__3AXK9XwaMjwJMrdjeYbDx4pZgaVN2mpimh7__3AKCefhcrijSwwWM671ahhMrPVrE7Je3j4s"
	// dataTransaction("188.25.72.128", &val, nil, nil)

	// lease("3AXK9XwaMjwJMrdjeYbDx4pZgaVN2mpimh7")

	startMonitor()
}
