package db

import (
	"camera-endpoint/conf"
	"camera-endpoint/util"
	"fmt"
	"github.com/dgraph-io/badger/v3"
)

func SimpleSet(k []byte, v []byte) {
	db := conf.BadgerCon()
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		return err
	})
	util.LogError(err)
}

func SimpleGet(k []byte) []byte {
	db := conf.BadgerCon()
	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		val, err := txn.Get(k)
		if err != nil {
			return err
		}
		fmt.Println(val)
		valCopy, _ = val.ValueCopy(nil)
		return nil
	})
	util.LogError(err)
	return valCopy
}

func GetDeviceId() []byte {
	return SimpleGet([]byte("deviceId"))
}
