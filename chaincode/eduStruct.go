package main



 /**
 产品唯一ID(溯源编号)：

 事件类型：

 简介：

 产品名称：

 产品规格：

 产品来源：

 加工方式：

 照片：

 备注信息：

 负责人：

 联系方式：

 录入时间：
  */
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
	PhoneNumber    string   `json:"PhoneNumber"`
	Photo	       string	`json:"Photo"`	            // 照片

	ShelfLife      string   `json:"shelfLife"`   //保质期
	StorageMethod  string   `json:"storageMethod"`  //储藏方式
	Brand          string   `json:"brand"`  //品牌
	Vendor         string   `json:"vendor"`  //生产厂商
	PlaceOfProduction   string   `json:"placeOfProduction"`  //生产地
	ExecutiveStandard   string   `json:"executiveStandard"`  //执行标准

 	Historys	   []HistoryItem	// 当前com的历史记录
	Time           string  `json:"Time"`   //时间
}


type HistoryItem struct {
	TxId	string
	Commodity   Commodity
}
