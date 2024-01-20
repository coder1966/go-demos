package a0037cassandra

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func read() {
	var passwdAuth gocql.PasswordAuthenticator
	passwdAuth.Username = "username"
	passwdAuth.Password = "passwd"

	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Timeout = 10 * time.Second
	cluster.Keyspace = "replication_factor"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = passwdAuth
	session, _ := cluster.CreateSession()
	defer session.Close()

	beginTime := currentTimeMillis()
	fmt.Println("read begin, read begin time is :", beginTime)

	var countNum int64
	var id gocql.UUID
	var text string

	//iter := session.Query(`SELECT id, text FROM tweet WHERE text = '1000'  limit 10000`).Iter()
	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? and text = '1100' ALLOW FILTERING`, "me").Iter()
	for iter.Scan(&id, &text) {
		countNum++
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Println(err)
	}

	endTime := currentTimeMillis()
	fmt.Println("read end, read end time is :", endTime)
	fmt.Println("use time is :", endTime-beginTime)
	fmt.Println("countNum is :", countNum)
}

func currentTimeMillis() int64 {
	return time.Now().Unix()
}

func demo() {
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "replication_factor2"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
}
