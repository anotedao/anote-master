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

	// val := "%s%s__3ACJZCJTDLSGu4n9i8UtnJgWaEDppXUodXs__3AShXVgRcRis82CwD7o9pz1Ac9vmRYMqELT"
	// dataTransaction("188.25.72.123", &val, nil, nil)

	// lease("3AThTowRX96bQ22oXo2XPWWpp8C1kRoQMdQ")
	// lease("3AUYBormLLPHoxwjkiFYcDqWvvdzKwNinaU")
	// lease("3ATYNXJA6xe3mBrXSQU4rKyDZSkT8MKUowU")
	// lease("3AC3LNRYZs7bj95D3bKruRFLdUmCJTmaidr")
	// lease("3AJXGdrAgPYt218fNNwx15BW4EdyHeuo1hq")

	startMonitor()
}
