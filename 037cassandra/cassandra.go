package a0037cassandra

import (
	"time"
)

// This test assumes that cassandra is running on default port locally and
// that the keySpace called 'test' already exists.

type Sale struct {
	Id         string
	CustomerId string
	SellerId   string
	Price      int
	Created    time.Time
}

// func doCassandra() {

// 	keySpace, err := gocassa.ConnectToKeySpace("not_exist_demo", []string{"127.0.0.1"}, "", "")

// 	if err != nil {
// 		panic(err)
// 	}
// 	salesTable := keySpace.Table("sale", Sale{}, gocassa.Keys{
// 		PartitionKeys: []string{"Id"},
// 	})

// 	// Create the table - we ignore error intentionally
// 	err = salesTable.Create()
// 	fmt.Println(err)
// 	// We insert the first record into our table - yay!
// 	err = salesTable.Set(Sale{
// 		Id:         "sale-1",
// 		CustomerId: "customer-1",
// 		SellerId:   "seller-1",
// 		Price:      42,
// 		Created:    time.Now(),
// 	}).Run()
// 	if err != nil {
// 		panic(err)
// 	}

// 	result := Sale{}
// 	if err := salesTable.Where(gocassa.Eq("Id", "sale-1")).ReadOne(&result).Run(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(result)
// }

// func doCassandra02() {
// 	// connect to the cluster
// 	cluster := gocql.NewCluster("127.0.0.1")
// 	cluster.Keyspace = "demo"
// 	cluster.Consistency = gocql.Quorum
// 	//设置连接池的数量,默认是2个（针对每一个host,都建立起NumConns个连接）
// 	cluster.NumConns = 3

// 	session, _ := cluster.CreateSession()
// 	time.Sleep(1 * time.Second) //Sleep so the fillPool can complete.
// 	fmt.Println(session.Pool.Size())
// 	defer session.Close()

// 	//unlogged batch, 进行批量插入，最好是partition key 一致的情况
// 	t := time.Now()
// 	batch := session.NewBatch(gocql.UnloggedBatch)
// 	for i := 0; i < 100; i++ {
// 		batch.Query(`INSERT INTO bigrow (rowname, iplist) VALUES (?,?)`, fmt.Sprintf("name_%d", i), fmt.Sprintf("ip_%d", i))
// 	}
// 	if err := session.ExecuteBatch(batch); err != nil {
// 		fmt.Println("execute batch:", err)
// 	}
// 	bt := time.Now().Sub(t).Nanoseconds()

// 	t = time.Now()
// 	for i := 0; i < 100; i++ {
// 		session.Query(`INSERT INTO bigrow (rowname, iplist) VALUES (?,?)`, fmt.Sprintf("name_%d", i), fmt.Sprintf("ip_%d", i))
// 	}
// 	nt := time.Now().Sub(t).Nanoseconds()

// 	t = time.Now()
// 	sbatch := session.NewBatch(gocql.UnloggedBatch)
// 	for i := 0; i < 100; i++ {
// 		sbatch.Query(`INSERT INTO bigrow (rowname, iplist) VALUES (?,?)`, "samerow", fmt.Sprintf("ip_%d", i))
// 	}
// 	if err := session.ExecuteBatch(sbatch); err != nil {
// 		fmt.Println("execute batch:", err)
// 	}
// 	sbt := time.Now().Sub(t).Nanoseconds()
// 	fmt.Println("bt:", bt, "sbt:", sbt, "nt:", nt)

// 	//----------out put------------------
// 	// ./rawtest
// 	// bt: 5795593 sbt: 3003774 nt: 261775
// 	//------------------------------------

// 	// insert a tweet
// 	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
// 		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
// 		log.Fatal(err)
// 	}

// 	var id gocql.UUID
// 	var text string

// 	/* Search for a specific set of records whose 'timeline' column matches
// 	 * the value 'me'. The secondary index that we created earlier will be
// 	 * used for optimizing the search */
// 	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
// 		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Tweet:", id, text)

// 	// list all tweets
// 	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
// 	for iter.Scan(&id, &text) {
// 		fmt.Println("Tweet:", id, text)
// 	}
// 	if err := iter.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// 	query := session.Query(`SELECT * FROM bigrow where rowname = ?`, "30")
// 	// query := session.Query(`SELECT * FROM bigrow `)

// 	var m map[string]interface{}
// 	m = make(map[string]interface{}, 10)
// 	err := query.Consistency(gocql.One).MapScan(m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%#v", m)
// }
