package main

import (
	"flag"
	"fmt"
	"sync"
)

var (
	file   = flag.String("f", "", "-f filename")
	host   = flag.String("h", "", "-h host")
	port   = flag.String("p", "", "-p port")
	schema = "http"
)

type statusCount struct {
	count map[int]int64
	mu    sync.Mutex
}

var sc = statusCount{
	count: make(map[int]int64),
}

func main() {
	flag.Parse()

	go Produce()

	Consume()
	// print result
	fmt.Println("request result:")
	for k, v := range sc.count {
		fmt.Printf("%d: %d\n", k, v)
	}
}
