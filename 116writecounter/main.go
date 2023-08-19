package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.baidu.com"
	// url = "http://127.0.0.1:9363/metrics"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	defer resp.Body.Close()
	// A agent used to count bytes.
	wCounter := &writeCounter{}
	testOutData(io.TeeReader(resp.Body, wCounter))
	fmt.Println("代理拿到的字节数：= ", wCounter.total)
}
func testOutData(in io.Reader) {
	total := 0
	buf := make([]byte, 256) // 最好4096 or 8192
	for {
		n, err := in.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		fmt.Println(string(buf))
		total += n
	}
	fmt.Println("实际拿到的字节数：= ", total)
}

// A agent used to count bytes.
type writeCounter struct {
	total uint64
}

// A agent used to count bytes.
func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.total += uint64(len(p))
	return 0, nil
}
