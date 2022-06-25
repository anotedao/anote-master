package main

import (
	"log"

	"github.com/wavesplatform/gowaves/pkg/client"
)

func initWaves() *client.Client {
	opts := client.Options{
		BaseUrl: "http://localhost:6869",
		ApiKey:  " ",
	}
	wc, err := client.NewClient(opts)
	if err != nil {
		log.Println(err.Error())
	}
	return wc
}
