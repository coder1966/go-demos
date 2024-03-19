package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRedisDo(t *testing.T) {
	type args struct {
		// index int
		ip string
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip: "10.100.65.11",
			},
			times: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i := 0; i < tt.times; i++ {
				index := i % 5
				RedisDo(index, tt.args.ip)
				time.Sleep(time.Millisecond * 100)
				t.Log(i)
				fmt.Print(i)
			}

		})
	}
}

func TestRedisD02(t *testing.T) {
	type args struct {
		// index int
		ip string
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip: "127.0.0.1",
			},
			times: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			RedisDo02(tt.times, tt.args.ip)

		})
	}
}

func TestRedisD03(t *testing.T) {
	type args struct {
		// index int
		ip string
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip: "127.0.0.1",
			},
			times: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			RedisDo03(tt.times, tt.args.ip)

		})
	}
}

func TestWriteCluster04(t *testing.T) {
	type args struct {
		// index int
		ip string
		pf int
		pe int
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip: "127.0.0.1",
				pf: 6379,
				pe: 6384,
			},
			times: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			WriteDo04(tt.times, tt.args.ip, tt.args.pf, tt.args.pe)

		})
	}
}

func TestWrite6376D05(t *testing.T) {
	type args struct {
		// index int
		ip   string
		port string
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip:   "127.0.0.1",
				port: "6376",
			},
			times: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			RedisDo05(tt.times, tt.args.ip, tt.args.port)

		})
	}
}
func TestWrite6377D05(t *testing.T) {
	type args struct {
		// index int
		ip   string
		port string
	}
	tests := []struct {
		name  string
		args  args
		times int
	}{
		{
			name: "ok",
			args: args{
				// index:1,
				// ip: "192.168.56.14",
				ip:   "127.0.0.1",
				port: "6377",
			},
			times: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			RedisDo05(tt.times, tt.args.ip, tt.args.port)

		})
	}
}
