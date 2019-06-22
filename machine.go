package main

import (
	"fmt"
	"gutil/machine"
	"net"
	"strings"
)

func getMac() string {

	ifas, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, ifa := range ifas {
		if strings.Contains(ifa.Name, "eth") {
			return ifa.HardwareAddr.String()
		}
	}
	return ""
}

func main() {

	fmt.Println(machine.ActiveNetIFace())
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
