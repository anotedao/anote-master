package main

import (
	"context"
	"fmt"
	"log"

	macaron "gopkg.in/macaron.v1"
)

func pingView(ctx *macaron.Context) {
	pr := &PingResponse{
		Success: true,
	}

	ip, err := getIP(ctx.Req.Request)
	if err != nil {
		log.Println(err.Error())
	}
	addressOwner := ctx.Params("addr")
	addressNode := ctx.Params("addrnode")

	log.Println(ip)

	cp, _, err := wc.Peers.Connected(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	done := false

	// for _, peer := range cp {
	for range cp {
		// if peer.Address.Addr.String() == ip {
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
				txid, err := lease(addressOwner)
				if err != nil {
					log.Println(err.Error())
				} else {
					val += fmt.Sprintf("__%s", txid)
					err := dataTransaction(addressNode, &val, nil, nil)
					if err != nil {
						log.Println(err.Error())
					}
					sendAsset(Fee*7, "", addressOwner)
					done = true
				}
			}
			// }
		}
	}

	ctx.JSON(200, pr)
}

type PingResponse struct {
	Success bool `json:"success"`
}
