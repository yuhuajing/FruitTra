package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/ethconn"
)

var (
	EthServer = "https://eth-sepolia.g.alchemy.com/v2/t39LUOfkEMNMz_uxQk2gkwK3kJ1HwDZF"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0xD088960Fc5aB18EA5d7fde419E5e8Aaf26bE0ecC"
)
var (
	Mongodburl              = "mongodb://clay:password@127.0.0.1:27017" // "mongodb://clay:password@127.0.0.1:27017"
	Dbname                  = "fruittra"
	DbcollectionUserInfo    = "userinfo"
	DbcollectionEcoInfo     = "prodinfo"
	DbcollectionTeaInfo     = "teainfo"
	DbcollectionProcessInfo = "processinfo"
	DbcollectionProdInfo    = "prodinfo"
	DbcollectionStoreInfo   = "storeinfo"
	DbcollectionLogisInfo   = "logisinfo"

	DbcollectionTousu = "tousu"
	Mongoclient       *mongo.Client
)
