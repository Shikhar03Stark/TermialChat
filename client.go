package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	connection net.Conn
	alias      string
	room       *room
	commands   chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.connection).ReadString('\n')

		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")

		cmd := args[0]

		switch cmd {
		case "/alias":
			c.commands <- command{
				id:     CMD_ALIAS,
				client: c,
				args:   args[1:],
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args[1:],
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args[1:],
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args[1:],
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args[1:],
			}
		default:
			c.err(fmt.Errorf("unknow command %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.connection.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(message string) {
	c.connection.Write([]byte("> " + message + "\n"))
}
