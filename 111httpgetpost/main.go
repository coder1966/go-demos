package main

/*

 */

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	httpGet()
	httpPost()
	fmt.Println("main 结束")
}

func httpGet() {
	urls := []string{
		"https://tenapi.cn/v2/yiyan",
		"https://dog.ceo/api/breeds/image/random",
	}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("get error: ", err)
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll((resp.Body))
		if err != nil {
			fmt.Println(" error: ", err)
			return
		}
		fmt.Println("获得body ： ", string(b))

		// // 解析成 obj
		// obj := NamespaceResponse{}
		// err = json.Unmarshal(b, &obj)
		// if err != nil {
		// 	_ = err
		// }
	}
}

func httpPost() {
	// curl https://tenapi.cn/v2/yiyan -X POST -d 'format=json'
	reqBodyStr := "format=json"
	reqBody := strings.NewReader(reqBodyStr)
	url := "https://tenapi.cn/v2/yiyan"
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Do HTTP.
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	fmt.Println("POST获得body ： ", string(b))

	// // 解析成 obj
	// obj := NamespaceResponse{}
	// err = json.Unmarshal(b, &obj)
	// if err != nil {
	// 	_ = err
	// }
}
