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
	for {
		// Create new HTTP client to send the transaction to public TestNet nodes
		cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
		if err != nil {
			log.Println(err)
		}

		// Context to cancel the request execution on timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pk, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
		if err != nil {
			log.Println(err)
		}
		addr, err := proto.NewAddressFromPublicKey(55, pk)
		if err != nil {
			log.Println(err)
		}

		opts := client.WithMatches(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

		de, _, err := cl.Addresses.AddressesData(ctx, addr, opts)
		if err != nil {
			log.Println(err)
		}

		for _, data := range de {
			addr := data.ToProtobuf().GetStringValue()

			ab, _, err := cl.Addresses.Balance(ctx, proto.MustAddressFromString(addr))
			if err != nil {
				log.Println(err.Error())
				return
			}

			if ab.Balance > MULTI8 {
				callDistributeReward(addr)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() {
	m := &Monitor{}
	go m.start()
}
