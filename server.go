package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func (s *server) runCMD() {
	for cmd := range s.commands {
		log.Printf("command from %s", cmd.client.alias)
		switch cmd.id {
		case CMD_ALIAS:
			s.updateAlias(cmd)
		case CMD_JOIN:
			s.joinRoom(cmd)
		case CMD_ROOMS:
			s.listRooms(cmd)
		case CMD_MSG:
			s.message(cmd)
		case CMD_QUIT:
			s.quit(cmd)
		}
	}
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newConn(conn net.Conn) {
	log.Println("New Client has connected", conn.RemoteAddr().String())

	c := &client{
		connection: conn,
		alias:      "Anonymous",
		room:       nil,
		commands:   s.commands,
	}

	c.readInput()

}

func (s *server) updateAlias(cmd command) {
	cmd.client.alias = strings.Join(cmd.args, " ")
	cmd.client.msg(fmt.Sprintf("You will be called %s!", cmd.client.alias))
}

func (s *server) joinRoom(cmd command) {
	roomName := strings.Join(cmd.args, " ")
	log.Printf("Join Room name: %s", roomName)
	r, ok := s.rooms[roomName]

	if !ok {
		//Room doesn't exists so create one.
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[cmd.client.connection.RemoteAddr()] = cmd.client

	//quit from previous room
	s.leaveRoom(cmd.client)

	//assign room to user
	cmd.client.room = r

	r.broadcast(cmd.client, fmt.Sprintf("%s has joined the room", cmd.client.alias))
	cmd.client.msg(fmt.Sprintf("%s welcome to room %s", cmd.client.alias, cmd.client.room.name))

}
func (s *server) listRooms(cmd command) {
	var rooms []string
	for room := range s.rooms {
		rooms = append(rooms, room)
	}
	log.Println(rooms)

	cmd.client.msg(fmt.Sprintf("Available rooms are:\n%s", strings.Join(rooms, "\n")))

}
func (s *server) message(cmd command) {

	if cmd.client.room == nil {
		cmd.client.err(errors.New("you must join a room to send message"))
	} else {
		cmd.client.room.broadcast(cmd.client, strings.Join(cmd.args, " "))
	}

}
func (s *server) quit(cmd command) {
	log.Printf("client has disconnected: %s", cmd.client.connection.RemoteAddr().String())

	s.leaveRoom(cmd.client)

	cmd.client.msg("Goodbye !!")

	cmd.client.connection.Close()
}

func (s *server) leaveRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.connection.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.alias))

		if len(c.room.members) == 0 {
			//destroy room with no members
			delete(s.rooms, c.room.name)
		}
	}
}
