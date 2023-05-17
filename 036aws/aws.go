package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/smithy-go"
)

var (
	keyC   = "/datakit/confd/host/mem.conf"
	keyP   = "/datakit/pipeline/metric/mem.p"
	valueC string
	valueP string
)

type valueStr struct {
	key   string
	value string
}

type Cfg struct {
	AccessKeyID     string
	SecretAccessKey string
}

func AwsDo(index int, ip string) {
	fmt.Println("===aws===index:", index)

	// v := viper.New()
	// v.SetConfigName("key.password.sample") // 配置文件名 (不带扩展格式)
	// v.SetConfigType("ini")                 // 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
	// v.AddConfigPath("./")                  // 配置文件的路径
	// err := v.ReadInConfig() //找到并读取配置文件
	// if err != nil {         // 捕获读取中遇到的error
	// 	fmt.Println("Fatal error config file: %w ", err)
	// 	return
	// }
	// if !v.IsSet("default.aws_access_key_id") {
	// 	fmt.Println("not have default.aws_access_key_id  ")
	// 	return
	// }
	// if !v.IsSet("default.ws_secret_access_key") {
	// 	fmt.Println("not default.ws_secret_access_key  ")
	// 	return
	// }
	// id := v.GetString("default.aws_access_key_id")
	// key := v.GetString("default.ws_secret_access_key")
	// 到这里，完成aws密码从key,password读取

	region := "cn-north-1"

	// 不要在应用程序中嵌入凭据。仅将此方法用于测试目的。
	// config, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(id, key, "")),
	// 	config.WithRegion(region),
	// )

	// will use secret file like ~/.aws/config
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Printf("ERROR config.LoadDefaultConfig : %v\n", err)
	}

	// Create Secrets Manager client
	conn := secretsmanager.NewFromConfig(config)

	// 创建2个值，valueC，valueP
	creatValues(index)

	// 处理
	switch index {
	case 0: // 相当于删除2个
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

	// _ = get(conn, keyP)
	// _ = get(conn, keyC)

}

func handle(conn *secretsmanager.Client, valueStr ...valueStr) {
	// 循环处理，可能有1~2组数据
	for _, v := range valueStr {
		key := v.key
		value := v.value
		if value == "" {
			del(conn, key)035aws
		} else {
			add(conn, key, value)
		}
	}
}

// 增
func add(conn *secretsmanager.Client, path, value string) {
	/*
	   逻辑：
	   查询 --> 有没有？返回
	   删除 --> 无论返回啥，已经删除成功了。
	   增加 --> 可能失败 --> (是存在)改写
	                   --> (是有删残留)恢复+改写
	   修改 --> 没有修改，就是增加 动作
	*/

	if path == "" {
		fmt.Printf("ERROR CreateSecret path == nil \n")
		return
	}

	input := &secretsmanager.CreateSecretInput{
		// Description:  aws.String(""),
		Name:         aws.String(path),
		SecretString: aws.String(value),
	}

	result, err := conn.CreateSecret(context.TODO(), input)
	if err != nil {
		operationError, ok := err.(*smithy.OperationError)
		if ok {
			errString := operationError.Err.Error()
			if strings.HasSuffix(errString, "already scheduled for deletion.") {
				// 存在删除的痕迹，-->恢复-->恢复成功-->修改
				if restore(conn, path) { // 如果 恢复 成功
					putSecretValue(conn, path, value)
					return
				} else {
					return
				}
			} else if strings.HasSuffix(errString, "already exists.") {
				// 存在现在有效的，-->修改
				putSecretValue(conn, path, value)
				return
			}
		}

		// 正常不会走到这里
		fmt.Printf("正常不会走到这里 ERROR CreateSecret :%s | %v\n", path, err)
		return
	}

	fmt.Println("ADD成功 : aws", path, result)
}

// 修改
func putSecretValue(conn *secretsmanager.Client, path, value string) {

	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(path),
		SecretString: aws.String(value),
	}
	_, err := conn.PutSecretValue(context.TODO(), input)
	if err != nil {
		fmt.Println("PutSecretValueInput aws ERROR: ", err.Error())
		return
	}
}

// 恢复
func restore(conn *secretsmanager.Client, path string) bool {
	input := &secretsmanager.RestoreSecretInput{
		SecretId: aws.String(path),
	}

	_, err := conn.RestoreSecret(context.TODO(), input)
	if err != nil {
		fmt.Println("恢复 aws ERROR: ", err.Error())
		return false
	}

	fmt.Println("恢复 aws 成功 ：", path)
	return true
}

// 删除
func del(conn *secretsmanager.Client, path string) {
	input := &secretsmanager.DeleteSecretInput{
		RecoveryWindowInDays: aws.Int64(7),
		SecretId:             aws.String(path),
	}
	result, err := conn.DeleteSecret(context.TODO(), input)
	if err != nil {
		fmt.Printf("ERROR conn.DeleteSecret :%s | %v\n", path, err)
	}
	fmt.Println("删除成功 : aws", result)
}

// 查,返回是否存在。别的子程序有用到这个bool
func get(conn *secretsmanager.Client, path string) bool {

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(path),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	result, err := conn.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		fmt.Printf("ERROR conn.GetSecretValue :%s | %v\n", path, err)
		return false
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	fmt.Printf("取回 aws KV %s : %s\n", path, secretString)
	return true
}

func creatValues(index int) {
	switch index {
	case 0:
		valueC = ""
		valueP = ""
	case 1:
		valueC = `
[[inputs.mem]]
  ##(optional) collect interval, default is 10 seconds
  interval = '11s'

[inputs.mem.tags]
  # some_tag = "some_value"
  # more_tag = "some_other_value"`
		valueP = `add_key(taws, 1)`
	case 2:
		valueC = `
[[inputs.mem]]
  ##(optional) collect interval, default is 10 seconds
  interval = '12s'

[inputs.mem.tags]
  # some_tag = "some_value"
  # more_tag = "some_other_value"`
		valueP = `add_key(taws, 2)`
	case 3:
		valueP = `add_key(taws, 3)`
	case 4:
		valueC = `
[[inputs.mem]]
  ##(optional) collect interval, default is 10 seconds
  interval = '14s'

[inputs.mem.tags]
  # some_tag = "some_value"
  # more_tag = "some_other_value"`
	default:
	}
}
