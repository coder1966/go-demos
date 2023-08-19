package person

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var personResponse = []Person{
	{
		Name:    "wahaha",
		Address: "shanghai",
		Age:     20,
	},
	{
		Name:    "lebaishi",
		Address: "shanghai",
		Age:     10,
	},
}

var personResponseBytes, _ = json.Marshal(personResponse)

func TestPublishWrongResponseStatus(t *testing.T) {
	// >读请求设置通过变量r *http.Request，写变量（也就是返回值）通过w http.ResponseWriter
	// >通过ts.URL来获取请求的URL（一般都是<http://ip:port>）
	// >通过r.Method来获取请求的方法，来测试判断我们的请求方法是否正确
	// >获取请求路径：r.URL.EscapedPath()，本例中的请求路径就是"/person"
	// >获取请求参数：r.ParseForm，r.Form.Get("addr")
	// >设置返回的状态码：w.WriteHeader(http.StatusOK)
	// >设置返回的内容（这就是我们想要的结果）：w.Write(personResponseBytes)，
	//  注意w.Write()接收的参数是[]byte，因此需要将object对象列表通过json.Marshal(personResponse)转换成字节。
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		switch r.URL.EscapedPath() {
		// case "/pools/default":
		// 	w.Write([]byte(pools_default))

		// case "/pools/default/buckets":
		// 	w.Write([]byte(pools_default_buckets))

		// case "/pools/default/tasks":
		// 	w.Write([]byte(pools_default_tasks))
		// default:
		// 	t.Errorf("r.URL.EscapedPath error', got: '%s'", r.URL.EscapedPath())
		// }
		w.Write(personResponseBytes)
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}
	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, _ := GetInfo(api)

	fmt.Println("reps:", resp)
}
