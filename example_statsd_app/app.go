package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9125")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30

	for i := 0; i < 10000; i++ {
		num := rand.Intn(max-min+1) + min
		msgS := fmt.Sprintf("duplicate_message.a.b.c:%d|c", num)
		msg := []byte(msgS)

		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
