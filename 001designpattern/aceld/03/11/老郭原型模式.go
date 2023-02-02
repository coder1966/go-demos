package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
原型模式，就是克隆，生成开销大,例如大量数据准备+数据库连接。
缓存一个对象，下次用的话，就直接克隆一个对象返回
*/

type CPU struct{ Name string }
type ROM struct{ Name string }
type Disk struct{ Name string }

type Computer struct {
	Cpu  CPU
	Rom  ROM
	Disk Disk
}

// 浅拷贝
func (c *Computer) Clone() *Computer {
	resume := *c
	return &resume
}

func (c *Computer) Backup() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, c); err != nil {
		panic(err.Error())
	}
	return pc
}

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

// 主程序
func main() {
	cpu := CPU{"英特尔"}
	rom := ROM{"金士顿"}
	disk := Disk{"三星"}

	c := Computer{
		Cpu:  cpu,
		Rom:  rom,
		Disk: disk,
	}

	c1 := c.Backup()
	fmt.Printf("c1: %v \n", *c1)

}
