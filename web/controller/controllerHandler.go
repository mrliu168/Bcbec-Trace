package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Bcbec-Trace/service"
	"time"
	"strconv"
)

var cuser User

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}
func (app *Application) Company(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "company.html", nil)
}
func (app *Application) Message(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "message.html", nil)
}
func (app *Application) News(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "news.html", nil)
}
func (app *Application) Disclaime(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "disclaimer.html", nil)
}
func (app *Application) Privacy(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "privacy.html", nil)
}
func (app *Application) Job(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "job.html", nil)
}
func (app *Application) Contact(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "contact.html", nil)
}
func (app *Application) Link(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "links.html", nil)
}
func (app *Application) Index(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		Flag bool
		Num string
	}{
		Flag:false,
		Num:string(strconv.Itoa(int(app.Setup.BlockNumber))),
	}
	ShowView(w, r, "index.html", data)
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser:cuser,
	}
	ShowView(w, r, "help.html", data)
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
		Num string
	}{
		CurrentUser:cuser,
		Flag:false,
		Num:string(strconv.Itoa(int(app.Setup.BlockNumber))),
	}

	if flag {
		// 登录成功
		ShowView(w, r, "index.html", data)
	}else{
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}
func (app *Application) Register(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "register.html", nil)
}
func (app *Application) AddEduShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "addEdu.html", data)
}


func (app *Application) AddEdu(w http.ResponseWriter, r *http.Request)  {

	com := service.Commodity{
		Type:r.FormValue("docType"),
		Primarykey:r.FormValue("primarykey"),
		Name:r.FormValue("name"),
		Des:r.FormValue("des"),
		Specification:r.FormValue("specification"),
		Source:r.FormValue("source"),
		Machining:r.FormValue("machining"),
		Remarks:r.FormValue("remarks"),
		Principal:r.FormValue("principal"),
		PhoneNumber:r.FormValue("phoneNumber"),
		Photo:r.FormValue("photo"),

		ShelfLife:r.FormValue("shelfLife"),
		StorageMethod:r.FormValue("storageMethod"),
		Brand:r.FormValue("brand"),
		Vendor:r.FormValue("vendor"),
		PlaceOfProduction:r.FormValue("placeOfProduction"),
		ExecutiveStandard:r.FormValue("executiveStandard"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}

	transactionID,err:=app.Setup.SaveCom(com)
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}
	if err!=nil{
		data.Msg=err.Error()
	}else{
		data.Msg="产品数据信息添加成功，区块交易hash为："+transactionID
	}
	ShowView(w,r,"addEdu.html",data)

	r.Form.Set("entityID", com.Primarykey)
	app.FindByID(w, r)
}

func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query.html", data)
}

// 根据物流编号与收件人查询信息
func (app *Application) FindComByCertNoAnd(w http.ResponseWriter, r *http.Request)  {
	primarykey:= r.FormValue("primarykey")
	name := r.FormValue("name")
	result, err:= app.Setup.FindComByCertNoAndName(primarykey, name)

	var com = service.Commodity{}
	json.Unmarshal(result, &com)
	fmt.Println("根据溯源编号与收件人名查询信息成功：")
	fmt.Println(com)
	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Com:com,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:false,
	}
	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}else{
		data.Msg="信息查询成功"
	}
	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query2.html", data)
}

func (app *Application) FindByID(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindComInfoByEntityID(entityID)
	var com = service.Commodity{}
	json.Unmarshal(result, &com)

	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Com:com,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindComInfoByEntityID(entityID)
	var com = service.Commodity{}
	json.Unmarshal(result, &com)

	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
	}{
		Com:com,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)
}
func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {
	com := service.Commodity{
		Type:r.FormValue("docType"),
		Primarykey:r.FormValue("primarykey"),
		Name:r.FormValue("name"),
		Des:r.FormValue("des"),
		Specification:r.FormValue("specification"),
		Source:r.FormValue("source"),
		Machining:r.FormValue("machining"),
		Remarks:r.FormValue("remarks"),
		Principal:r.FormValue("principal"),
		PhoneNumber:r.FormValue("phoneNumber"),
		Photo:r.FormValue("photo"),

		ShelfLife:r.FormValue("shelfLife"),
		StorageMethod:r.FormValue("storageMethod"),
		Brand:r.FormValue("brand"),
		Vendor:r.FormValue("vendor"),
		PlaceOfProduction:r.FormValue("placeOfProduction"),
		ExecutiveStandard:r.FormValue("executiveStandard"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}

	transactionID,err:=app.Setup.ModifyCom(com)
	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
	}{
		Com:com,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}
	if err!=nil{
		data.Msg=err.Error()
	}else{
		data.Msg="产品数据来源信息更新成功，区块交易Hash为："+transactionID
	}
	ShowView(w,r,"modify.html",data)

	r.Form.Set("entityID", com.Primarykey)
	app.FindByID(w, r)
}
