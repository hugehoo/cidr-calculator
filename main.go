package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Print("Input CIDR :")
	reader := bufio.NewReader(os.Stdin)

	ip, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input", err)
		return
	}

	ipString := strings.TrimSpace(ip)
	//fmt.Println("Input Ip : ", ipString)

	cidr, ipNet, err := net.ParseCIDR(ipString)
	if err != nil {
		fmt.Println("Parsing Error ", err)
		return
	}
	fmt.Println("First IP : ", cidr)
	fmt.Println("ipNet : ", ipNet)
	fmt.Println("ipNet.Mask : ", ipNet.Mask)
	fmt.Println("ipNet.IP : ", ipNet.IP)
}
