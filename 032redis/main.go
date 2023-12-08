package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	keyC   = "/data/c/net"
	keyP   = "/data/p/net.p"
	valueC string
	valueP string
)

type valueStr struct {
	key   string
	value string
}

func main() {
	RedisDo(1, "127.0.0.1")
}
func RedisDo(index int, ip string) {

	fmt.Println("===redis===index:", index)

	// 初始化redis客户端
	conn := redis.NewClient(&redis.Options{
		Addr: ip + ":6379",
		// Password: "654123", // no password set
		DB: 0, // use default DB
	})

	// 创建2个值，valueC，valueP
	creatValues(index)

	// 批量写数据
	batchLen := 3000
	// 写字符串
	keySlice := []string{}
	for i := 0; i < batchLen; i++ {
		l := strings.Trim(fmt.Sprintf("%d", i), "")
		keySlice = append(keySlice, "keySlice"+l)
	}
	for i, v := range keySlice {
		handle(conn, valueStr{key: v, value: strings.Repeat("0123456789abcdef", i)})
	}
	// 写字符串
	keyZSets := []string{}
	for i := 0; i < batchLen; i++ {
		l := strings.Trim(fmt.Sprintf("%d", i), "")
		keyZSets = append(keySlice, "keyZSet"+l)
	}
	for i, keyZSet := range keyZSets {
		zSets := []*redis.Z{}
		for j := 0; j < i+1; j++ {
			zSet := &redis.Z{Score: float64(j), Member: "Member" + strings.Trim(fmt.Sprintf("%d", j), "")}
			zSets = append(zSets, zSet)
		}
		handleZSet(conn, keyZSet, zSets)
	}

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

	// get(conn, keyP)
	// get(conn, keyC)

}

func RedisDo02(index int, ip string) {

	fmt.Println("===redis===index:", index)

	// 初始化redis客户端
	conn := redis.NewClient(&redis.Options{
		Addr: ip + ":6379",
		// Password: "654123", // no password set
		DB: 0, // use default DB
	})
	// 初始化context
	ctx := context.Background()

	for i := 0; i < index; i++ {
		keyC = "keyC" + fmt.Sprint(i)
		valueC = "valueC" + fmt.Sprint(i)

		// 写
		err := conn.Set(ctx, keyC, valueC, 0).Err()
		if err != nil {
			fmt.Println("redis put error: ", err)
		}
		fmt.Println("redis put : ", i)
	}

}

func RedisDo03(index int, ip string) {
	// 初始化redis客户端
	conn := redis.NewClient(&redis.Options{
		Addr: ip + ":6379",
		// Password: "654123", // no password set
		DB: 0, // use default DB
	})
	// 批量写数据
	batchLen := index
	// 写字符串
	keySlice := []string{}
	for i := 0; i < batchLen; i++ {
		l := strings.Trim(fmt.Sprintf("%d", i), "")
		keySlice = append(keySlice, "wo"+l)
	}
	for i, v := range keySlice {
		handle(conn, valueStr{key: v, value: strings.Repeat("0123456789abcdef", i)})
	}
}

func handleZSet(conn *redis.Client, k string, v []*redis.Z) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_ = conn.ZAdd(ctx, k, v...).Err()
}

func handle(conn *redis.Client, valueStr ...valueStr) {
	// 循环处理，可能有1~2组数据
	for _, v := range valueStr {
		key := v.key
		value := v.value
		if value == "" {
			del(conn, key)
		} else {
			add(conn, key, value)
		}

		// 发布订阅
		if strings.Index(v.key, "/data/c") == 0 {
			publish(conn, "/data/c")
		} else if strings.Index(v.key, "/data/p") == 0 {
			publish(conn, "/data/p")
		}

	}
}

// 发布订阅
func publish(conn *redis.Client, path string) {
	// 初始化context
	ctx := context.Background()

	publishPath := "__keyspace@0__:" + path + "*"

	// 发布订阅
	// n, err := conn.Publish(ctx, "__keyspace@0__:/data*", "set").Result()
	n, err := conn.Publish(ctx, publishPath, "set").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d clients received the message\n", n)
}

// 增
func add(conn *redis.Client, path, value string) {
	if path == "" {
		return
	}

	// 初始化context
	ctx := context.Background()

	// 写
	err := conn.Set(ctx, path, value, 0).Err()
	if err != nil {
		fmt.Println("redis put error: ", err)
	}
	fmt.Println("创建成功", path)
}

// 删除
func del(conn *redis.Client, path string) {
	// 初始化context
	ctx := context.Background()
	ret := conn.Del(ctx, path)

	fmt.Println("redis删除成功条数 : ", ret.Val())
}

// 查
func get(conn *redis.Client, path string) {

	// 初始化context
	ctx := context.Background()

	// 读
	val, err := conn.Get(ctx, path).Result()
	if err != nil {
		fmt.Println("redis get error: ", err)
	}
	fmt.Println("redis get : ", path, val)
}

func creatValues(index int) {
	switch index {
	case 0:
		valueC = ""
		valueP = ""
	case 1:
		valueC = `
		DATA
`
		valueP = `add_key(tredis, 1)`
	case 2:
		valueC = `
		DATA
		DATA
`
		valueP = `add_key(tredis, 2)`
	case 3:
		valueP = `add_key(tredis, 3)`
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
