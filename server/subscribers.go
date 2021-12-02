package server

import (
	"camera-endpoint/conf"
	"camera-endpoint/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
)

func SubCloudMsg() *nats.Subscription {
	nc := conf.NatsCon()
	subscription, err := nc.Subscribe("d.1.s1206.78E4A0164CCE064A4489013E66FB7A7D", func(msg *nats.Msg) {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		data := msg.Data

		datamap := make(map[string]interface{})
		err := json.Unmarshal(data, &datamap)
		util.LogError(err)

		id := datamap["id"]
		var res interface{}
		if id == "setSyncParking" {
			res = SetSynchronizeParking_(data)
		}
		if id == "setBrt" {
			res = SetAdjustBrightness_(data)
		}

		jsonRes, _ := json.Marshal(res)

		//If need to reply
		if msg.Reply != "" {
			replyMsg := nats.Msg{
				Subject: msg.Reply,
				Data:    jsonRes,
			}
			err := nc.PublishMsg(&replyMsg)
			util.LogError(err)
		}

	})
	util.LogError(err)

	return subscription
}
