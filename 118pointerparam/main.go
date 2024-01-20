package main

func main() {

}

type People interface {
	run(int)
	eat()
}

type People1 interface {
	run(int)
	eat()
}

type Man struct {
	name string
	age  int
}

func (Man) run(_ int) {
	panic("not implemented") // TODO: Implement
}

func (Man) eat() {
	panic("not implemented") // TODO: Implement
}

type Man1 struct {
	name string
	age  int
}

func (Man1) run(_ int) {
	panic("not implemented") // TODO: Implement
}

func (Man1) eat() {
	panic("not implemented") // TODO: Implement
}
