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
	network, firstIp, lastIp, broadcast := CalculateIPRange(ipString)
	fmt.Println("네트워크 주소:", network)
	fmt.Println("첫 번째 호스트 주소:", firstIp)
	fmt.Println("마지막 호스트 주소:", lastIp)
	fmt.Println("브로드캐스트 주소:", broadcast)
}

func CalculateIPRange(cidr string) (network, firstIp, lastIp, broadcast net.IP) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("CIDR 파싱 중 오류:", err)
		return nil, nil, nil, nil
	}

	network = ipNet.IP
	firstIP, lastIP, broadcast := make(net.IP, len(network)), make(net.IP, len(network)), make(net.IP, len(network))
	copy(firstIP, network)
	copy(lastIP, network)
	copy(broadcast, network)

	for i := range network {
		network[i] &= ipNet.Mask[i]
	}

	for i := len(network) - 1; i >= 0; i-- {
		lastIP[i] |= ^ipNet.Mask[i]
		firstIP[i] |= 1
	}

	for i := range broadcast {
		broadcast[i] = lastIP[i]
	}
	return network, firstIP, lastIP, broadcast
}
