package main

import (
	"fmt"
	"net"
	"net/url"
)

// see https://www.kandaoni.com/news/38305.html

func main() {
	inters, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}

	for _, inter := range inters {
		// Get all network card addresses
		//
		// example data
		// [0]: "127.0.0.1"
		// [1]: "::1"
		// [2]: "10.100.65.11"
		// [3]: "fe80::ceee:f501:a51f:da2e"
		// [4]: "172.18.0.1"
		// [5]: "172.31.0.1"
		// [6]: "172.17.0.1"
		// [7]: "fe80::42:5bff:fe17:448c"
		// [8]: "fe80::5824:92ff:feb0:aa6a"
		addrs, err := inter.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipnet, _ := addr.(*net.IPNet)
			fmt.Println("===ip : ", ipnet.IP.String())
		}
	}

}

func testIP(remote string) string {
	host := ""
	// try get 'host' tag from remote URL.
	if u, err := url.Parse(remote); err == nil && u.Host != "" { // like scheme://host:[port]/...
		host = u.Host
		if ip, _, err := net.SplitHostPort(u.Host); err == nil {
			host = ip
		}
	} else { // not URL, only IP:Port
		if ip, _, err := net.SplitHostPort(remote); err == nil {
			host = ip
		}
	}

	if host == "" {
		address := net.ParseIP(remote)
		if address == nil {
			return ""
		} else {
			return address.String()
		}
	}

	return ""
}
