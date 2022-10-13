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

		CommunityAddress := addr

		opts := client.WithMatches(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

		de, _, err := cl.Addresses.AddressesData(ctx, addr, opts)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		peers, _, err := cl.Peers.All(ctx)
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

			if ab != nil && ab.Balance > MULTI8 {
				callDistributeReward(addr)
			}

			found := false

			for _, peer := range peers {
				ls := time.Unix(int64(peer.LastSeen)/1000, 0)
				if strings.Contains(peer.Address.String(), data.GetKey()) && time.Since(ls) < time.Hour*10 {
					found = true
				}
			}

			if !found {
				dataTransaction(data.ToProtobuf().GetStringValue(), nil, nil, nil)
				dataTransaction(data.GetKey(), nil, nil, nil)
				log.Println(data.ToProtobuf().GetStringValue() + " " + data.GetKey())
				logTelegram(data.ToProtobuf().GetStringValue() + " " + data.GetKey())
			}
		}

		leases, _, err := cl.Leasing.Active(ctx, CommunityAddress)
		if err != nil {
			log.Println(err.Error())
			logTelegram(err.Error())
		}

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
				log.Println("Cancel lease: " + lease.Recipient.String())
				logTelegram("Cancel lease: " + lease.Recipient.String())
			}
		}

		log.Println(len(leases))

		for _, peer := range peers {
			miner, err := getData(peer.Address.Addr.String())
			ls := time.Unix(int64(peer.LastSeen)/1000, 0)
			if err != nil && miner == nil && time.Since(ls) < time.Hour {
				log.Println(peer.Address.Addr)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() {
	m := &Monitor{}
	go m.start()
}
