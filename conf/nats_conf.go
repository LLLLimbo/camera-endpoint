package conf

import "github.com/nats-io/nats.go"

var nats_con *nats.Conn

func ConnectNats(natsAddr *string) error {
	nc, err := nats.Connect(*natsAddr)
	nats_con = nc
	return err
}

func NatsCon() *nats.Conn {
	return nats_con
}
