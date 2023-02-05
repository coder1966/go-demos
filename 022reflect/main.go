package main

/*
	心里要有这个pair<type:*os.File, value:"/dev/tty" 文件描述符>
	是不断传递下来的。

	反射：ValueOf TypeOf
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	// 变量类型对
	doPair()
	// 反射
	doReflect()
	// Tag
	doTag()
	// Tag
	doJson()

}

func doPair() {
	var a string
	// pair<statictype:string, value:"abcde">
	a = "abcde"

	var allType interface{}
	// pair<type:string, value:"abcde">
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
	// =======================================

	// tty: pair<type:*os.File, value:"/dev/tty" 文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println(" error: ", err)
	}

	// pair<type: , value: >
	var r io.Reader

	// pair<type:*os.File, value:"/dev/tty" 文件描述符>
	r = tty

	// pair<type: , value: >
	var w io.Writer

	// pair<type:*os.File, value:"/dev/tty" 文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("Hello "))

	// =================================

	// pair<type:*Book, value:Book{} 的地址>
	b := &Book{}

	// pair<type: , value: >
	var rr Reader
	// pair<type:*Book, value:Book{} 的地址>
	rr = b
	rr.ReadBook()

	var ww Writer
	// pair<type:*Book, value:Book{} 的地址>
	ww = rr.(Writer) // 断言成功，是因为 ww rr 具体类型都是 Book，是一致的
	ww.WriteBook()
}

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}
type Book struct{}

func (b *Book) ReadBook() {
	fmt.Println("实现接口 Reader 的 ReadBook()")
}
func (b *Book) WriteBook() {
	fmt.Println("实现接口 Writer 的 WriteBook()")
}

// ====================================
// 反射
// ====================================
func doReflect() {
	var num float64 = 1.2345
	reflectNum(num)
}
func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))

	// =======================================

	u := User{1, "张三", 57}

	DoFiledAndMethod(u)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Printf("User的方法 user is called ...")
	fmt.Printf("%v \n", u)
}

func DoFiledAndMethod(arg interface{}) {
	// 获取type
	argType := reflect.TypeOf(arg)
	fmt.Println("argType: ", argType.Name())
	// 获取value
	argValue := reflect.ValueOf(arg)
	fmt.Println("argValue: ", argValue)

	// 通过type获取里面的字段
	// 1.获取 reflect.Type，通过 Type得到字段总数NumFiled，然后遍历
	// 2.得到每个Filed的数据类型
	// 3.通过Filed有一个Interface()方法的到value
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argValue.Field(i).Interface()
		fmt.Println("类型合集+字段名+字段类型+值：", field, field.Name, field.Type, value)
	}

	// 通过type获取里面的方法
	for i := 0; i < argType.NumMethod(); i++ {
		m := argType.Method(i)
		fmt.Println("方法 ", m.Name, m.Type)
	}
}

// ====================================
// Tag
// ====================================
func doTag() {
	var re resume
	findTag(&re)
}

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem() // Elem()表示当前结构体全部的元素

	for i := 0; i < t.NumField(); i++ {
		fmt.Println("Tag info: ", t.Field(i).Tag.Get("info"))
		fmt.Println("Tag doc: ", t.Field(i).Tag.Get("doc"))
	}
}

// ====================================
// Json
// ====================================
func doJson() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"星爷", "张柏芝"}}
	j, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println("编码 json.Marshal(movie): ", string(j))

	myMovie := Movie{}
	err = json.Unmarshal(j, &myMovie)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println("解码 json.Unmarshal(j,&myMovie): ", myMovie)
}

type Movie struct {
	Title  string   `json:"name"`
	Year   int      `json:"-"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}
