package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

type Producer interface {
	Produce()
}

var uriCh = make(chan string, 100)
var timeStart time.Time

type DefaultProducer struct{}

var dp = DefaultProducer{}

func (DefaultProducer) Produce() {
	fd, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	defer close(uriCh)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		t, uri, err := parseText(scanner.Text())
		if err == nil {
			send(*t, uri)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Produce() {
	dp.Produce()
}

func send(t time.Time, uri string) {
	if reflect.DeepEqual(timeStart, time.Time{}) {
		timeStart = t
	}

	select {
	case <-time.After(t.Sub(timeStart)):
		uriCh <- uri
	}
}

func parseText(rawText string) (*time.Time, string, error) {
	text := strings.Split(rawText, "\t")
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(text[0], " "))
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	return &t, strings.Trim(text[1], " "), nil
}
