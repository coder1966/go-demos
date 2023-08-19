// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package collectors used to collect metrics.
package collectors

import (
	"fmt"
	"net/http"
)

type MetricContext struct {
	ClusterName  string
	NodeHostname string
	BucketName   string
	Keyspace     string
	Source       string
	Target       string
}

type Option struct {
	TLSOpen        bool
	CacertFile     string
	CertFile       string
	KeyFile        string
	Scheme         string
	Host           string
	Port           int
	AdditionalPort int
	User           string
	Password       string
}

type Client struct {
	client *http.Client
	Opt    *Option
}

func (c *Client) Get(path, port string) {
	switch port {
	case "Port":
		err := c.get(c.url(path))
		if err != nil {
			fmt.Println(" error: ", err)
		}
	case "AdditionalPort":
		err := c.get(c.indexerURL(path))
		if err != nil {
			fmt.Println(" error: ", err)
		}
	}

}

func (c *Client) request(url string) (*http.Response, error) {
	req, err := c.getReq(url)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) getReq(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if c.Opt.User != "" {
		req.SetBasicAuth(c.Opt.User, c.Opt.Password)
	}

	return req, err
}

func (c *Client) SetClient(cli *http.Client) {
	c.client = cli
}
