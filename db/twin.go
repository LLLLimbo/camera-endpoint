package db

import (
	"camera-endpoint/conf"
	"camera-endpoint/util"
	"fmt"
	"github.com/dgraph-io/badger/v3"
)

func SetProperty(k []byte, v []byte) {
	db := conf.BadgerCon()
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		util.LogError(err)
		return err
	})
	util.LogError(err)
}

func GetProperty(k []byte) []byte {
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
