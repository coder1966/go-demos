package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
net.Listen("tcp", "127.0.0.1:0")：监听一个未被占用的端口，并返回 Listener。
调用 http.Serve(ln, nil) 启动 http 服务。
使用 http.Get 发起一个 Get 请求，检查返回值是否正确。
尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景。
*/
// 那我们可以创建真实的网络连接进行测试：

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

/*
针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效。

上述的测试用例改写如下：
*/
// test code

func TestConn2(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}

/* 自动生成的 */
func Test_helloHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHealthCheckHandler01(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	rr := httptest.NewRecorder()

	//直接使用HealthCheckHandler，传入参数rr,req
	HealthCheckHandler(rr, req)

	// 检测返回的状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检测返回的数据
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				w: httptest.NewRecorder(),
				r: mockResponseWriter(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HealthCheckHandler(tt.args.w, tt.args.r)

			// 检测返回的状态码
			assert.Equal(t, tt.args.w.Code, http.StatusOK)

			// 检测返回的数据
			assert.Equal(t, tt.args.w.Body.String(), `{"alive": true}`)
		})
	}
}

func mockResponseWriter() *http.Request {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		return nil
	}
	return req
}
