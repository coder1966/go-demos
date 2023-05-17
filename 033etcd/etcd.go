package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"go.etcd.io/etcd/client"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

var (
	keyC   = "/data/c/u"
	keyP   = "/data/p/metric/c.p"
	valueC string
	valueP string
)

type valueStr struct {
	key   string
	value string
}

func main() {
	Etcdv3Do(1, "127.0.0.1")
}
func Etcdv3Do(index int, ip string) {
	fmt.Println("===etcdv3===index:", index)

	conn, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{ip + ":2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(" error: ", err)
	}
	defer conn.Close()

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

	// get(conn, keyP)
	// get(conn, keyC)

}

// NewEtcdClient returns an *etcd.Client with a connection to named machines.
func NewEtcdClient(machines []string, cert, key, caCert string, clientInsecure bool, basicAuth bool, username string, password string) (*Client, error) {
	var c client.Client
	var kapi client.KeysAPI
	var err error
	var transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: clientInsecure,
	}

	cfg := client.Config{
		Endpoints:               machines,
		HeaderTimeoutPerRequest: time.Duration(3) * time.Second,
	}

	if basicAuth {
		cfg.Username = username
		cfg.Password = password
	}

	if caCert != "" {
		certBytes, err := ioutil.ReadFile(caCert)
		if err != nil {
			return &Client{kapi}, err
		}

		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM(certBytes)

		if ok {
			tlsConfig.RootCAs = caCertPool
		}
	}

	if cert != "" && key != "" {
		tlsCert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return &Client{kapi}, err
		}
		tlsConfig.Certificates = []tls.Certificate{tlsCert}
	}

	transport.TLSClientConfig = tlsConfig
	cfg.Transport = transport

	c, err = client.New(cfg)
	if err != nil {
		return &Client{kapi}, err
	}

	kapi = client.NewKeysAPI(c)
	return &Client{kapi}, nil
}

// Client is a wrapper around the etcd client
type Client struct {
	client client.KeysAPI
}

func handle(conn *etcdv3.Client, valueStr ...valueStr) {
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
func add(conn *etcdv3.Client, path, value string) {
	if path == "" {
		return
	}

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := conn.Put(ctx, path, value)
	cancel()
	if err != nil {
		fmt.Println(" error: ", err)
	}

	fmt.Println("创建成功", path)
}

// 删除
func del(conn *etcdv3.Client, path string) {
	// 删除多个key,以什么为前缀
	if _, err := conn.Delete(context.TODO(), path, etcdv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}
	// //第三个参数可选,clientv3.WithPrevKV()表示会赋值deleteResp.PrevKvs
	// if deleteResp, err := conn.Delete(context.TODO(), "/cron/jobs/job1", etcdv3.WithPrevKV()); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //被删除之前的k和v,上面第三个参数要设置,否则是没有的
	// if len(deleteResp.PrevKvs) != 0 {
	// 	for _, v := range deleteResp.PrevKvs {
	// 		fmt.Println(string(v.Key))
	// 		fmt.Println(string(v.Value))
	// 	}
	// }

	fmt.Println("删除成功 : etcd")
}

// 查
func get(conn *etcdv3.Client, path string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := conn.Get(ctx, path, etcdv3.WithPrefix())
	defer cancel()
	if err != nil {
		fmt.Println("取 etcdV3  error: ", err)
	}
	fmt.Printf("取回 etcdV3 KV条数 %d  ;", len(resp.Kvs))
	for _, v := range resp.Kvs {
		fmt.Printf("取回 etcdV3 KV %s : %s", v.Key, v.Value)
	}
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
		valueP = `add_key(tetcdv3, 1)`
	case 2:
		valueC = `
		DATA
		DATA

`
		valueP = `add_key(tetcdv3, 2)`
	case 3:
		valueP = `add_key(tetcdv3, 3)`
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
