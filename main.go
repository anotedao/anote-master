package main

import (
	"log"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	// val := "%s%s__3AUYBormLLPHoxwjkiFYcDqWvvdzKwNinaU__3AA1ziNL6nYGrJr9hDw8KY5JKD15HQQzLbz"
	// dataTransaction("161.35.146.50", &val, nil, nil)

	// lease("3AUYBormLLPHoxwjkiFYcDqWvvdzKwNinaU")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
