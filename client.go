package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var conn *net.UDPConn
var done chan bool

func ping(idx int) {
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
	var err error

	da := flag.String("addr", "127.0.0.1", "destination IP address")
	dp := flag.Int("port", 5683, "destination IP address")
	flag.Parse()

	daddr := &net.UDPAddr{
		IP:   net.ParseIP(*da),
		Port: *dp,
	}

	conn, err = net.DialUDP("udp", nil, daddr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	done = make(chan bool, 1)
	for i := 0; i < 3; i++ {
		ping(i)
		<-done
	}
}
