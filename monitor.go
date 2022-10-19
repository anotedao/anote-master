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

			if ab != nil && ab.Balance > MULTI8 {
				callDistributeReward(addr)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func (m *Monitor) monitor() {
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

		peers, _, err := cl.Peers.All(ctx)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		leases, _, err := cl.Leasing.Active(ctx, addr)
		if err != nil {
			log.Println(err.Error())
		}

		log.Println(len(peers))
		log.Println(len(leases))

		for _, l := range leases {
			found := false
			// minerStr := ""
			ipStr := ""
			lsStr := ""
			for _, p := range peers {
				miner, _ := getData(p.Address.Addr.String(), nil)
				if miner != nil && strings.Contains(miner.(string), l.Recipient.String()) {
					// minerStr = miner.(string)
					ipStr = p.Address.Addr.String()
					ls := time.Unix(int64(p.LastSeen)/1000, 0)
					lsStr = ls.String()
					if time.Since(ls) < time.Hour {
						found = true
					}
					// log.Println(l.Recipient.String() + " " + p.Address.Addr.String() + " " + ls.String())
				}
			}

			if !found {
				log.Println(l.Recipient.String() + " " + ipStr + " " + lsStr)
			}
		}

		log.Println("Done.")

		time.Sleep(time.Second * MonitorTick)
	}
}

func initMonitor() {
	m := &Monitor{}
	go m.start()
	go m.monitor()
}
