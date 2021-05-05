package main

import (
	"fmt"
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, message string) {
	for addr, member := range r.members {
		if addr != sender.connection.RemoteAddr() {
			b_msg := fmt.Sprintf("(%s) %s", sender.alias, message)
			member.msg(b_msg)
		}
	}
}
