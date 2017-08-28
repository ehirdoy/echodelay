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

	buf := make([]byte, 1024)
	for i := 0; i < 3; i++ {
		n, err := conn.Write([]byte("Ping"))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		n, err = conn.Read(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		log.Printf("%d: Received data: %s", i, string(buf[:n]))
	}
}
