package gio_redis

import (
	"net"
)

type Client struct {
	conn net.Conn	
}

func Dial(address string) *Client {
	clt := &Client{}
	clt.conn, _ = net.Dial("tcp", address)
	return clt
}