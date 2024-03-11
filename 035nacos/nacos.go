// package nacos

// import (
// 	"fmt"
// 	"time"

// 	"github.com/nacos-group/nacos-sdk-go/clients"
// 	"github.com/nacos-group/nacos-sdk-go/common/constant"
// 	"github.com/nacos-group/nacos-sdk-go/vo"
// )

// func Nacos() {
// 	// 至少一个ServerConfig
// 	serverConfigs := []constant.ServerConfig{
// 		{
// 			IpAddr: "127.0.0.1",
// 			Port:   8848,
// 		},
// 	}

// 	// 创建clientConfig
// 	clientConfig := constant.ClientConfig{
// 		// NamespaceId:         "1cf91be1-d0e3-4494-aef7-b3cb8177e04e", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
// 		NamespaceId:         "datakit", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
// 		TimeoutMs:           5000,
// 		NotLoadCacheAtStart: true,
// 		LogDir:              "tmp/nacos/log",
// 		CacheDir:            "tmp/nacos/cache",
// 		// RotateTime:          "1h",
// 		// MaxAge:              3,
// 		LogLevel: "debug",
// 		Username: "nacos",
// 		Password: "nacos",
// 	}
// 	// 创建动态配置客户端的另一种方式 (推荐)
// 	configClient, err := clients.NewConfigClient(
// 		vo.NacosClientParam{
// 			ClientConfig:  &clientConfig,
// 			ServerConfigs: serverConfigs,
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//获取配置信息
// 	//content, err := configClient.GetConfig(vo.ConfigParam{
// 	//	DataId: "user-web.yaml",
// 	//	Group:  "dev"})
// 	//if err != nil {
// 	//	fmt.Println("GetConfig err: ",err)
// 	//}

// 	//监听配置
// 	err = configClient.ListenConfig(vo.ConfigParam{
// 		DataId: "user-web.yaml",
// 		Group:  "dev",
// 		OnChange: func(namespace, group, dataId, data string) {
// 			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println("err = configClient.ListenConfig ", err)
// 		return
// 	}
// 	for i := 0; i < 1000; i++ {
// 		time.Sleep(time.Second * 2)
// 		fmt.Printf("%d", i)
// 	}

// }

// // func Nacos02() {
// // 	sc := []constant.ServerConfig{
// // 		*constant.NewServerConfig("192.168.10.130", 8848, constant.WithContextPath("/nacos")),
// // 	}
// // 	//设置namespace的id    日志目录
// // 	cc := *constant.NewClientConfig(
// // 		constant.WithNamespaceId("-944c-******-944c-******"),
// // 		constant.WithTimeoutMs(5000),
// // 		constant.WithNotLoadCacheAtStart(true),
// // 		constant.WithLogDir("tmp/nacos/log"),
// // 		constant.WithCacheDir("tmp/nacos/cache"),
// // 		constant.WithLogLevel("debug"),
// // 	)
// // 	//建立连接
// // 	client, err := clients.NewConfigClient(
// // 		vo.NacosClientParam{
// // 			ClientConfig:  &cc,
// // 			ServerConfigs: sc,
// // 		},
// // 	)
// // 	if err != nil {
// // 		fmt.Printf("PublishConfig err:%+v \n", err)
// // 	}
// // 	//获取配置集
// // 	content, err := client.GetConfig(vo.ConfigParam{
// // 		DataId: "user-web.json",
// // 		Group:  "dev",
// // 	})
// // 	//fmt.Println(content)
// // 	//这里是自己实例化的struct
// // 	serverConfig := config.ServerConfig{}
// // 	//想要将一个字符串转换成struct需要去设置这个struct的tag
// // 	json.Unmarshal([]byte(content), &serverConfig)
// // 	fmt.Println(serverConfig)
// // }
