package main

import (
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type Consumer interface {
	Consume()
}
type DefaultConsumer struct{}

var dc = DefaultConsumer{}

func (DefaultConsumer) Consume() {
	for uri := range uriCh {
		url := schema + "://" + *host + ":" + *port + uri
		wg.Add(1)
		go consume(url)
	}
	wg.Wait()
}

func Consume() {
	dc.Consume()
}

func consume(url string) {
	defer wg.Done()
	c := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	incr(res.StatusCode)
	res.Body.Close()
}

func incr(status int) {
	sc.mu.Lock()
	sc.count[status] += 1
	sc.mu.Unlock()
}
