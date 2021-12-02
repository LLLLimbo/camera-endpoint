package main

import (
	"camera-endpoint/conf"
	"camera-endpoint/db"
	"camera-endpoint/server"
	"camera-endpoint/util"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

const IotHubConfigureEndPoint = "http://localhost:17011/device/deviceInstance/configure"
const Storage = "E:\\data\\badger"
const sn_ = "78E4A0164CCE064A4489013E66FB7A7D"
const device_id = "f6c3b854cb454b74840e5f5facd1b5b6"

type ReqShadow struct {
	Act      string `json:"act"`
	DeviceId string `json:"deviceId"`
}

func main() {
	var configureEndPoint = flag.String("s", IotHubConfigureEndPoint, "configuration end point")
	//var sn = flag.String("sn", sn_, "serial number")
	var deviceId = flag.String("d", device_id, "device id")
	var storage = flag.String("b", Storage, "badger data directory")
	var natsAddr = flag.String("n", nats.DefaultURL, "nats address")

	conf.ConnectBadger(storage)
	db.SetProperty([]byte("deviceId"), []byte(*deviceId))

	StartUpConfiguration(deviceId, configureEndPoint)

	//ConnectNats nats server and subscribe topics
	conf.ConnectNats(natsAddr)
	/*err := conf.ConnectNats()
	if err != nil {
		return
	} else {
		nc := conf.NatsCon()
		req := ReqShadow{
			Act:      "get",
			DeviceId: *deviceId,
		}
		data, err := json.Marshal(req)
		util.LogError(err)

		msg := nats.Msg{
			Subject: "u.sdw.projectKey.productKey.deviceKey",
			Header:  nats.Header{"act": {"get"}, "deviceId": {*deviceId}},
			Data:    data,
		}
		err = nc.PublishMsg(&msg)
		util.LogError(err)
	}*/

	server.SubCloudMsg()

	r := gin.Default()

	v1 := r.Group("/ParkAPI")
	server.HeartBeat(v1)
	server.RecResultPunishment(v1)
	server.SetSynchronizeParking(v1)
	server.SetAdjustBrightness(v1)

	_ = r.Run("0.0.0.0:9030")
}

type DeviceConfRes struct {
	DeviceInstance DeviceInstance `json:"deviceInstance"`
}

type Metadata struct {
}

type DeviceProperty struct {
}

type DeviceInstance struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	ProductKey  string `json:"productKey"`
	ProjectId   string `json:"projectId"`
	ProjectKey  string `json:"projectKey"`
	PositionId  string `json:"positionId"`
	GeoPath     string `json:"geoPath"`
	State       string `json:"state"`
	Sn          string `json:"sn"`
	Status      string `json:"status"`
	CreateTime  string `json:"createTime"`
	ModifyTime  string `json:"modifyTime"`
}

func StartUpConfiguration(deviceId, configureEndPoint *string) {
	client := resty.New()
	device := DeviceInstance{Id: *deviceId, Sn: *deviceId}
	res, err := client.R().SetBody(&device).Post(*configureEndPoint)
	util.LogError(err)
	err = json.Unmarshal(res.Body(), &device)
	util.LogError(err)
	log.Info().Msgf("Received configuration %s.", res.String())
}
