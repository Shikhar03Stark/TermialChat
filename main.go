package main

import (
	"log"
	"net"
)

func main() {
	ser := newServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()
	log.Printf("Server running on :8080\n")

	go ser.runCMD()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go ser.newConn(conn)
	}

}
