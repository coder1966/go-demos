// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package collectors

import (
	"fmt"
	"io/ioutil"
)

func (c *Client) get(u string) error {
	resp, err := c.request(u)
	if err != nil {
		return fmt.Errorf("failed to Get %s : %w", u, err)
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s : %w", u, err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to Get 200 response status from %s : %d", u, resp.StatusCode)
	}

	fmt.Println("OUTPUT : ", string(bts))

	return nil
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s://%s:%d/%s", c.Opt.Scheme, c.Opt.Host, c.Opt.Port, path)
}

func (c *Client) indexerURL(path string) string {
	return fmt.Sprintf("%s://%s:%d/%s", c.Opt.Scheme, c.Opt.Host, c.Opt.AdditionalPort, path)
}
