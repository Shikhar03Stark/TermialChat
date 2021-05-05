package main

type commandID int

const (
	CMD_ALIAS commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	client *client
	id     commandID
	args   []string
}
