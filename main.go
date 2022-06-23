package main

import (
	"fmt"
	"log"

	"github.com/mr-tron/base58"
)

var conf *Config

var nodeAddress string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	// pk, _ := crypto.NewPublicKeyFromBase58(conf.PublicKey)

	// a, _ := proto.NewAddressFromPublicKey(55, pk)

	// log.Println(a.String())

	// mess := "%s%s__3AP5zrscjAo7Jb8urXd6836XRTGakZf7Z9B__178.62.103.70"

	// err := dataTransaction("3ASLefwuE3dz9cW9bhP6ZC9N73pLqs2vPEH", &mess, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

	// sum := 0

	// abdr, _ := gowaves.WNC.AssetsBalanceDistribution("Gvs59WEEXVAQiRZwisUosG7fVNr8vnzS8mjkgqotrERT", 3166500, 100, "")
	// for a, _ := range abdr.Items {
	// 	abr, _ := gowaves.WNC.AssetsBalance(a, "66DUhUoJaoZcstkKpcoN3FUcqjB6v8VJd5ZQd6RsPxhv")
	// 	if abr.Balance > 0 {
	// 		log.Println(a)
	// 		if a != "3PBmmxKhFcDhb8PrDdCdvw2iGMPnp7VuwPy" && a != "3P2EtZMgEN4W49hLXy966D53oHiE52gawhn" && a != "3PDb1ULFjazuzPeWkF2vqd1nomKh4ctq9y2" && a != "3PQEVtX7SukU7zVfpgkKDmnrX7NFw1pHBVd" {
	// 			sum += int(abr.Balance)
	// 			log.Println(abr.Balance)
	// 		}
	// 	}

	// }

	// log.Println(sum)

	// ride := []byte("WAVES")
	// decoded, _ := base58.Decode("WAVES")
	// log.Println(decoded)
	encoded, _ := base58.Decode("PrpurRJQhiW3G5ezdGjpdzaZBFDkRvCAhLQXb2UG7WVwexxqwgT8S17tFQLzuH2dzqsEtooCLEXs9iqqRrNW7JEZukFv3rgEeHwfNqLZ3JfQVeYZBFSLkUvt5VaLCDh")
	// Show the encoded data.
	fmt.Println("Encoded Data:", string(encoded))
}
