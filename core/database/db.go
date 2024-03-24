package database

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/common/config"
	"main/common/dbconn"
	"main/common/tabletypes"
	"main/core/sendtx"
	"reflect"
	"strings"
)

func InsertUserInfo(username, password, email, phone string, identity tabletypes.Identity) error {
	username = strings.ToLower(username)
	email = strings.ToLower(email)
	filter := bson.M{"username": username}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("InsertUserInfo:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.UserInfo{
			Id:       utils.UUIDv4(),
			Username: username,
			Password: password,
			Email:    email,
			Phone:    phone,
			Identity: identity,
			Approved: 0,
		}
		err := InsertDocument(config.DbcollectionUserInfo, res)
		if err != nil {
			return fmt.Errorf("InsertUserInfo:err in inserting NFTData")
		}
		return nil
	}
	return fmt.Errorf("Already Registed Username: %s", email)
}

func InsertInfoTea(planet string, location string, grower string, species string, quarantine string, packages string, pick string, picktime string, picker string) (error, string) {
	rid := uuid.New().String()
	txhash := sendtx.AddOrupdateTeaData(rid, planet, location, grower, species, quarantine, packages, pick, picktime, picker)
	var res = tabletypes.TeaResData{
		Id:         rid,
		Planet:     planet,
		Location:   location,
		Grower:     grower,
		Species:    species,
		Quarantine: quarantine,
		Packages:   packages,
		Pick:       pick,
		PickTime:   picktime,
		Picker:     picker,
		Hash:       txhash,
	}
	err := InsertDocument(config.DbcollectionTeaInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash
}

func InsertInfoProd(id string, processa string, processtime string, processer string, company string) (error, string) {
	txhash := sendtx.AddOrUpdateProdData(id, processa, processtime, processer, company)
	var res = tabletypes.ProductionInfo{
		Id:             id,
		Production:     processa,
		ProductionTime: processtime,
		Producer:       processer,
		Company:        company,
		Hash:           txhash,
	}
	err := InsertDocument(config.DbcollectionTeaInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash
}

func InsertInfoProcess(id string, processa string, processtime string, processer string, company string) (error, string) {
	txhash := sendtx.AddOrupdateProcessData(id, processa, processtime, processer, company)
	var res = tabletypes.ProcessInfo{
		Id:          id,
		Process:     processa,
		ProcessTime: processtime,
		Processer:   processer,
		Company:     company,
		Hash:        txhash,
	}
	err := InsertDocument(config.DbcollectionTeaInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash
}

func InsertInfoStore(id string, store string, storetime string, company string) (error, string) {
	txhash := sendtx.AddOrupdateStoreData(id, store, storetime, company)
	var res = tabletypes.StorageInfo{
		Id:        id,
		Store:     store,
		StoreTime: storetime,
		Company:   company,
		Hash:      txhash,
	}
	err := InsertDocument(config.DbcollectionTeaInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash
}

func InsertInfoLogis(id string, path string, logisway string, logistime string, driver string, company string) (error, string) {
	txhash := sendtx.AddOrupdateLogissData(id, path, logisway, logistime, driver, company)
	var res = tabletypes.LogisInfo{
		Id:        id,
		Path:      path,
		LogisWay:  logisway,
		Driver:    driver,
		LogisTime: logistime,
		Company:   company,
		Hash:      txhash,
	}
	err := InsertDocument(config.DbcollectionTeaInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash
}

func QueryUserInfos(username, password string) (error, *tabletypes.UserInfo) {
	username = strings.ToLower(username)
	filter := bson.M{"username": username, "password": password, "approved": 1}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	var resdata *tabletypes.UserInfo
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err), resdata
	}
	if len(idres) == 0 {
		return fmt.Errorf("Non-Registed/Not-Approved user, Please register firstly."), resdata
	}
	resdata = idres[0].(*tabletypes.UserInfo)
	return nil, resdata
}

func QueryAllUserInfo() (error, []*tabletypes.UserInfo) {
	filter := bson.M{}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("QueryAllUserInfo err : %s", err), nil
	}
	if len(idres) != 0 {
		resdata := make([]*tabletypes.UserInfo, 0)
		for _, data := range idres {
			res := data.(*tabletypes.UserInfo)
			resdata = append(resdata, res)
		}
		return nil, resdata
	}
	return nil, nil
}

func QueryTeaChainData() []*tabletypes.TeaResData {
	filter := bson.M{}
	alldata := make([]*tabletypes.TeaResData, 0)
	err, idres := GetDocuments(config.DbcollectionTeaInfo, filter, &tabletypes.TeaResData{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.TeaResData)
			alldata = append(alldata, res)
		}
	}
	return alldata
}
func QueryProdChainData() []*tabletypes.ProductionInfo {
	filter := bson.M{}
	alldata := make([]*tabletypes.ProductionInfo, 0)
	err, idres := GetDocuments(config.DbcollectionProdInfo, filter, &tabletypes.ProductionInfo{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.ProductionInfo)
			alldata = append(alldata, res)
		}
	}
	return alldata
}
func QueryProcessChainData() []*tabletypes.ProcessInfo {
	filter := bson.M{}
	alldata := make([]*tabletypes.ProcessInfo, 0)
	err, idres := GetDocuments(config.DbcollectionProcessInfo, filter, &tabletypes.ProcessInfo{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.ProcessInfo)
			alldata = append(alldata, res)
		}
	}
	return alldata
}
func QueryStoreChainData() []*tabletypes.StorageInfo {
	filter := bson.M{}
	alldata := make([]*tabletypes.StorageInfo, 0)
	err, idres := GetDocuments(config.DbcollectionStoreInfo, filter, &tabletypes.StorageInfo{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.StorageInfo)
			alldata = append(alldata, res)
		}
	}
	return alldata
}
func QueryLogisChainData() []*tabletypes.LogisInfo {
	filter := bson.M{}
	alldata := make([]*tabletypes.LogisInfo, 0)
	err, idres := GetDocuments(config.DbcollectionLogisInfo, filter, &tabletypes.LogisInfo{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.LogisInfo)
			alldata = append(alldata, res)
		}
	}
	return alldata
}

func QueryChainData() []tabletypes.DataRes {
	filter := bson.M{}
	ids := make([]string, 0)
	alldata := make([]tabletypes.DataRes, 0)
	err, idres := GetDocuments(config.DbcollectionTeaInfo, filter, &tabletypes.TeaResData{})
	if err != nil {
		fmt.Println(err)
		return alldata
	}
	if len(idres) != 0 {
		fmt.Println(123)
		for _, data := range idres {
			res := data.(*tabletypes.TeaResData)
			ids = append(ids, res.Id)
		}
	}

	for _, id := range ids {
		tea := QueryTeaByIdChainData(id)
		prod := QueryProdByIdChainData(id)
		process := QueryProcessByIdChainData(id)
		store := QueryStoreByIdChainData(id)
		logis := QueryLogisByIdChainData(id)
		alldata = append(alldata, tabletypes.DataRes{tea, prod, process, store, logis})
	}
	fmt.Println(alldata)
	return alldata
}

func QueryChainDataById(id string) []tabletypes.DataRes {
	alldata := make([]tabletypes.DataRes, 0)
	tea := QueryTeaByIdChainData(id)
	prod := QueryProdByIdChainData(id)
	process := QueryProcessByIdChainData(id)
	store := QueryStoreByIdChainData(id)
	logis := QueryLogisByIdChainData(id)
	alldata = append(alldata, tabletypes.DataRes{tea, prod, process, store, logis})
	return alldata
}

func QueryTeaByIdChainData(id string) *tabletypes.TeaResData {
	filter := bson.M{"id": id}
	opts := options.Find().SetLimit(1)
	var res *tabletypes.TeaResData
	err, idres := GetDocuments(config.DbcollectionTeaInfo, filter, &tabletypes.TeaResData{}, opts)
	if err != nil {
		fmt.Println(err)
		return res
	}
	if len(idres) != 0 {
		res = idres[0].(*tabletypes.TeaResData)
	}
	return res
}

func QueryProdByIdChainData(id string) *tabletypes.ProductionInfo {
	filter := bson.M{"id": id}
	opts := options.Find().SetLimit(1)
	var res *tabletypes.ProductionInfo
	err, idres := GetDocuments(config.DbcollectionProdInfo, filter, &tabletypes.ProductionInfo{}, opts)
	if err != nil {
		fmt.Println(err)
		return res
	}
	if len(idres) != 0 {
		res = idres[0].(*tabletypes.ProductionInfo)
	}
	return res
}
func QueryProcessByIdChainData(id string) *tabletypes.ProcessInfo {
	filter := bson.M{"id": id}
	opts := options.Find().SetLimit(1)
	var res *tabletypes.ProcessInfo
	err, idres := GetDocuments(config.DbcollectionProcessInfo, filter, &tabletypes.ProcessInfo{}, opts)
	if err != nil {
		fmt.Println(err)
		return res
	}
	if len(idres) != 0 {
		res = idres[0].(*tabletypes.ProcessInfo)
	}
	return res
}

func QueryStoreByIdChainData(id string) *tabletypes.StorageInfo {
	filter := bson.M{"id": id}
	opts := options.Find().SetLimit(1)
	var res *tabletypes.StorageInfo
	err, idres := GetDocuments(config.DbcollectionStoreInfo, filter, &tabletypes.StorageInfo{}, opts)
	if err != nil {
		fmt.Println(err)
		return res
	}
	if len(idres) != 0 {
		res = idres[0].(*tabletypes.StorageInfo)
	}
	return res
}

func QueryLogisByIdChainData(id string) *tabletypes.LogisInfo {
	filter := bson.M{"id": id}
	opts := options.Find().SetLimit(1)
	var res *tabletypes.LogisInfo
	err, idres := GetDocuments(config.DbcollectionLogisInfo, filter, &tabletypes.LogisInfo{}, opts)
	if err != nil {
		fmt.Println(err)
		return res
	}
	if len(idres) != 0 {
		res = idres[0].(*tabletypes.LogisInfo)
	}
	return res
}

func QueryUserInfoByName(username string) (error, *tabletypes.UserInfo) {
	filter := bson.M{"username": username}
	var res *tabletypes.UserInfo
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfoByName err :%s", err), res
	}
	if len(idres) != 0 {
		res := idres[0].(*tabletypes.UserInfo)
		return nil, res
	}
	return fmt.Errorf("QueryUserInfoByName err :%s", err), res
}

func ApproveUserInfo(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		update := bson.M{"$set": bson.M{"approved": 1}}
		err := UpdateDocument(config.DbcollectionUserInfo, filter, update)
		if err != nil {
			return fmt.Errorf("ApproveUserInfo err")
		}
		return nil
	}
	return fmt.Errorf("Approve non-exist user.")
}

func DeleteChaindata(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionTeaInfo, filter, &tabletypes.TeaResData{})
	if err != nil {
		return fmt.Errorf("DeleteChaindata: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionTeaInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DbcollectionEcoInfo: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist eco.")
}

func DeleteUserInfo(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionUserInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func DeleteTea(id string) error {
	filter := bson.M{"_id": id}
	err, idres := GetDocuments(config.DbcollectionTeaInfo, filter, &tabletypes.TeaResData{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionTeaInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func DeleteLogis(id string) error {
	filter := bson.M{"_id": id}
	err, idres := GetDocuments(config.DbcollectionLogisInfo, filter, &tabletypes.LogisInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionLogisInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func DeleteStore(id string) error {
	filter := bson.M{"_id": id}
	err, idres := GetDocuments(config.DbcollectionStoreInfo, filter, &tabletypes.StorageInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionStoreInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func DeleteProcess(id string) error {
	filter := bson.M{"_id": id}
	err, idres := GetDocuments(config.DbcollectionProcessInfo, filter, &tabletypes.ProcessInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionProcessInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func DeleteProd(id string) error {
	filter := bson.M{"_id": id}
	err, idres := GetDocuments(config.DbcollectionProdInfo, filter, &tabletypes.ProductionInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionProdInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func UpdateAdminUserInfo(username, phone, email string) error {
	//filter := bson.M{"username": strings.ToLower(username), "approved": 1}
	filter := bson.M{"username": strings.ToLower(username)}
	update := bson.M{"$set": bson.M{"email": email, "phone": phone}}
	err := UpdateDocument(config.DbcollectionUserInfo, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateUserInfo err")
	}
	return nil

}

func InsertDocument(collectionname string, data interface{}) error {
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	return nil
}

func GetDocuments(collectionname string, filter bson.M, result interface{}, opts ...*options.FindOptions) (error, []interface{}) {
	collection := dbconn.GetCollection(collectionname)
	cur, err := collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to get documents: %v", err)), nil
	}
	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		// Create a new instance of the result type for each document
		res := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		err := cur.Decode(res)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to decode document: %v", err)), nil
		}
		results = append(results, res)
	}
	if err := cur.Err(); err != nil {
		return fmt.Errorf(fmt.Sprintf("cursor error: %v", err)), nil
	}
	return nil, results
}
func UpdateDocument(collectionname string, filter bson.M, update bson.M) error {
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}
