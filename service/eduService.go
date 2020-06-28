
package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"encoding/json"
	"fmt"
)

func (t *ServiceSetup) SaveCom(com Commodity) (string, error) {

	eventID := "eventAddCom"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(com)
	if err != nil {
		return "", fmt.Errorf("指定的com对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addCom", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}


func (t *ServiceSetup) FindComInfoByEntityID(entityID string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryComInfoByEntityID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindComByCertNoAndName(primarykey, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryComByCertNoAndName", Args: [][]byte{[]byte(primarykey ), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) ModifyCom(com Commodity) (string, error) {

	eventID := "eventModifyCom"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(com)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateCom", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) DelCom(entityID string) (string, error) {

	eventID := "eventDelEdu"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "delCom", Args: [][]byte{[]byte(entityID), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}

