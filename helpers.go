package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func dataTransaction(key string, valueStr *string, valueInt *int64, valueBool *bool) error {
	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Current time in milliseconds
	ts := time.Now().Unix() * 1000

	tr := proto.NewUnsignedDataWithProofs(2, sender, Fee, uint64(ts))

	if valueStr == nil && valueInt == nil && valueBool == nil {
		tr.Entries = append(tr.Entries,
			&proto.DeleteDataEntry{
				Key: key,
			},
		)
	}

	if valueStr != nil {
		tr.Entries = append(tr.Entries,
			&proto.StringDataEntry{
				Key:   key,
				Value: *valueStr,
			},
		)
	}

	if valueInt != nil {
		tr.Entries = append(tr.Entries,
			&proto.IntegerDataEntry{
				Key:   key,
				Value: *valueInt,
			},
		)
	}

	if valueBool != nil {
		tr.Entries = append(tr.Entries,
			&proto.BooleanDataEntry{
				Key:   key,
				Value: *valueBool,
			},
		)
	}

	err = tr.Sign(55, sk)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = cl.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	return nil
}

func getData(key string) (interface{}, error) {
	pk, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		return nil, err
	}

	a, err := proto.NewAddressFromPublicKey(55, pk)
	if err != nil {
		return nil, err
	}

	ad, _, err := wc.Addresses.AddressesDataKey(context.Background(), a, key)
	if err != nil {
		return nil, err
	}

	if ad.GetValueType().String() == "string" {
		return ad.ToProtobuf().GetStringValue(), nil
	}

	if ad.GetValueType().String() == "boolean" {
		return ad.ToProtobuf().GetBoolValue(), nil
	}

	if ad.GetValueType().String() == "integer" {
		return ad.ToProtobuf().GetIntValue(), nil
	}

	return "", nil
}

// NewSHA256 ...
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func lease(address string) (txid string, err error) {
	networkByte := byte(55)
	nodeURL := AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	rec, err := proto.NewRecipientFromString(address)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	tr := proto.NewUnsignedLeaseWithSig(sender, rec, 1000*MULTI8, Fee, ts)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return "", err
	}

	return tr.ID.String(), nil
}

func leaseCancel(txid string) {
	networkByte := byte(55)
	nodeURL := AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	lease, err := crypto.NewDigestFromBase58(txid)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// tr := proto.NewUnsignedLeaseWithSig(sender, rec, 1000*MULTI8, Fee, ts)
	tr := proto.NewUnsignedLeaseCancelWithSig(sender, lease, Fee, ts)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

func sendAsset(amount uint64, assetId string, recipient string) error {
	var networkByte byte
	var nodeURL string

	networkByte = 55
	nodeURL = AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Current time in milliseconds
	ts := time.Now().Unix() * 1000

	asset, err := proto.NewOptionalAssetFromString(assetId)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	assetW, err := proto.NewOptionalAssetFromString("")
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	rec, err := proto.NewAddressFromString(recipient)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	tr := proto.NewUnsignedTransferWithSig(sender, *asset, *assetW, uint64(ts), amount, Fee, proto.Recipient{Address: &rec}, nil)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	return nil
}

func waitForScript(address string) {
	cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return
	}

	a, err := proto.NewAddressFromString(address)
	if err != nil {
		log.Println(err.Error())
		logTelegram(err.Error())
		return
	}

	script := ""

	for len(script) == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		asi, _, err := cl.Addresses.ScriptInfo(ctx, a)
		if err != nil {
			log.Println(err.Error())
			logTelegram(err.Error())
			return
		}
		script = asi.Script

		time.Sleep(time.Second * 2)
	}
}

func callDistributeReward(address string) error {
	var networkByte = byte(55)
	var nodeURL = AnoteNodeURL

	// Create new HTTP client to send the transaction to public TestNet nodes
	cl, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	addr, err := proto.NewAddressFromString(address)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	abi, _, err := cl.Addresses.Balance(ctx, addr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(string(getPublicKey(address)))
	if err != nil {
		log.Println(err.Error() + " " + address)
		logTelegram(err.Error())
		return err
	}

	rec, err := proto.NewRecipientFromString("3AVkEwYsZeooN1GEc81a66N2zmnKFw1ZxyB")
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	call := proto.FunctionCall{
		Name:      "distributeMinerReward",
		Arguments: proto.Arguments{},
	}

	payments := proto.ScriptPayments{}
	payments.Append(proto.ScriptPayment{
		Amount: abi.Balance - RewardFee,
	})

	fa := proto.OptionalAsset{}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	tr := proto.NewUnsignedInvokeScriptWithProofs(
		2,
		networkByte,
		sender,
		rec,
		call,
		payments,
		fa,
		RewardFee,
		ts)

	tr.Proofs = proto.NewProofs()

	// // Send the transaction to the network
	_, err = cl.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
		return err
	}

	return nil
}

func getPublicKey(address string) string {
	pk := ""

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a, err := proto.NewAddressFromString(address)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	transactions, _, err := client.Transactions.Address(ctx, a, 100)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	for _, tr := range transactions {
		at := AnoteTransaction{}
		trb, err := json.Marshal(tr)
		json.Unmarshal(trb, &at)
		pk, err := crypto.NewPublicKeyFromBase58(at.SenderPublicKey)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}
		addr, err := proto.NewAddressFromPublicKey(55, pk)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}
		if addr.String() == address {
			return at.SenderPublicKey
		}
	}

	return pk
}

type AnoteTransaction struct {
	SenderPublicKey string `json:"senderPublicKey"`
}

func getCallerInfo() (info string) {

	// pc, file, lineNo, ok := runtime.Caller(2)
	_, file, lineNo, ok := runtime.Caller(2)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	// funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // The Base function returns the last element of the path
	return fmt.Sprintf("%s:%d: ", fileName, lineNo)
}

func logTelegram(message string) {
	message = "anote-master:" + getCallerInfo() + url.PathEscape(url.QueryEscape(message))

	_, err := http.Get(fmt.Sprintf("http://localhost:5002/log/%s", message))
	if err != nil {
		log.Println(err)
	}
}
