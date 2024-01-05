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

	// val := "%s%s__3ARcT2sqE7V4kz1EdLuSgSciR9AT4jzWDmr__3AA1ziNL6nYGrJr9hDw8KY5JKD15HQQzLbz"
	// dataTransaction("188.25.72.126", &val, nil, nil)

	// lease("3ARcT2sqE7V4kz1EdLuSgSciR9AT4jzWDmr")

	startMonitor()
}
