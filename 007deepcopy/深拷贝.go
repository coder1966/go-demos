package main

import (
	"bytes"
	"encoding/gob"
)

func deepCopy(dst, src interface{}) error {
	// 缓冲区
	var buf bytes.Buffer
	// 调用gob，创建编码器，buf地址给他，对src编码
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	// 缓冲区内容，解码
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)

}
