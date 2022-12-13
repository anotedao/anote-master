package main

import (
	macaron "gopkg.in/macaron.v1"
)

func pingView(ctx *macaron.Context) {
	// addressOwner := ctx.Params("addr")
	// addressNode := ctx.Params("addrnode")
	// ip := ctx.Params("ip")

	// cp, _, err := wc.Peers.Connected(context.Background())
	// if err != nil {
	// 	log.Println(err.Error())
	// 	logTelegram(err.Error())
	// }

	done := true

	// for _, peer := range cp {
	// 	if peer.Address.Addr.String() == ip {
	// 		valueIp, err := getData(ip, nil)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 			// logTelegram(err.Error())
	// 		}

	// 		value := "%s%s__" + addressNode + "__" + addressOwner

	// 		if valueIp == nil && !done {
	// 			dataTransaction(ip, &value, nil, nil)

	// 			sendAsset(Fee*7, "", addressNode)

	// 			go func() {
	// 				waitForScript(addressNode)
	// 				lease(addressNode)
	// 			}()

	// 			done = true
	// 		} else if !done {
	// 			log.Println("Ping not done: ", value)
	// 		}
	// 	}
	// }

	pr := &PingResponse{
		Success: done,
	}

	ctx.JSON(200, pr)
}

type PingResponse struct {
	Success bool `json:"success"`
}
