package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var conn *net.UDPConn
var done chan bool

func reply(addr *net.UDPAddr, idx int) {
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
	var err error

	la := flag.String("addr", "127.0.0.1", "local IP address")
	lp := flag.Int("port", 5683, "local port")
	flag.Parse()

	laddr := &net.UDPAddr{
		IP:   net.ParseIP(*la),
		Port: *lp,
	}

	conn, err = net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	done = make(chan bool, 1)
	buf := make([]byte, 1024)
	log.Println("Starting udp server...")
	for i := 0; i < 3; i++ {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		log.Printf("%d: Reciving data: %s from %s", i, string(buf[:n]), addr.String())
		go reply(addr, i)
		<-done
	}
}
