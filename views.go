package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mo7zayed/reqip"
	macaron "gopkg.in/macaron.v1"
)

func pingView(ctx *macaron.Context) {
	pr := &PingResponse{
		Success: true,
	}

	ip := reqip.GetClientIP(ctx.Req.Request)
	addressOwner := ctx.Params("addr")
	addressNode := ctx.Params("addrnode")

	cp, _, err := wc.Peers.Connected(context.Background())
	if err != nil {
		log.Println(err)
	}

	for _, peer := range cp {
		if peer.Address.Addr.String() == ip {
			value, err := getData(addressOwner)
			if err != nil {
				log.Println(err)
			}
			if value == nil {
				val := fmt.Sprintf("%%s%%s__%s__%s", addressNode, ip)
				err := dataTransaction(addressOwner, &val, nil, nil)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	ctx.JSON(200, pr)
}

type PingResponse struct {
	Success bool `json:"success"`
}
