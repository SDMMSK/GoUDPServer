// Build: go build -ldflags "-s -w"
// Run: goudpserver -ip=0.0.0.0 -port=10003
// Test: ncat -u localhost 10003

package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
)

func lstn(connection *net.UDPConn, alarm chan struct{}) {
	buffer := make([]byte, 1024)
	n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		n, remoteAddr, err = connection.ReadFromUDP(buffer)
		fmt.Println("from", remoteAddr, "-", buffer[:n])
	}

	fmt.Println("Listener failed (", err, ")!")
	alarm <- struct{}{}
}

func main() {
	argIP := flag.String("ip", "0.0.0.0", "Listen IP address")
	argPort := flag.Int("port", 10003, "Listen Port number")
	flag.Parse()

	fmt.Println("IP Address: ", *argIP)
	fmt.Println("Port Number: ", *argPort)

	addr := net.UDPAddr{
		Port: *argPort,
		IP:   net.ParseIP(*argIP),
	}
	connection, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	alarm := make(chan struct{})
	fmt.Println("CPU Threads: ", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go lstn(connection, alarm)
	}

	msg := <-alarm
	fmt.Println(msg)
}
