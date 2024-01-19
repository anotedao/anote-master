package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type Monitor struct{}

func (m *Monitor) start() {
	cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}, ApiKey: " "})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	for {
		// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		ctx := context.Background()
		// defer cancel()

		pk, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}
		addr, err := proto.NewAddressFromPublicKey(55, pk)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		// opts := client.WithMatches(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

		de, resp, err := cl.Addresses.AddressesData(ctx, addr)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}
		resp.Body.Close()

		log.Println(addr)

		log.Println(prettyPrint(de))

		for _, data := range de {
			addr := data.GetKey()
			// addr := data.ToProtobuf().GetStringValue()
			// addr = strings.Split(addr, "__")[1]

			// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			// defer cancel()
			ctx := context.Background()

			log.Println(addr)

			ab, resp, err := cl.Addresses.Balance(ctx, proto.MustAddressFromString(addr))
			if err != nil {
				log.Println(err.Error())
				logTelegram(err.Error())
			}
			resp.Body.Close()

			if ab != nil && ab.Balance >= MULTI8 {
				callDistributeReward(addr)
			}

			ctx = nil
		}

		de = nil
		ctx = nil

		time.Sleep(time.Second * MonitorTick)
	}
}

func startMonitor() {
	m := &Monitor{}
	m.start()
}
