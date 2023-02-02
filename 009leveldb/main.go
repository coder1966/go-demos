package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

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
	db, err := leveldb.OpenFile("./leveldb/blockstore.db", nil)
	if err != nil {
		fmt.Printf("leveldb.OpenFile err: %v ", err)
	}
	defer db.Close()

	// 增加
	Write(db, []byte("name5"), []byte(" 山5"))
	Write(db, []byte("name6"), []byte(" 山6"))
	Write(db, []byte("name7"), []byte(" 山7"))
	Write(db, []byte("name1"), []byte(" 大山"))
	Write(db, []byte("name2"), []byte(" 小山"))
	Write(db, []byte("name3"), []byte(" 火山"))
	Write(db, []byte("name4"), []byte(" 高山"))

	Write(db, []byte("nome1"), []byte(" 高山1"))
	Write(db, []byte("nome2"), []byte(" 高山2"))
	Write(db, []byte("nome3"), []byte(" 高山3"))
	Write(db, []byte("nome4"), []byte(" 高山4"))

	Batch(db, []BatchData{
		{[]byte("bame1"), []byte(" 山b1")},
		{[]byte("bame2"), []byte(" 山b2")},
		{[]byte("bame3"), []byte(" 山b3")},
	}...)

	// 读取
	fmt.Println(Get(db, []byte("name1")))

	// 遍历|最后一条|查找然后迭代
	Walk(db)

	// 前缀遍历
	Prefix(db, []byte("name"))

	// 掐段遍历
	Sub(db, []byte("name2"), []byte("name6"))

	// 删除
	Del(db, []byte("name4"))

	// Walk(db)

	//获取db快照
	snapshot, i := db.GetSnapshot()
	fmt.Println(snapshot) //leveldb.Snapshot{22}
	fmt.Println(i)        //<nil>
	//注意: The snapshot must be released after use, by calling Release method.
	//也就是说snapshot在使用之后,必须使用它的Release方法释放!
	snapshot.Release()

	Bloom()
	// u := &User{}
	// u.ID = 1
	// u.Name = "大山子"
	// u.Age = 18
	// u.Sex = "男"

	// db.Put("user-1", u, nil)
}

// 写
func Write(db *leveldb.DB, k, v []byte) error {
	db.Put(k, v, nil)
	return nil
}

// 批量写
func Batch(db *leveldb.DB, bs ...BatchData) error {
	batch := new(leveldb.Batch)
	for _, b := range bs {
		batch.Put(b.K, b.V)
	}

	// 也可以批量删除
	batch.Delete([]byte("baz"))

	err := db.Write(batch, nil)

	return err

}

// 读
func Get(db *leveldb.DB, k []byte) ([]byte, error) {
	ids, err := db.Get(k, nil)
	if err != nil {
		return nil, fmt.Errorf("leveldb db.Get err: %v ", err)
	}
	return ids, nil
}

// 遍历
func Walk(db *leveldb.DB) error {
	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	//  最后一条
	if iter.Last() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	//  查找然后迭代
	fmt.Println("=========查找然后迭代")
	for ok := iter.Seek([]byte("name3")); ok; ok = iter.Next() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	iter.Release()
	err := iter.Error()

	return err
}

// 读取某个前缀的所有KEY数据
func Prefix(db *leveldb.DB, pre []byte) error {
	iter := db.NewIterator(util.BytesPrefix(pre), nil)

	for iter.Next() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	//  最后一条
	if iter.Last() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	return nil
}

// 迭代数据库内容的子集：(前包含后不包含)
func Sub(db *leveldb.DB, start, limit []byte) error {
	iter := db.NewIterator(&util.Range{Start: start, Limit: limit}, nil)

	for iter.Next() {
		fmt.Println(string(iter.Key()) + string(iter.Value()))
	}

	return nil
}

// 删除
func Del(db *leveldb.DB, k []byte) error {
	return db.Delete(k, nil)
}

// 带上bloom过滤器
func Bloom() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	dbBloom, err := leveldb.OpenFile("./leveldbbloom", o)
	if err != nil {
		fmt.Printf("leveldb.OpenFile err: %v ", err)
	}
	defer dbBloom.Close()

}
