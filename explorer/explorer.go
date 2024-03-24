package explorer

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	html "github.com/gofiber/template/html/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"log"
	"main/common/tabletypes"
	"main/core/database"
	txdata2 "main/core/txdata"
	"time"
)

func Explorer() {
	// 路由分之
	app := fiber.New(fiber.Config{
		Views: html.New("./explorer/views", ".html"),
	})
	//静态界面
	app.Static("/", "./explorer/public")
	//默认进入 index.html界面
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})
	app.Post("/login", deallogin)           //登录分支
	app.Post("/registerUser", registerUser) //注册分支
	app.Post("/adminfunc", manageuser)      //管理员界面
	app.Post("/produserfunc", produserfunc) //
	app.Post("/planetfunc", produserfunc)   //
	app.Post("/pickfunc", produserfunc)     //
	app.Post("/processfunc", produserfunc)  //
	app.Get("/approveuser", approveuser)
	app.Post("/giveapprove", giveapprove)
	app.Post("/removeapprove", removeapprove)
	app.Post("/updateAdminUser", updateAdminUser)

	app.Post("/addtea", addtea)
	app.Post("/addprod", addprod)
	app.Post("/addprocess", addprocess)
	app.Post("/addstorage", addstorage)
	app.Post("/addlogis", addlogis)

	app.Get("/getchaindata", getchaindata)
	app.Get("/getchaindatauser", getchaindatauser)
	app.Post("/delchaindata", delchaindata)
	app.Get("/gettxbyhash", gettxbyhash)

	app.Get("/checkdata", checkdata)
	app.Get("/checktx", checktx)

	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	//}))

	log.Fatal(app.Listen(":3005"))
}

type Tousu struct {
	User  string `json:"user"`
	Tousu string `json:"tousu"`
}

func getchaindatauser(c *fiber.Ctx) error {
	resdata := database.QueryChainData()
	return c.Render("chaindatauser", fiber.Map{
		"Data": resdata,
	})
}

func checkdata(c *fiber.Ctx) error {
	resdata := database.QueryChainData()
	return c.Render("chaindatacheck", fiber.Map{
		"Data": resdata,
	})
}
func getchaindata(c *fiber.Ctx) error {
	resdata := database.QueryChainData()
	return c.Render("chaindata", fiber.Map{
		"Data": resdata,
	})
}

func approveuser(c *fiber.Ctx) error {
	_, resdata := database.QueryAllUserInfo()
	return c.Render("approveuser", fiber.Map{
		"Data": resdata,
	})
}

type ecodatas struct {
	Traceid    string `json:"traceid"`
	Batchid    string `json:"batchid"`
	Prodinfo   string `json:"prodinfo"`
	Logisinfo  string `json:"logisinfo"`
	Soreinfo   string `json:"soreinfo"`
	Salestring string `json:"salestring"`
	//Logisinfo  string `json:"logisinfo"`

}

func checktx(c *fiber.Ctx) error {
	//id := c.Query("id")
	//err, traceid, batchid, prodinfo, logisinfo, soreinfo, salestring := sendtx.ReadDataByID(id)
	//if err != nil {
	//	return c.SendStatus(400)
	//}
	//return c.Status(200).JSON(ecodatas{
	//	Traceid:    traceid,
	//	Batchid:    batchid,
	//	Prodinfo:   prodinfo,
	//	Logisinfo:  logisinfo,
	//	Soreinfo:   soreinfo,
	//	Salestring: salestring,
	//})
	return c.SendStatus(200)
}

type txdata struct {
	Chainid   uint64    `json:"chainid"`
	Blockhash string    `json:"blockhash"`
	Blocknum  uint64    `json:"blocknum"`
	Txtime    time.Time `json:"txtime"`
	Contract  string    `json:"contract"`
	Gas       uint64    `json:"gas"`
	Confirmed uint64    `json:"confirmed"`
}

func gettxbyhash(c *fiber.Ctx) error {
	hash := c.Query("hash")
	err, chainid, blockhash, blocknum, txtime, contract, gas, blocks := txdata2.Gettxbyhash(hash)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(txdata{
		Chainid:   chainid,
		Blockhash: blockhash,
		Blocknum:  blocknum,
		Txtime:    txtime,
		Contract:  contract,
		Gas:       gas,
		Confirmed: blocks,
	})
}

func produserfunc(c *fiber.Ctx) error {
	payload := &tabletypes.Info{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, resdata := database.QueryUserInfoByName(payload.Username)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return c.Status(200).JSON(tabletypes.UserInfo{
		Username: resdata.Username,
		Id:       resdata.Id,
		Email:    resdata.Email,
		Phone:    resdata.Phone,
		Identity: resdata.Identity,
		Approved: resdata.Approved,
	})
}
func manageuser(c *fiber.Ctx) error {
	payload := &tabletypes.Info{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, resdata := database.QueryUserInfoByName(payload.Username)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return c.Status(200).JSON(tabletypes.UserInfo{
		Username: resdata.Username,
		Id:       resdata.Id,
		Email:    resdata.Email,
		Phone:    resdata.Phone,
		Identity: resdata.Identity,
		Approved: resdata.Approved,
	})
}

func giveapprove(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.ApproveUserInfo(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

func delchaindata(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.DeleteChaindata(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}
func removeapprove(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.DeleteUserInfo(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

type UserRes struct {
	Token string `json:"token"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
}

func deallogin(c *fiber.Ctx) error {
	payload := &tabletypes.UserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//fmt.Println(payload)
	err, userinfo := database.QueryUserInfos(payload.Username, payload.Password)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	myClaims := &jwt.MapClaims{
		"name": "John Doe",
		"uid":  utils.UUIDv4(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	c.Locals("user", myClaims)
	fmt.Println(t)
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Token:   t,
		Data:    payload.Username,
		Type:    userinfo.Identity,
	})
}

type DataResponse struct {
	Error   string              `json:"error"`
	Success bool                `json:"success" default:"false"`
	Data    string              `json:"data"`
	Token   string              `json:"token"`
	Type    tabletypes.Identity `json:"type"`
}

func registerUser(c *fiber.Ctx) error {
	//fmt.Println(string(c.Body()))
	payload := &tabletypes.UserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err := database.InsertUserInfo(payload.Username, payload.Password, payload.Email, payload.Phone, payload.Identity)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "Duplicated",
		})
	}
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Token:   "",
		Data:    "Register successfully!!",
		Type:    "",
	})
}
func updateAdminUser(c *fiber.Ctx) error {
	payload := &tabletypes.NewAdminUserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	err := database.UpdateAdminUserInfo(payload.Username, payload.Phone, payload.Email)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

func addtea(c *fiber.Ctx) error {
	payload := &tabletypes.TeaResData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertInfoTea(payload.Planet, payload.Location, payload.Grower, payload.Species, payload.Quarantine, payload.Packages, payload.Pick, payload.PickTime, payload.Picker)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}

func addprocess(c *fiber.Ctx) error {
	payload := &tabletypes.ProcessInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertInfoProcess(payload.Id, payload.Process, payload.ProcessTime, payload.Processer, payload.Company)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}

func addprod(c *fiber.Ctx) error {
	payload := &tabletypes.ProductionInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertInfoProd(payload.Id, payload.Production, payload.ProductionTime, payload.Producer, payload.Company)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}

func addstorage(c *fiber.Ctx) error {
	payload := &tabletypes.StorageInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertInfoStore(payload.Id, payload.Store, payload.StoreTime, payload.Company)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}

func addlogis(c *fiber.Ctx) error {
	payload := &tabletypes.LogisInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertInfoLogis(payload.Id, payload.Path, payload.LogisWay, payload.LogisTime, payload.Driver, payload.Company)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}
