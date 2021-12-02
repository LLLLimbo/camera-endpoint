package conf

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/rs/zerolog/log"
)

var badger_con *badger.DB

func ConnectBadger(directory *string) *badger.DB {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions(*directory))
	if err != nil {
		log.Fatal().Err(err)
	}
	badger_con = db
	return db
}

func BadgerCon() *badger.DB {
	return badger_con
}
