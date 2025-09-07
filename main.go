package main

import (
	"log"
	"net"
)

func main() {
	cantor := NewCantor()
	go cantor.chant()

	l, err := net.Listen("tcp", configAddr())
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	log.Println("server initialized listening on: ", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("got connection from: ", conn.RemoteAddr())
		f := NewFollower(conn, cantor)
		go f.poll()
	}
}
