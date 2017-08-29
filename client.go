package main

import (
	"log"
	"net"
	"os"
)

func ping(conn *net.UDPConn, done chan bool, idx int) {
	n, err := conn.Write([]byte("Ping"))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Printf("%d: Received data: %s", idx, string(buf[:n]))
	done <- true
}

func main() {
	dst := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	conn, err := net.DialUDP("udp", nil, dst)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan bool, 1)
	for i := 0; i < 3; i++ {
		ping(conn, done, i)
		<-done
	}
}
