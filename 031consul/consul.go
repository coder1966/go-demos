package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

var (
	keyC   = "data/c/n"
	keyP   = "data/p/metric/n.p"
	valueC string
	valueP string
)

type valueStr struct {
	key   string
	value string
}

func main() {
	ConsulDo(1, "127.0.0.1")
}

func ConsulDo(index int, ip string) {
	fmt.Println("===consul===index:", index)
	// 创建终端
	client, err := api.NewClient(&api.Config{
		Address: "http://" + ip + ":8500",
	})
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// 获得KV句柄
	conn := client.KV()

	// 创建2个值，valueC，valueP
	creatValues(index)

	// 处理
	switch index {
	case 0:
		handle(conn, valueStr{key: keyC, value: valueC}, valueStr{key: keyP, value: valueP})
	case 1:
		handle(conn, valueStr{key: keyC, value: valueC}, valueStr{key: keyP, value: valueP})
	case 2:
		handle(conn, valueStr{key: keyC, value: valueC}, valueStr{key: keyP, value: valueP})
	case 3:
		handle(conn, valueStr{key: keyP, value: valueP})
	case 4:
		handle(conn, valueStr{key: keyC, value: valueC})
	}

	get(conn, keyP)
	get(conn, keyC)

}

func handle(conn *api.KV, valueStr ...valueStr) {
	// 循环处理，可能有1~2组数据
	for _, v := range valueStr {
		key := v.key
		value := v.value
		if value == "" {
			del(conn, key)
		} else {
			add(conn, key, value)
		}

	}
}

// 增
func add(conn *api.KV, path, value string) {
	if path == "" {
		return
	}

	p := &api.KVPair{Key: path, Value: []byte(value), Flags: 32}
	_, err := conn.Put(p, nil)
	if err != nil {
		fmt.Println("consul put error: ", err)
		return
	}

	fmt.Println("consul 创建成功", path)
}

// 删除
func del(conn *api.KV, path string) {
	// 初始化context

	_, err := conn.Delete(path, nil)
	if err != nil {
		fmt.Println("consul del error: ", err)
		return
	}

	fmt.Println("consul 删除成功: ")
}

// 查
func get(conn *api.KV, path string) {

	// 取回数据
	pair, _, err := conn.Get(path, nil)
	if err != nil {
		fmt.Println("consul 取 error: ", err)
		return
	}
	_ = pair
	if pair == nil {
		fmt.Println("consul 取回来KV: ", pair)
	} else {
		fmt.Println("consul 取回来KV: ", pair.Key, string(pair.Value))
	}

}

func creatValues(index int) {
	switch index {
	case 0:
		valueC = ""
		valueP = `add_key(tconsul, 0)`
	case 1:
		valueC = `
		DATA
`
		valueP = `add_key(tconsul, 1)`
	case 2:
		valueC = `
		DATA	
	    DATA
`
		valueP = `add_key(tconsul, 2)`
	case 3:
		valueP = `add_key(tconsul, 3)`
	case 4:
		valueC = `
		DATA
		DATA
		DATA
		DATA
`
	default:
	}
}
