package utilities

import (
	"encoding/binary"
	"net"
)

func GetLocalIPv4() (ip4 uint32) {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ip := ipnet.IP.To4(); ip != nil && !ip.IsLoopback() &&
				(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168) {
				ip4 = binary.BigEndian.Uint32(ip)
				break
			}
		}
	}
	return
}
