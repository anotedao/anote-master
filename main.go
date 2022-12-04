package main

import (
	"log"

	"github.com/wavesplatform/gowaves/pkg/client"
	"gopkg.in/macaron.v1"
)

var conf *Config

var nodeAddress string

var m *macaron.Macaron

var wc *client.Client

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	// m = initMacaron()

	// wc = initWaves()

	// initMonitor()

	// cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}, ApiKey: " "})
	// if err != nil {
	// 	log.Println(err)
	// 	logTelegram(err.Error())
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// defer cancel()

	// pk, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	// if err != nil {
	// 	log.Println(err)
	// 	logTelegram(err.Error())
	// }
	// addr, err := proto.NewAddressFromPublicKey(55, pk)
	// if err != nil {
	// 	log.Println(err)
	// 	logTelegram(err.Error())
	// }

	// opts := client.WithMatches(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	// de, _, err := cl.Addresses.AddressesData(ctx, addr, opts)
	// if err != nil {
	// 	log.Println(err)
	// 	logTelegram(err.Error())
	// }

	// list := make(map[string]int)
	// // list2 := make(map[string]int)

	// for _, data := range de {
	// 	addr := data.ToProtobuf().GetStringValue()
	// 	owner := strings.Split(addr, "__")[2]
	// 	addr = strings.Split(addr, "__")[1]

	// 	// log.Println(addr + " " + owner)

	// 	count := list[owner]

	// 	if count == 0 {
	// 		list[owner] = 1
	// 	} else {
	// 		list[owner]++
	// 	}
	// }

	// for a, b := range list {
	// 	if b > 1 {
	// 		for _, data := range de {
	// 			addr := data.ToProtobuf().GetStringValue()
	// 			owner := strings.Split(addr, "__")[2]
	// 			addr = strings.Split(addr, "__")[1]

	// 			if owner == a {
	// 				log.Println(owner + " " + addr)

	// 				ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// 				defer cancel()
	// 				leases, _, err := cl.Leasing.Active(ctx, proto.MustAddressFromString(addr))
	// 				if err != nil {
	// 					log.Println(err.Error())
	// 				}
	// 				for _, l := range leases {
	// 					leaseCancel(l.ID.String())
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// log.Println(counter)

	// log.Println(prettyPrint(list))

	// m.Run("127.0.0.1", Port)
}
