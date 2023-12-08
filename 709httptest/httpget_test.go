package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getData(t *testing.T) {
	tests := []struct {
		name    string
		source  *httptest.Server
		want    string
		wantErr bool
	}{
		{
			name:    "ok",
			source:  mockServer(),
			want:    "{\"code\":200,\"data\":{\"alive\":true},\"msg\":\"ok\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := mockServer()
			defer server.Close()

			got, err := getData(server.URL)
			assert.NoError(t, err)

			assert.Equal(t, got, tt.want)
		})
	}
}

type Response struct {
	Code int64                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func mockServer() *httptest.Server {
	// API调用处理函数
	healthHandler := func(rw http.ResponseWriter, r *http.Request) {
		response := Response{
			Code: 200,
			Msg:  "ok",
			Data: map[string]interface{}{"alive": true},
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(rw).Encode(response)
	}

	// 适配器转换
	return httptest.NewServer(http.HandlerFunc(healthHandler))
}
