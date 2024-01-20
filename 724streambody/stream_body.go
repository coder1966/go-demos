package a0601streambody

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

const bufLen = 128

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, bufLen)
	},
}

func streamBody(res http.ResponseWriter, req *http.Request) (result []string, err error) {
	defer res.WriteHeader(http.StatusNoContent)

	r := req.Body
	defer r.Close()
	br := bufio.NewReader(r)

	buf := bufPool.Get().([]byte)
	defer func() {
		// buf.Reset()
		bufPool.Put(&buf)
	}()

	bufEnd := false
	tail := len(buf)
	lastTail := tail

	for !bufEnd {
		time.Sleep(time.Microsecond * 10)
		// 1 读取buf，可能 前面有残留
		// 1.1 先搬移残留
		if tail != lastTail {
			// 有残留
			if len(buf)-tail >= len(buf)/2 {
				// 残留大于一半，申请新的
				bufNew := bufPool.Get().([]byte)
				defer func() {
					// buf.Reset()
					bufPool.Put(&buf)
				}()
				copy(bufNew, buf[tail:])
				buf = bufNew
			} else {
				// 残留小于一半，直接copy
				copy(buf, buf[tail:])
			}
		}

		// 1.2 读取
		n, err := br.Read(buf[len(buf)-tail:])
		if err != nil {
			if err == io.EOF && n == 0 {
				if lastTail-tail != 0 {
					result = append(result, string(buf[:lastTail-tail]))
				}
				return result, nil
			}
			return nil, err
		}

		lastTail = n + len(buf) - tail

		// 2 处理，可能需要残留一些Byte
		r2 := bytes.NewReader(buf[:n+len(buf)-tail])
		br2 := bufio.NewReader(r2)
		tail = 0
		for {
			l, e := br2.ReadBytes('\n')
			if e != nil && len(l) == 0 {
				break
			}
			if l[len(l)-1] != byte(10) {
				// 不是换行结尾 （整个数据源的尾巴，另外处理）
				break
			}
			result = append(result, string(l))
			tail += len(l)
		}
	}
	return
}
