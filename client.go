package main

import (
	"log"
	"net"
	"os"
)

func main() {
	dst := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	conn, err := net.Dial("udp", dst.String())
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Ping"))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	recvBuf := make([]byte, 1024)
	n, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Printf("Received data: %s", string(recvBuf[:n]))
}
