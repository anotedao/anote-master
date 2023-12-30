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
	// lease("3A9i3q423TrD1GrV5NUF7fKPtanWms58dH5")
	// lease("3AN54JjyzDW4w3r7Jat3vTdw3CsP8j7KL6c")
	// lease("3APqq58wCCkfkm9QJNKZBbF6xJR46zpPLes")

	// leaseCancel("BTmizMBFcmZV6tzagUVCjAZCedduQTycFyYNR1vBSwb4")
	// leaseCancel("7LFwQTLha2HTtup3LFxtPK6Q8uk3xWAvXLe3EWmyZKou")
	// leaseCancel("6igZetgfATBvKL2nmo9SFezD2CxMJ7BxVfqR1UZ46Fom")

	startMonitor()
}
