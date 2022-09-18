package main

import (
	"context"
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

			valueIp, err := getData(ip)
			if err != nil {
				log.Println(err.Error())
			}

			if value == nil && valueIp == nil && !done {
				val := addressOwner
				valIp := addressNode

				dataTransaction(addressNode, &val, nil, nil)

				dataTransaction(ip, &valIp, nil, nil)

				sendAsset(Fee*7, "", addressNode)

				go func() {
					waitForScript(addressNode)
					lease(addressNode)
				}()

				done = true
			} else if !done {
				log.Println("Ping not done: ", value, valueIp)
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
