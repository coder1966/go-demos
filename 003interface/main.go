package main

type DemoInter01 interface {
	Get01(...string) (interface{}, error)
}
type DemoInter02 interface {
	Get02(...interface{}) (interface{}, error)
}
type DemoInter03 interface {
	Get03()
	Set03()
	Del03()
}

type User01 struct{}

func (*User01) Get01(...string) (interface{}, error) {
	return "", nil
}

type User02 struct{}

func (*User02) Get02(...interface{}) (interface{}, error) {
	return "", nil
}

type User03 struct{}

var _ DemoInter03 = &User03

func main() {
	var user01 DemoInter01 = &User01{}

	var user02 DemoInter02 = &User02{}
}
