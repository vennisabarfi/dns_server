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
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println("Error setting up udp connection ", err)
	}

	defer conn.Close()

	// Read from UDP listener in endless loop
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("> ", string(buf[0:]))

		// Write back the message over UPD
		conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
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
