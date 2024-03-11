package main

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger/v3"
)

func main() {
	// Open the Badger database located in the badger directory.
	// It will be created if it doesn't exist.
	db, err := Open("./badger")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set a key with a value.
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("mykey"), []byte("myvalue"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// for i := 0; i < 9999; i++ {
	// 	Write(db, []byte("nomeX"+fmt.Sprint(i)), []byte(" 高山YYYYx"+fmt.Sprint(i)))
	// }

	// Get the value back out.
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("mykey"))
		if err != nil {
			return err
		}

		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		log.Printf("The value for 'mykey' is: %s\n", val)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// To delete a key, use the same transactional method as setting a key.
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte("mykey"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Open(path string) (*badger.DB, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	opts := badger.DefaultOptions(path)
	opts.Dir = path
	opts.ValueDir = path
	opts.SyncWrites = false
	opts.ValueThreshold = 256
	opts.CompactL0OnClose = true
	// opts.WithInMemory(true)  // 内存模式
	db, err := badger.Open(opts)
	if err != nil {
		log.Println("badger open failed", "path", path, "err", err)
		return nil, err
	}
	return db, nil
}

func Write(db *badger.DB, k, v []byte) error {
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("mykey"), []byte("myvalue"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
