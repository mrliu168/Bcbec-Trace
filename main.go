package main

import (
	"os"
	"fmt"
	"time"
	"github.com/Bcbec-Trace/sdkInit"
	"github.com/Bcbec-Trace/service"
	"encoding/json"
	"github.com/Bcbec-Trace/web/controller"
	"github.com/Bcbec-Trace/web"
)

const (
	configFile = "config.yaml"
	initialized = false
	ComCC = "comcc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/Bcbec-Trace/fixtures/artifacts/chainhero.channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.hf.chainhero.io",

		ChaincodeID: "ComCC",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/Bcbec-Trace/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID:ComCC,
		Client:channelClient,
	}

	coms := []service.Commodity{
		service.Commodity{Type:"中国邮政EMS", Primarykey:"GSI80000A113", Name:"出口普洱茶", Des:"从地里采摘", Specification:"1500kg", Source:"云南普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/01.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"中国邮政EMS", Primarykey:"GSI80000A114", Name:"出口三七", Des:"从地里采摘", Specification:"2500kg", Source:"云南昆明", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/02.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"国际快运", Primarykey:"GSI80000A115", Name:"进口龙虾", Des:"从海里打捞", Specification:"1000kg", Source:"缅甸", Machining:"淡水养殖",Remarks:"无",Principal:"peeterchen",PhoneNumber:"123456789", Photo:"/static/photo/03.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"缅甸", Vendor:"缅甸某某厂", PlaceOfProduction:"缅甸", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"德邦物流", Primarykey:"GSI80000A116", Name:"进口三文鱼", Des:"从海里打捞", Specification:"2000kg", Source:"缅甸", Machining:"淡水养殖",Remarks:"无",Principal:"Aaron",PhoneNumber:"123456789", Photo:"/static/photo/04.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"泰国", Vendor:"泰国某厂", PlaceOfProduction:"曼谷", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
	}

	for _,v := range coms {
		msg, err := serviceSetup.SaveCom(v)
		if err != nil {
			fmt.Println(err.Error())
		}else {
			fmt.Println("信息发布成功, 交易编号为: " + msg)
		}
	}

	result, err := serviceSetup.FindComInfoByEntityID("GSI80000A113")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var com service.Commodity
		json.Unmarshal(result, &com)
	}

	coms= []service.Commodity{
		service.Commodity{Type:"国际快运", Primarykey:"GS116000A119", Name:"进口水果", Des:"从地里采摘", Specification:"1500kg", Source:"泰国", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/05.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"泰国某某厂", PlaceOfProduction:"泰国", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"国际快运", Primarykey:"GS116000A120", Name:"进口动植物油脂", Des:"生产加工", Specification:"2500kg", Source:"越南", Machining:"人工加工",Remarks:"无",Principal:"Aaron",PhoneNumber:"123456789", Photo:"/static/photo/06.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"越南", Vendor:"越南某某厂", PlaceOfProduction:"越南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"国际快运", Primarykey:"GS116000A121", Name:"进口塑料", Des:"生产加工", Specification:"4500kg", Source:"缅甸", Machining:"人工加工",Remarks:"无",Principal:"Aaron",PhoneNumber:"123456789", Photo:"/static/photo/07.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"泰国", Vendor:"泰国某某厂", PlaceOfProduction:"泰国", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"国际快运", Primarykey:"GS116000A122", Name:"进口棕榈油", Des:"生产加工", Specification:"3500kg", Source:"马来西亚", Machining:"人工加工",Remarks:"无",Principal:"Alice",PhoneNumber:"123456789", Photo:"/static/photo/08.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"马来西亚", Vendor:"马来西亚某某厂", PlaceOfProduction:"马来西亚", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
	}

	for _,v := range coms {
		msg, err := serviceSetup.ModifyCom(v)
		if err != nil {
			fmt.Println(err.Error())
		}else {
			fmt.Println("信息操作成功, 交易编号为: " + msg)
		}
	}

	result, err = serviceSetup.FindComInfoByEntityID("GSI80000A113")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var com service.Commodity
		json.Unmarshal(result, &com)
		fmt.Println(com)
	}

	//===========================================//
result,err=serviceSetup.FindComByCertNoAndName("GSI80000A113","出口普洱茶")
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		var com service.Commodity
		json.Unmarshal(result,&com)
		fmt.Println("根据GSI溯源码以及产品名称查询成功")
		fmt.Println(com)
	}
	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
