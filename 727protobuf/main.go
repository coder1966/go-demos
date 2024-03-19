package main

import (
	"fmt"
	"log"

	"proto/person"

	"google.golang.org/protobuf/proto"
)

func main() {

	func2()
}

func func2() {

	li4 := &person.Person{Name: "li-si", Id: 18, Detail: []byte("Hello World")}
	// encode ============
	// 序列化成 protobuf 二进制格式
	b, err := proto.Marshal(li4)
	if err != nil {
		log.Fatal("Marshaling error:", err)
	} else {
		fmt.Println("encode GOT: ", string(string(b)))
	}

	// decode ===========
	// 反序列化二进制数据到 BinaryData 实例
	var li5 person.Person
	if err := proto.Unmarshal(b, &li5); err != nil {
		log.Fatal("Unmarshaling error:", err)
	}

	// 使用解码后的数据
	// receivedData := decodedData.Data
	fmt.Println("decode GOT: ", li5)
}

// var binaryData = []byte{0x00, 0xFF, 0xAB, 0xCD}
