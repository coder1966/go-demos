package main

import (
	"fmt"
	"godemos/923gethttpandtls/collectors"
	"net/http"
	"net/url"
)

type Input struct {
	Scheme         string
	Host           string
	Port           int
	AdditionalPort int
	User           string
	Password       string
	TLSOpen        bool
	CacertFile     string
	CertFile       string
	KeyFile        string

	client *collectors.Client
}

func main() {
	// http://Administrator:123456@127.0.0.1:8091"
	ipt := &Input{
		Scheme:         "http",
		Host:           "127.0.0.1",
		Port:           8091,
		AdditionalPort: 9102,
		User:           "Administrator",
		Password:       "123456",
		TLSOpen:        false,
		CacertFile:     "",
		CertFile:       "/var/cb/clientcertfiles/travel-sample.pem",
		KeyFile:        "/var/cb/clientcertfiles/travel-sample.key",
	}

	testCase := map[string]string{
		"pools/default": "Port",
		"api/v1/stats":  "AdditionalPort",
	}

	err := ipt.Init()
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}

	for path, port := range testCase {
		ipt.client.Get(path, port)
	}
}

func (i *Input) Init() error {
	if err := i.parseURL(); err != nil {
		return err
	}

	cli, err := i.newClient()
	if err != nil {
		return err
	}
	i.client = cli

	return nil
}

func (i *Input) parseURL() error {
	u := i.Scheme + "://" + i.Host + ":" + fmt.Sprint(i.Port)
	if _, err := url.Parse(u); err != nil {
		return fmt.Errorf("parse url %s failed: %w", u, err)
	}

	if i.Port < 1 || i.Port > 65535 {
		return fmt.Errorf("parse port error: %d", i.Port)
	}

	if i.AdditionalPort < 1 || i.AdditionalPort > 65535 {
		return fmt.Errorf("parse additional port error: %d", i.AdditionalPort)
	}

	if !(i.Scheme == "http" || i.Scheme == "https") {
		return fmt.Errorf("parse additional scheme error: %s", i.Scheme)
	}

	return nil
}

func (i *Input) newClient() (*collectors.Client, error) {
	opt := &collectors.Option{
		TLSOpen:        i.TLSOpen,
		CacertFile:     i.CacertFile,
		CertFile:       i.CertFile,
		KeyFile:        i.KeyFile,
		Scheme:         i.Scheme,
		Host:           i.Host,
		Port:           i.Port,
		AdditionalPort: i.AdditionalPort,
		User:           i.User,
		Password:       i.Password,
	}

	client := &collectors.Client{
		Opt: opt,
	}
	client.SetClient(&http.Client{})

	return client, nil
}
