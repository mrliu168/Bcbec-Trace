package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Commodity struct {
	ObjectType	   string	`json:"docType"`
	Type           string	`json:"type"`  //事件类型
	Primarykey     string   `json:"primarykey"`  //主键，唯一Id
	Name	       string	`json:"name"`
	Des            string   `json:"des"`  //描述
	Specification  string   `json:"specification"`  //规格
	Source         string   `json:"source"`  //产品来源
	Machining      string   `json:"machining"`    //加工
	Remarks        string   `json:"remarks"`    //备注信息
	Principal      string   `json:"principal"`  //负责人
	PhoneNumber    string   `json:"phoneNumber"`
	Photo	       string	`json:"Photo"`	 // 照片

	ShelfLife      string   `json:"shelfLife"`   //保质期
	StorageMethod  string   `json:"storageMethod"`  //储藏方式
	Brand          string   `json:"brand"`  //品牌
	Vendor         string   `json:"vendor"`  //生产厂商
	PlaceOfProduction   string   `json:"placeOfProduction"`  //生产地
	ExecutiveStandard   string   `json:"executiveStandard"`  //执行标准

	Historys       []HistoryItem	// 当前com的历史记录
	Time           string  `json:"Time"`   //时间
}

type HistoryItem struct {
	TxId	string
	Commodity   Commodity
}

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
	BlockNumber uint64
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string)(uint64, error) {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)

		return ccEvent.BlockNumber, nil
	case <-time.After(time.Second * 20):
		return 0, fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return 0, nil
}
