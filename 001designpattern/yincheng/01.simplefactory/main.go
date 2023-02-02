package main

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	if str == "en" {
		return &English{}
	} else if str == "cn" {
		return &Chinese{}
	}
	return nil
}

type Chinese struct{}

func (*Chinese) Say(name string) string {
	return "您好: " + name
}

type English struct{}

func (*English) Say(name string) string {
	return "hello: " + name
}

func main() {
	api := NewAPI("cn")
	server := api.Say("王大锤")
	fmt.Println(server)

	apiEn := NewAPI("en")
	serverEn := apiEn.Say("alex")
	fmt.Println(serverEn)

}
