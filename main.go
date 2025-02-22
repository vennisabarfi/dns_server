package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	port := ":1053"

	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Println("Error resolving udp server to port ", err)
	}
	fmt.Println(udpAddr)

	// start listening for udp packages on the address
	conn, err := net.ListenPacket("udp", port)
	if err != nil {
		log.Println("Error setting up udp connection ", err)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println("Error buffering data from udp server", err)
			continue
		}

		go serve(conn, addr, buf[:n])

		log.Println("Server buffering data")
	}

}

func serve(conn net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	conn.WriteTo(buf, addr)

}

func HealthServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is healthy")
	w.Write([]byte("Server is healthy"))

}
