package a0601streambody

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func Test_streamBody(t *testing.T) {

	tests := []struct {
		name       string
		mockByte   []byte
		wantResult []string
		wantErr    bool
	}{
		{
			name:       "01 testing",
			mockByte:   []byte(mockData),
			wantResult: want,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				Method: http.MethodPut,
				URL: &url.URL{
					Path:     "/prom_remote_write",
					RawQuery: "foo=bar&remoteip=1.2.3.4&__source=2.2.2.2",
				},
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     make(http.Header),
				Host:       "1.1.1.1",
				Body:       io.NopCloser(bytes.NewReader(tt.mockByte)),
			}
			res := httpResponseWriter{}

			gotResult, err := streamBody(res, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("streamBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("streamBody() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

var mockData = `AAAAAAAAAAAAAA
BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC

DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD
EEEEEEEEEEEEEEEEEEEEEE
FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG`

var want = []string{
	"AAAAAAAAAAAAAA\n",
	"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB\n",
	"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC\n",
	"\n",
	"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD\n",
	"EEEEEEEEEEEEEEEEEEEEEE\n",
	"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF\n",
	"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
}

type httpResponseWriter struct{}

func (m httpResponseWriter) Header() http.Header        { return http.Header{} }
func (m httpResponseWriter) Write([]byte) (int, error)  { return 0, nil }
func (m httpResponseWriter) WriteHeader(statusCode int) {}
