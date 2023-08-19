package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("/tmp/cb/server.pem", "/tmp/cb/server.key")
	if err != nil {
		log.Println(err)
		return
	}
	certBytes, err := ioutil.ReadFile("/tmp/cb/client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	ln, err := tls.Listen("tcp", ":10443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

// import (
// 	"bufio"
// 	"crypto/tls"
// 	"log"
// 	"net"
// )

// func main() {
// 	cert, err := tls.LoadX509KeyPair("/tmp/cb/server.pem", "/tmp/cb/server.key")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	config := &tls.Config{Certificates: []tls.Certificate{cert}}
// 	ln, err := tls.Listen("tcp", ":443", config)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer ln.Close()
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		go handleConn(conn)
// 	}
// }
// func handleConn(conn net.Conn) {
// 	defer conn.Close()
// 	r := bufio.NewReader(conn)
// 	for {
// 		msg, err := r.ReadString('\n')
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		println(msg)
// 		n, err := conn.Write([]byte("world\n"))
// 		if err != nil {
// 			log.Println(n, err)
// 			return
// 		}
// 	}
// }
