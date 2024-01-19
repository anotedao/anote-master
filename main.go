package main

import (
	"log"
	// _ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	conf = initConfig()

	// initWaves()

	// val := "%s%s__3AF7rRTQrDazDvdrvcdRCFaNKZSn21qikNG__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6"
	// dataTransaction("188.25.138.156", &val, nil, nil)

	// lease("3AF7rRTQrDazDvdrvcdRCFaNKZSn21qikNG")

	startMonitor()

	// val := "%s%d__3AY3UN4EsZK6MMYxJUptoeX8uTGS6hr4QNy__550000"
	// dataTransaction("3AYAhn6tKGssduMsCvnjEPa2MMSduRPW2gr", &val, nil, nil)

	// val = "%s%d__3ATqxdbGWSp1tKtPSkZ2pX52a1RWFjNrGZx__550000"
	// dataTransaction("3AUgfHWfFfv2asmizunrdzQtvn2xEWKL8hh", &val, nil, nil)

	// val = "%s%d__3AGtguyZWGoVLLUU2hxyKrucGLdTihBit1j__550000"
	// dataTransaction("3AJXGdrAgPYt218fNNwx15BW4EdyHeuo1hq", &val, nil, nil)

	// val = "%s%d__3AKCefhcrijSwwWM671ahhMrPVrE7Je3j4s__550000"
	// dataTransaction("3AAJ2VBwKooTUkq2rSWxrZ3dcYBUBPJc3g8", &val, nil, nil)

	// // val = "%s%d__3ABxPm5GGwC6CR6g9YTx2d49EsBdejG2eou__550000"
	// // dataTransaction("3AC3LNRYZs7bj95D3bKruRFLdUmCJTmaidr", &val, nil, nil)

	// val = "%s%d__3AA1ziNL6nYGrJr9hDw8KY5JKD15HQQzLbz__550000"
	// dataTransaction("3AUYBormLLPHoxwjkiFYcDqWvvdzKwNinaU", &val, nil, nil)

	// val = "%s%d__3AY3UN4EsZK6MMYxJUptoeX8uTGS6hr4QNy__550000"
	// dataTransaction("3AKZGfBftW9py18vUX8cuNT5jM3aDd4N1xM", &val, nil, nil)

	// val = "%s%d__3AEvVZB8ecyLkYZt8QFoPxy5q1zKoQoifJP__550000"
	// dataTransaction("3A9i3q423TrD1GrV5NUF7fKPtanWms58dH5", &val, nil, nil)

	// val = "%s%d__3AGtguyZWGoVLLUU2hxyKrucGLdTihBit1j__550000"
	// dataTransaction("3AArU5bWZ7xZBR2ygoB9nKiYQzvGFiq3PXF", &val, nil, nil)

	// val = "%s%d__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6__550000"
	// dataTransaction("3ADvaNaDR7K3jrBumDjFkyTdxqkPAv8Jc3W", &val, nil, nil)

	// val = "%s%d__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6__550000"
	// dataTransaction("3ATm4u9ANgtUd4R4fXhuy9FCUjzARKCmhir", &val, nil, nil)

	// val = "%s%d__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6__550000"
	// dataTransaction("3AF7rRTQrDazDvdrvcdRCFaNKZSn21qikNG", &val, nil, nil)

	// val = "%s%d__3AShXVgRcRis82CwD7o9pz1Ac9vmRYMqELT__550000"
	// dataTransaction("3ACJZCJTDLSGu4n9i8UtnJgWaEDppXUodXs", &val, nil, nil)

	// val = "%s%d__3AA1ziNL6nYGrJr9hDw8KY5JKD15HQQzLbz__550000"
	// dataTransaction("3AV4aJ82rmTwRnd78wgVRvNLAXnvyHtcjF5", &val, nil, nil)

	// val = "%s%d__3AA1ziNL6nYGrJr9hDw8KY5JKD15HQQzLbz__550000"
	// dataTransaction("3ARcT2sqE7V4kz1EdLuSgSciR9AT4jzWDmr", &val, nil, nil)

	// val = "%s%d__3A9Rb3t91eHg1ypsmBiRth4Ld9ZytGwZe9p__550000"
	// dataTransaction("3AThTowRX96bQ22oXo2XPWWpp8C1kRoQMdQ", &val, nil, nil)

	// val = "%s%d__3ALJLTzt3Q6uhosTdqvf935U3SwHnqyNdyH__550000"
	// dataTransaction("3AGfi22WcZrAWrYjtD8V3a5PV1Z3b1pt15z", &val, nil, nil)

	// val = "%s%d__3AKCefhcrijSwwWM671ahhMrPVrE7Je3j4s__550000"
	// dataTransaction("3AXK9XwaMjwJMrdjeYbDx4pZgaVN2mpimh7", &val, nil, nil)

	// val = "%s%d__3AVB7NmMDnDM16VJxNw9ZE2mt4SFz5wepsg__550000"
	// dataTransaction("3ACFS96MCnpQWG1Jca77euUcY7hRF3gtGy5", &val, nil, nil)

	// val = "%s%d__3ALJLTzt3Q6uhosTdqvf935U3SwHnqyNdyH__550000"
	// dataTransaction("3AMRkeWMB18Az4eGEdQmncWMCzgL9Y6aoc7", &val, nil, nil)

	// val = "%s%d__3APCiVkxSRNzJsdSbUAMQnoZBkhSE3nyd9U__550000"
	// dataTransaction("3APqq58wCCkfkm9QJNKZBbF6xJR46zpPLes", &val, nil, nil)

	// val = "%s%d__3APfqnJXpucsE8W8PCJ7UnaVZtrKEHQVpWo__550000"
	// dataTransaction("3AB2QTHWWavvjXiGWmBRnTXyqWEcdC8S69i", &val, nil, nil)

	// val = "%s%d__3AShXVgRcRis82CwD7o9pz1Ac9vmRYMqELT__550000"
	// dataTransaction("3APWcSs37en5WatNYqW3zf4UAbjriGd9nv6", &val, nil, nil)

	// val = "%s%d__3AWTbhsG5Lqbx5Bwa6LokS8SuvWXvrAKAFr__550000"
	// dataTransaction("3AN54JjyzDW4w3r7Jat3vTdw3CsP8j7KL6c", &val, nil, nil)

	// val = "%s%d__3AM2KY5qur6ms2yP2KeUwFgoKBApnwmGaEn__550000"
	// dataTransaction("3A9zxBpyTGJFxsspy7gpBhqiEXywtpxuuNk", &val, nil, nil)

	// val = "%s%d__3AMHWXNjR95eVAWL3U1o23VR6adceCayoyJ__550000"
	// dataTransaction("3ALr1kmyz3JZXKnbqSWGP1xU3UKXQLEs9f1", &val, nil, nil)
}
