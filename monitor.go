package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type Monitor struct{}

func (m *Monitor) start() {
	for {
		cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}, ApiKey: " "})
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

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

		opts := client.WithMatches(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

		de, _, err := cl.Addresses.AddressesData(ctx, addr, opts)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		for _, data := range de {
			addr := data.ToProtobuf().GetStringValue()
			addr = strings.Split(addr, "__")[1]

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			ab, _, err := cl.Addresses.Balance(ctx, proto.MustAddressFromString(addr))
			if err != nil {
				log.Println(err.Error())
				logTelegram(err.Error())
			}

			if ab != nil && ab.Balance > (MULTI8+RewardFee) {
				callDistributeReward(addr)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func startMonitor() {
	m := &Monitor{}
	m.start()
}
