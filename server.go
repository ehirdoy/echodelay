package main

import (
	"log"
	"net"
	"os"
	"time"
)

func reply(conn *net.UDPConn, addr *net.UDPAddr, done chan bool, idx int) {
	dur := time.Duration(idx) * time.Second
	log.Printf("%d: Sleeping %d[sec]", idx, idx)
	//dur := time.Duration(idx) * time.Minute
	//log.Printf("%d: Sleeping %d[min]", idx, idx)
	time.Sleep(dur)
	log.Printf("%d: Sending data..", idx)
	conn.WriteTo([]byte("Pong"), addr)
	log.Printf("%d: Complete Sending data..", idx)
	done <- true
}

func main() {
	src := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	conn, err := net.ListenUDP("udp", src)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	done := make(chan bool, 1)
	buf := make([]byte, 1024)
	log.Println("Starting udp server...")
	for i := 0; i < 3; i++ {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		log.Printf("%d: Reciving data: %s from %s", i, string(buf[:n]), addr.String())
		go reply(conn, addr, done, i)
		<-done
	}
}
