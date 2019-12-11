package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

var fakeMetrics = map[int]string{
	1: "a",
	2: "b",
	3: "c",
}

func main() {
	conn, err := net.Dial("udp", "monit_statsd_exp_1:9125")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30

	log.Println("starting statsd example app")
	for {
		num := rand.Intn(max-min+1) + min
		tagA := fakeMetrics[rand.Intn(3-1+1)+1]
		msgS := fmt.Sprintf("duplicate_message.%s.b.c:%d|c", tagA, num)
		msg := []byte(msgS)

		_, err = conn.Write(msg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
