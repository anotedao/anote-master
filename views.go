package main

import (
	"context"
	"fmt"
	"log"

	macaron "gopkg.in/macaron.v1"
)

func pingView(ctx *macaron.Context) {
	addressOwner := ctx.Params("addr")
	addressNode := ctx.Params("addrnode")
	ip := ctx.Params("ip")

	cp, _, err := wc.Peers.Connected(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	done := false

	for _, peer := range cp {
		// for range cp {
		if peer.Address.Addr.String() == ip {
			value, err := getData(addressNode)
			if err != nil {
				log.Println(err.Error())
			}
			if value == nil && !done {
				val := fmt.Sprintf("%%s%%s__%s__%s", addressOwner, ip)
				err := dataTransaction(addressNode, &val, nil, nil)
				if err != nil {
					log.Println(err.Error())
				} else {
					txid, err := lease(addressNode)
					if err != nil {
						log.Println(err.Error())
					} else {
						val += fmt.Sprintf("__%s", txid)
						err := dataTransaction(addressNode, &val, nil, nil)
						if err != nil {
							log.Println(err.Error())
						}
						sendAsset(Fee*7, "", addressNode)
						done = true
					}
				}
			}
		}
	}

	pr := &PingResponse{
		Success: done,
	}

	ctx.JSON(200, pr)
}

type PingResponse struct {
	Success bool `json:"success"`
}
