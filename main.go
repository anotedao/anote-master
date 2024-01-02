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

	// val := "%s%s__3AJJjHbfw9MtY58HFi8oi1ajhGbWeKAgMKb__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6"
	// dataTransaction("188.25.72.100", &val, nil, nil)

	// lease("3AYAhn6tKGssduMsCvnjEPa2MMSduRPW2gr")
	// lease("3AArU5bWZ7xZBR2ygoB9nKiYQzvGFiq3PXF")
	// lease("3AKZGfBftW9py18vUX8cuNT5jM3aDd4N1xM")
	// lease("3AC3LNRYZs7bj95D3bKruRFLdUmCJTmaidr")
	// lease("3AJXGdrAgPYt218fNNwx15BW4EdyHeuo1hq")

	startMonitor()
}
