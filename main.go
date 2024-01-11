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

	// val := "%s%s__3AMRkeWMB18Az4eGEdQmncWMCzgL9Y6aoc7__3ALJLTzt3Q6uhosTdqvf935U3SwHnqyNdyH"
	// dataTransaction("206.189.59.181", &val, nil, nil)

	// lease("3AN54JjyzDW4w3r7Jat3vTdw3CsP8j7KL6c")
	// leaseCancel("DuXMGMsrj5tLrTkCgX9Zr3W8eK1jCi19bY2Fs1P7McUL")

	startMonitor()
}
