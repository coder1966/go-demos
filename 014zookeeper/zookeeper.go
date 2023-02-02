package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	keyC   = "/datakit/confd/netstat"
	keyP   = "/datakit/pipeline/metric/disk.p"
	valueC string
	valueP string
)

type valueStr struct {
	key   string
	value string
}

func main() {
	ZookeeperDo(1, "127.0.0.1")
}
func ZookeeperDo(index int, ip string) {
	fmt.Println("===zookeeper===index:", index)
	hosts := []string{ip + ":2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println("conn, _, err := zk.Connect error: ", err)
	}
	defer conn.Close()
	// 创建一级目录节点
	onlyAdd(conn, "/data", "")
	// 创建二级目录节点
	onlyAdd(conn, "/data/c", "")
	onlyAdd(conn, "/data/p", "")
	// 创建三级目录节点
	onlyAdd(conn, "/data/p/metrics", "")
	onlyAdd(conn, "/data/p/metric", "")
	onlyAdd(conn, "/data/p/network", "")
	onlyAdd(conn, "/data/p/keyevent", "")
	onlyAdd(conn, "/data/p/object", "")
	onlyAdd(conn, "/data/p/custom_object", "")
	onlyAdd(conn, "/data/p/logging", "")
	onlyAdd(conn, "/data/p/tracing", "")
	onlyAdd(conn, "/data/p/rum", "")
	onlyAdd(conn, "/data/p/security", "")
	onlyAdd(conn, "/data/p/profiling", "")

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

	// 查
	// get(conn, keyC)
	// get(conn, keyP)
}

func handle(conn *zk.Conn, valueStr ...valueStr) {
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

// 只增加，有了，就不修改
func onlyAdd(conn *zk.Conn, path, value string) {
	if path == "" {
		return
	}

	var data = []byte(value)

	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	_, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Println("only 创建 error: ", err, path)
		return
	}
	fmt.Println("only 创建成功", path)
}

// 增
func add(conn *zk.Conn, path, value string) {
	if path == "" {
		return
	}

	var data = []byte(value)

	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	_, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Println("创建 error: ", err, path)
		modify(conn, path, value)
		return
	}
	fmt.Println("创建成功", path)
}

// 删除
func del(conn *zk.Conn, path string) {
	if path == "" {
		return
	}
	_, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		fmt.Println("zookeeper 删除 error: ", err)
		return
	}
	fmt.Println("zookeeper 删除成功 : ", path)
}

// 查
func get(conn *zk.Conn, path string) {
	if path == "" {
		return
	}
	data, _, err := conn.Get(path)
	if err != nil {
		fmt.Println("查 error: ", err)
		return
	}
	fmt.Println("查询 : ", path, string(data))
}

// 改
func modify(conn *zk.Conn, path, value string) {
	if path == "" {
		return
	}
	var data = []byte(value)
	_, sate, _ := conn.Get(path)
	_, err := conn.Set(path, data, sate.Version)
	if err != nil {
		fmt.Println("修改 error: ", err)
		return
	}
	fmt.Println("修改成功", path)
}

// // 不阻塞监听 ExistsW
// func creatWatch(path string) {

// 	// 创建监听的option，用于初始化zk
// 	eventCallbackOption := zk.WithEventCallback(callback)
// 	if path == "" {
// 		return
// 	}
// 	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*5, eventCallbackOption)
// 	if err != nil {
// 		fmt.Println(" error: ", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// 开始监听path
// 	_, _, _, err = conn.ExistsW(path)
// 	if err != nil {
// 		fmt.Println("监听 error: ", err)
// 		return
// 	}

// 	_, _, _, err = conn.ExistsW(path)
// 	if err != nil {
// 		fmt.Println("监听 error: ", err)
// 		return
// 	}
// 	// 触发监听

// }

// // 阻塞监听
// func creatWatchNew(path string) {

// 	// 创建监听的option，用于初始化zk

// 	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*5)
// 	if err != nil {
// 		fmt.Println(" error: ", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// 开始监听path
// 	_, _, _, err = conn.ChildrenW(path)
// 	if err != nil {
// 		fmt.Println("监听 error: ", err)
// 		return
// 	}

// 	fmt.Println("监听 触发: ")

// }

// func callback(event zk.Event) {
// 	fmt.Println("@@@@@@@@@@@@")
// 	fmt.Println("callback path : ", event.Path)
// 	fmt.Println("callback type : ", event.Type.String())
// 	fmt.Println("callback state : ", event.State.String())
// 	fmt.Println("@@@@@@@@@@@@")
// }

func creatValues(index int) {
	switch index {
	case 0:
		valueC = ""
		valueP = ""
	case 1:
		valueC = `
DATA
`
		valueP = `add_key(tconfd, 1)`
	case 2:
		valueC = `
DATA
DATA

`
		valueP = `add_key(tconfd, 2)`
	case 3:
		valueP = `add_key(tconfd, 3)`
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
