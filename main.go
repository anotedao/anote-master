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

	// lease("3AJJjHbfw9MtY58HFi8oi1ajhGbWeKAgMKb")

	// leaseCancel("H6Ks1kQVQWBgoCEYRHg5S9zQLBPcSfGmz2TaJTm2Zucm")
	// leaseCancel("GDicpYUefUenvUY3NDP6zMD3zDz7uyef2doeecjHzBHM")

	// lease("3AAJ2VBwKooTUkq2rSWxrZ3dcYBUBPJc3g8")
	// lease("3APWcSs37en5WatNYqW3zf4UAbjriGd9nv6")

	startMonitor()
}
