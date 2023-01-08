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

	// lease("3AM8xxmwB9f8f8PUYpBdCCzhU1e7x4AR79o")

	key := "137.184.114.121"
	value := "%s%s__3AAV8552hNYiKrQ1TwUZohmdUeoouhPSXR6__3AQ7JCAQKWLjPnduYUuzGjHJuRWDXNW6UVx"
	dataTransaction(key, &value, nil, nil)

	key = "138.68.82.175"
	value = "%s%s__3AYDqADQzxnBCUWAgwYqa7HDHc7tTb9VFe1__3AFi9Xs6gJzfzeAheAaprXNReNq7WWvs1Qg"
	dataTransaction(key, &value, nil, nil)

	key = "159.223.215.254"
	value = "%s%s__3AAJ2VBwKooTUkq2rSWxrZ3dcYBUBPJc3g8__3AKCefhcrijSwwWM671ahhMrPVrE7Je3j4s"
	dataTransaction(key, &value, nil, nil)

	key = "161.35.20.45"
	value = "%s%s__3ASw3cTnQ3WWx7egu9urZQYG6SF12sAFobP__3AWo9PMERGweYgK2GjzscyMtq4VEqFSvF6j"
	dataTransaction(key, &value, nil, nil)

	key = "167.99.36.0"
	value = "%s%s__3AVfTvnYCNdLkSy1zMYQayWntsPbLjyKV7E__3AHWBJDGTQr1gB32gef2ggmY5yhWiBRvZhE"
	dataTransaction(key, &value, nil, nil)

	key = "167.99.36.116"
	value = "%s%s__3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM__3A9Rb3t91eHg1ypsmBiRth4Ld9ZytGwZe9p"
	dataTransaction(key, &value, nil, nil)

	key = "174.138.9.35"
	value = "%s%s__3AGPzYFwwbpyzQVcpRfMRvaxiaz2jQM5fMU__3AWhzpyhXDDVBBrTjcq7dgfmWNWsQFguXPB"
	dataTransaction(key, &value, nil, nil)

	key = "178.62.199.229"
	value = "%s%s__3APZKhnT9CoBg6Gk1pPidWifT4cvDcCwLxG__3AF4JjMnExbNYYxDF3AKes6Ce1M1NyuSYz7"
	dataTransaction(key, &value, nil, nil)

	key = "178.62.4.80"
	value = "%s%s__3ANHygcoZS7d4zRvcFrXcsMh8UCjFqUksF8__3ASLefwuE3dz9cW9bhP6ZC9N73pLqs2vPEH"
	dataTransaction(key, &value, nil, nil)

	key = "193.149.180.83"
	value = "%s%s__3AGAMvupfs74c48CjdPeAzRXvz3jsonuuWo__3ATvsMDMGEjx59WyVTT8CsXsHXAatZi2Sfq"
	dataTransaction(key, &value, nil, nil)

	key = "207.154.245.85"
	value = "%s%s__3ABzhn4LbMbFMu7ZUK8sHDbet8z3m1Q2jDd__3AG8wmARQR4pxhCqMcxXvzJ4vkcud89noNt"
	dataTransaction(key, &value, nil, nil)

	key = "45.153.48.47"
	value = "%s%s__3AYKGNVEJHb4AQ4XbNpQHSvfC3cYM4pY6FU__3AC3Pw3pNTxo1wEeKsz83JdFvcLtPMzi4ow"
	dataTransaction(key, &value, nil, nil)

	key = "64.225.103.253"
	value = "%s%s__3AM8xxmwB9f8f8PUYpBdCCzhU1e7x4AR79o__3AD6JBzYKQYBYhcLGd9Z1XXxnTBoSEnjPj8"
	dataTransaction(key, &value, nil, nil)

	key = "64.227.122.134"
	value = "%s%s__3ABabTX7Q8Xec4TsyLWrrdcwtt45cK1xYdt__3ABkYZ8dF17cdrUdi2aBbYpNZ1DdPmngGrZ"
	dataTransaction(key, &value, nil, nil)

}
