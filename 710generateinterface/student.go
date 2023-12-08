package main

type man interface {
	Say(string)
	Run()
	Get() string
}

type student struct {
	name string
	age  string
}
