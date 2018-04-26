package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type NSQHandlerA struct {
}

func (this *NSQHandlerA) HandleMessage(message *nsq.Message) error {
	log.Println("recv a:", string(message.Body))
	return nil
}

type NSQHandlerB struct {
}

func (this *NSQHandlerB) HandleMessage(message *nsq.Message) error {
	log.Println("recv b:", string(message.Body))
	return nil
}

// func (this *NSQHandler) HandleMessage(message *nsq.Message) error {
// 	log.Println("recv b:", string(message.Body))
// 	return nil
// }

func main() {
	waiter := sync.WaitGroup{}
	waiter.Add(2)

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test", "ch-a", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandlerA{})
		err = consumer.ConnectToNSQLookupd("109.254.2.199:4161")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test", "ch-b", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandlerB{})
		err = consumer.ConnectToNSQLookupd("109.254.2.199:4161")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}
