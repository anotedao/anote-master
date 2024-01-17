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

	// val := "%s%s__3ADvaNaDR7K3jrBumDjFkyTdxqkPAv8Jc3W__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6"
	// dataTransaction("188.25.138.154", &val, nil, nil)

	// lease("3ADvaNaDR7K3jrBumDjFkyTdxqkPAv8Jc3W")
	// leaseCancel("DuXMGMsrj5tLrTkCgX9Zr3W8eK1jCi19bY2Fs1P7McUL")

	startMonitor()
}
