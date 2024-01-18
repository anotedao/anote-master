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

	// val := "%s%s__3ATm4u9ANgtUd4R4fXhuy9FCUjzARKCmhir__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6"
	// dataTransaction("188.25.138.155", &val, nil, nil)

	// lease("3ATm4u9ANgtUd4R4fXhuy9FCUjzARKCmhir")

	startMonitor()
}
