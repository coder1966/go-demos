$ sudo apt  install protobuf-compiler

$ sudo apt  install libprotobuf-dev

$ sudo apt  install protoc-gen-go

创建 .proto

$ protoc --go_out=. *.proto

protoc --go_out=. --go_opt=paths=source_relative person.proto

https://geektutu.com/post/quick-go-protobuf.html

https://protobuf.dev/programming-guides/proto3/