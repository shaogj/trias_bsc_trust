package trustbscparlia

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"strconv"

	"net/http"
)

//1108add for generate
func CreateKey() (privs, addrs string) {
	//创建私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return priv, address
}

	/*
	getPriv,getAddr :=CreateKey()
	fmt.Println("after CreateKey() get getPriv is:%s,getAddr is:%s,len(getAddr) is:%d", getPriv,getAddr,len(getAddr))
	return 0, nil,nil
	*/