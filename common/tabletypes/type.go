package tabletypes

import "main/trace"

type Identity string

var (
	User  Identity = "user"
	Admin Identity = "admin"
	Prod  Identity = "prod"
	Logis Identity = "logis"
	Store Identity = "store"
	Sale  Identity = "sale"
)

type UserInfo struct {
	Id       string   `json:"id" gorm:"primary_key"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Identity Identity `json:"identity"`
	Approved int64    `json:"approved"`
}

type Info struct {
	Username string `json:"username"`
}

type TraceInfo struct {
	TraceID       string `json:"traceID"`
	ProdID        string `json:"prodID" `
	ProdBaseInfo  string `json:"prodBaseInfo"`
	ProdTxHash    string `json:"prodTxHash"`
	StoreID       string `json:"storeID" `
	StoreBaseInfo string `json:"storeBaseInfo"`
	StoreTxHash   string `json:"storeTxHash"`
	LogisID       string `json:"logisID" `
	LogisBaseInfo string `json:"logisBaseInfo"`
	LogisTxHash   string `json:"logisTxHash"`
	SaleID        string `json:"saleID" `
	SaleBaseInfo  string `json:"saleBaseInfo"`
	SaleTxHash    string `json:"saleTxHash"`
}

type TeaResData struct {
	Id         string `json:"id"`
	Planet     string `json:"planet"`
	Location   string `json:"location"`
	Grower     string `json:"grower"`
	Species    string `json:"species"`
	Quarantine string `json:"quarantine"`
	Packages   string `json:"packages"`
	Pick       string `json:"pick"`
	PickTime   string `json:"picktime"`
	Picker     string `json:"picker"`
	Hash       string `json:"hash"`
}

type ProductionInfo struct {
	Id             string `json:"id"`
	Production     string `json:"production"`
	ProductionTime string `json:"productiontime"`
	Producer       string `json:"producer"`
	Company        string `json:"company"`
	Hash           string `json:"hash"`
}
type ProcessInfo struct {
	Id          string `json:"id"`
	Process     string `json:"process"`
	ProcessTime string `json:"processtime"`
	Processer   string `json:"processer"`
	Company     string `json:"company"`
	Hash        string `json:"hash"`
}

type StorageInfo struct {
	Id        string `json:"id"`
	Store     string `json:"store"`
	StoreTime string `json:"storetime"`
	Company   string `json:"company"`
	Hash      string `json:"hash"`
}

type LogisInfo struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	LogisWay  string `json:"logisway"`
	Driver    string `json:"driver"`
	LogisTime string `json:"logistime"`
	Company   string `json:"company"`
	Hash      string `json:"hash"`
}

type NewUserInfo struct {
	Username    string `json:"username"`
	Newusername string `json:"newusername"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
type NewAdminUserInfo struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type Res struct {
	trace.TeatraceTeaInfo
	trace.TeatraceprocessInfo
	trace.TeatraceproductionInfo
	trace.TeatracestorageInfo
	trace.TeatracelogisInfo
}

type DataRes struct {
	*TeaResData
	*ProductionInfo
	*ProcessInfo
	*StorageInfo
	*LogisInfo
}
