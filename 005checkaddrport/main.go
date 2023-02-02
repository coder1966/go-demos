package main

import (
	"fmt"
	"net/netip"
	"net/url"
)

func main() {
	ad1 := "123.123.4.5"
	parse, err := url.Parse(ad1)
	fmt.Println(parse, err)

	ad2 := "123.123.4.5:90"
	AddrPort, err := netip.ParseAddrPort(ad2)
	AddrPort.Addr()

	fmt.Println(AddrPort, AddrPort.Addr(), AddrPort.Port(), err)
	AddrPort, err = netip.ParseAddrPort(ad1)
	fmt.Println(AddrPort, err)

}
