package main

import (
	"log"

	"github.com/wavesplatform/gowaves/pkg/client"
)

func initWaves() *client.Client {
	opts := client.Options{
		BaseUrl: AnoteNodeURL,
		ApiKey:  " ",
	}
	wc, err := client.NewClient(opts)
	if err != nil {
		log.Println(err.Error())
	}
	return wc
}
