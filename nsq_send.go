package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {

	producer, err := nsq.NewProducer("109.254.2.199:4150", nsq.NewConfig())
	if err != nil {
		log.Println("NewProducer:%s", err.Error())
	}

	i := 1
	for {
		if err := producer.Publish("test", []byte(fmt.Sprintf("Hello World %d", i))); err != nil {
			log.Println("Publish:%s", err.Error())
		}

		log.Printf(fmt.Sprintf("Hello World %d", i))

		time.Sleep(time.Second * 2)

		i++
	}

	select {}
}
