package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ipaddr *IPAddr) String() string {
	var string_ipaddr [4]string
	for index, hex := range ipaddr {
		if hex == 0 {
			string_ipaddr[index] = "0"
		}
		num := ""
		for ; hex > 0; hex /= hex {
			num = string(hex%10) + num
		}
		string_ipaddr[index] = num
	}
	return strings.Join(string_ipaddr[:], ":")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
