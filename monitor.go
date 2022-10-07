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
		// Create new HTTP client to send the transaction to public TestNet nodes
		cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}, ApiKey: " "})
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		// Context to cancel the request execution on timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

		peers, _, err := cl.Peers.Connected(ctx)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		for _, data := range de {
			addr := data.ToProtobuf().GetStringValue()

			ab, _, err := cl.Addresses.Balance(ctx, proto.MustAddressFromString(addr))
			if err != nil {
				log.Println(err.Error())
				logTelegram(err.Error())
			}

			if ab.Balance > MULTI8 {
				callDistributeReward(addr)
			}

			found := false

			for _, peer := range peers {
				if strings.Contains(peer.Address.String(), data.GetKey()) {
					found = true
				}
			}

			if !found {
				dataTransaction(data.ToProtobuf().GetStringValue(), nil, nil, nil)
				dataTransaction(data.GetKey(), nil, nil, nil)
			}
		}

		leases, _, err := cl.Leasing.Active(ctx, addr)

		for _, lease := range leases {
			found := false
			for _, data := range de {
				addr := data.ToProtobuf().GetStringValue()
				if addr == lease.Recipient.String() {
					found = true
				}
			}
			if !found {
				leaseCancel(lease.ID.String())
			}
		}

		log.Println(len(leases))

		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() {
	m := &Monitor{}
	go m.start()
}
