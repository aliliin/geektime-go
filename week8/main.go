package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hhxsv5/go-redis-memory-analysis"
	"time"
)

var (
	rdb *redis.Client
)

// 初始化 reds 连接
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	_ = context.Background()
}

func main() {
	write(10000, "len10-10k", generateVal(10))
	write(50000, "len10-50k", generateVal(10))

	write(10000, "len1000-10k", generateVal(1000))
	write(50000, "len1000-50k", generateVal(1000))

	write(10000, "len5000-10k", generateVal(5000))
	write(10000, "len5000-50k", generateVal(5000))

	analysis()
}

func write(n int, key, value string) {
	for i := 0; i < n; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := rdb.Set(k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func generateVal(size int) string {
	array := make([]byte, size)
	for i := 0; i < size; i++ {
		array[i] = 'A'
	}
	return string(array)
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection("127.0.0.1", 6379, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}

	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
