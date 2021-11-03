package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

func main() {
	keys, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	password := "*********************************"
	key, err := keystore.DecryptKey(keys, password)

	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)

	if err != nil {
		panic(err)
	}
	fmt.Printf("\naddress :\t %x", key.Address)
	fmt.Printf("\n*** 🎉 Key : \t %x", hexutil.Encode(privateKeyBytes)[2:])
}
