package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/common/tabletypes"
	"main/trace"
	"math/big"
)

func gentx(gaslimit uint64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("05e4be45dfd1bbed7c0b541630ec25fe3e94c7688e83566e6137391364453fc2")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrupdateTeaData(id string, planet string, location string, grower string, species string, quarantine string, packages string, pick string, picktime string, picker string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(1000000)
	tx, err := instance.AddOrupdateTeaData(auth, id, planet, location, grower, species, quarantine, packages, pick, picktime, picker)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
}
func AddOrUpdateProdData(id string, processa string, processtime string, processer string, company string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(1000000)
	tx, err := instance.AddOrupdateProdData(auth, id, processa, processtime, processer, company)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
}
func AddOrupdateProcessData(id string, processa string, processtime string, processer string, company string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(1000000)
	tx, err := instance.AddOrupdateProcessData(auth, id, processa, processtime, processer, company)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
}

func AddOrupdateLogissData(id string, path string, logisway string, logistime string, driver string, company string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(1000000)
	tx, err := instance.AddOrupdateLogissData(auth, id, path, logisway, logistime, driver, company)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
}
func AddOrupdateStoreData(id string, store string, storetime string, company string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(1000000)
	tx, err := instance.AddOrupdateStoreData(auth, id, store, storetime, company)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
}

func ReadDataByID(id string) tabletypes.Res {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	tx, err := instance.Tracedata(nil, id)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tabletypes.Res{
		tx.Tea,
		tx.Process,
		tx.Production,
		tx.Store,
		tx.Logis,
	}
}
