package server

import (
	"camera-endpoint/conf"
	"camera-endpoint/db"
	"camera-endpoint/util"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"net/http"
)

const Host = "http://192.168.14.188"

type RecEvent struct {
	DeviceId string    `json:"deviceId"`
	Act      string    `json:"act"`
	Output   RecResult `json:"output"`
}

type RecResult struct {
	VehicleLaneKey    string `json:"vehicleLaneKey"`
	Ipaddr            string `json:"ipaddr"`
	License           string `json:"license"`
	ColorType         string `json:"colorType"`
	Type              string `json:"type"`
	Confidence        string `json:"confidence"`
	ScanTime          string `json:"scanTime"`
	ImageFile         string `json:"imageFile"`
	ImageFragmentFile string `json:"imageFragmentFile"`
	TriggerType       string `json:"triggerType"`
}

func HeartBeat(r *gin.RouterGroup) {
	r.POST("/Heartbeat", func(c *gin.Context) {
		type HeartBeatResult struct {
			VehicleLaneKey string `json:"vehicleLaneKey"`
		}
		hbres := HeartBeatResult{}
		err := c.BindJSON(&hbres)
		if err != nil {
			return
		}
		log.Info().Msgf("vehicleLaneKey = %s", hbres.VehicleLaneKey)
		c.JSON(http.StatusOK, gin.H{})

		nc := conf.NatsCon()
		data, _ := json2.Marshal(hbres)
		msg := nats.Msg{
			Subject: "u.eve.s1206.heartBeat",
			Data:    data,
		}
		nc.PublishMsg(&msg)
	})
}

func RecResultPunishment(r *gin.RouterGroup) {
	r.POST("/sendScanCar", func(c *gin.Context) {
		recRes := RecResult{}
		err := c.BindJSON(&recRes)
		if err != nil {
			log.Err(err)
			return
		}
		file1 := util.Base64toFile(recRes.ImageFile)
		file2 := util.Base64toFile(recRes.ImageFragmentFile)
		recRes.ImageFile = file1
		recRes.ImageFragmentFile = file2
		log.Info().Msgf("RecResult = %v", recRes)

		event := RecEvent{
			DeviceId: string(db.GetDeviceId()),
			Act:      "report",
			Output:   recRes,
		}

		nc := conf.NatsCon()
		data, _ := json2.Marshal(event)
		msg := nats.Msg{
			Subject: "u.eve.QGS1206A.carIdentify",
			Data:    data,
		}
		nc.PublishMsg(&msg)
	})
}

func SetSynchronizeParking(r *gin.RouterGroup) {
	r.POST("/SetSynchronizeParking", func(c *gin.Context) {
		input, _ := c.GetRawData()
		jsonRes := SetSynchronizeParking_(input)
		c.JSON(200, jsonRes)
	})
}

func SetAdjustBrightness(r *gin.RouterGroup) {
	r.POST("/SetAdjustBrightness", func(c *gin.Context) {
		input, _ := c.GetRawData()
		jsonRes := SetAdjustBrightness_(input)
		c.JSON(200, jsonRes)
	})
}
