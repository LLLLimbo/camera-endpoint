package server

import (
	"camera-endpoint/util"
	json2 "encoding/json"
	"github.com/go-resty/resty/v2"
)

type QGDeviceRes struct {
	Info       string `json:"info"`
	ResultCode int32  `json:"resultCode"`
	Data       string `json:"data"`
}

type FuncInputSetSyncParking struct {
	Id      string `json:"id"`
	ReqTime int64  `json:"reqTime"`
	Type    string `json:"type"`
	Act     string `json:"act"`
	Input   struct {
		Channelout   int32 `json:"channelout"`
		Parkingspace int32 `json:"parkingspace"`
	} `json:"input"`
}

type FuncInputSetBrt struct {
	Id      string `json:"id"`
	ReqTime int64  `json:"reqTime"`
	Type    string `json:"type"`
	Act     string `json:"act"`
	Input   struct {
		Channelout int32 `json:"channelout"`
		Brightness int32 `json:"brightness"`
	} `json:"input"`
}

type ReqSetSyncParking struct {
	Key          string `json:"key"`
	Channelout   int32  `json:"channelout"`
	Parkingspace int32  `json:"parkingspace"`
}

type ReqSetBrt struct {
	Key        string `json:"key"`
	Channelout int32  `json:"channelout"`
	Brightness int32  `json:"brightness"`
}

func SetSynchronizeParking_(input []byte) QGDeviceRes {
	var funcInput = FuncInputSetSyncParking{}
	err := json2.Unmarshal(input, &funcInput)
	util.LogError(err)

	req := ReqSetSyncParking{}
	req.Key = funcInput.Id
	req.Channelout = funcInput.Input.Channelout
	req.Parkingspace = funcInput.Input.Parkingspace

	rawreq, err := json2.Marshal(req)
	util.LogError(err)
	client := resty.New()
	var url = Host + "/Home/SetSynchronizeParking"
	response, err := util.SimpleFormPost(client, rawreq, url)
	util.LogError(err)

	body := response.Body()
	var jsonRes = QGDeviceRes{}
	err = json2.Unmarshal(body, &jsonRes)
	util.LogError(err)

	return jsonRes
}

func SetAdjustBrightness_(input []byte) QGDeviceRes {
	var funcInput = FuncInputSetBrt{}
	err := json2.Unmarshal(input, &funcInput)
	util.LogError(err)

	req := ReqSetBrt{}
	req.Key = funcInput.Id
	req.Channelout = funcInput.Input.Channelout
	req.Brightness = funcInput.Input.Brightness

	rawreq, err := json2.Marshal(req)
	util.LogError(err)
	client := resty.New()
	var url = Host + "/Home/SetAdjustBrightness"
	response, err := util.SimpleFormPost(client, rawreq, url)
	util.LogError(err)

	body := response.Body()
	var jsonRes = QGDeviceRes{}
	err = json2.Unmarshal(body, &jsonRes)
	util.LogError(err)

	return jsonRes
}
