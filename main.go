package main

import (
	"log"
	// _ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	// startMonitor()

	// leaseCancel("FXZ36ZEdEVbkTVohnJjEeJbVTRjwCyhaC2kZEi838LLR")
	// leaseCancel("B3dVWFM8mEyt9xJ9JuZKLrkAtaEGByZen8geYXzJbfqf")
	// leaseCancel("8Ca2mX56VQayj5rAnffGR6J9NDHY19yLL1NKGerkKoBK")
	// leaseCancel("BpxrYokyi4QW2k8qJkiGJV5t7fkpmCVVae2Un7pEVjyQ")
	// leaseCancel("FNR5pyKQ98YgkgiZtd4iTdX3EMxaWF2wRXGaMq9jTeYL")

	// dataTransaction("3APSKZokE7Hg7dnVDmVVWVQpqVEgiVW9XG9", nil, nil, nil)
}
