package httplib

import (
	"fmt"
	"net/http"
)

func NewServer(addr string, port int) (*http.Server, error) {
	a := fmt.Sprintf("%s:%d", addr, port)
	server := http.Server{
		Addr: a,
	}
	return &server, nil
}
