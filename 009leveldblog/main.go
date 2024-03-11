package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

/*
  如何知道最后写入的块是什么？（）
  bloom 咋用的，我不想用
*/

// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// 	Sex  string
// }

type BatchData struct {
	K []byte
	V []byte
}

func main() {
	// 打开、创建
	db, err := leveldb.OpenFile("./store", &opt.Options{})
	if err != nil {
		fmt.Printf("leveldb.OpenFile err: %v ", err)
	}
	defer db.Close()

	// 遍历
	idx, sumKey, sumVal, err := Walk(db)
	fmt.Println("运行开始", idx, sumKey, sumVal, err)

	for i := 0; i < 1000; i++ {
		fmt.Println("running index: ", i)
		batchOpt(db, i*10000)
	}

	// 遍历
	idx, sumKey, sumVal, err = Walk(db)
	fmt.Println("运行结束", idx, sumKey, sumVal, err)

}

func batchOpt(db *leveldb.DB, add int) {
	// 写入
	for i := 0; i < 9999; i++ {
		Write(db, []byte(fmt.Sprint(i+add)), []byte(fmt.Sprint(i+add)))
	}

	// 删除
	// for i := 0; i < 9999; i++ {
	// 	Del(db, []byte(fmt.Sprint(i+add)))
	// }
}

// 写
func Write(db *leveldb.DB, k, v []byte) error {
	db.Put(k, v, nil)
	return nil
}

// 删除
func Del(db *leveldb.DB, k []byte) error {
	return db.Delete(k, nil)
}

// 遍历
func Walk(db *leveldb.DB) (idx, sumKey, sumVal int, err error) {
	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		idx++
		k, _ := strconv.Atoi(string(iter.Key()))
		sumKey += k
		v, _ := strconv.Atoi(string(iter.Value()))
		sumVal += v
	}

	iter.Release()
	err = iter.Error()

	return
}
