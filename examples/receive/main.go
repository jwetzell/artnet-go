package main

import (
	"fmt"
	"net"

	"github.com/jwetzell/artnet-go"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":6454")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	if err != nil {
		panic(err)
	}

	for {
		bytesRead, _, err := connection.ReadFromUDP(buffer)

		if err != nil {
			panic(err)
		}

		packet, err := artnet.Decode(buffer[0:bytesRead])

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%+v\n", packet)
	}
}
