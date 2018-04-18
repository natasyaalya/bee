package conn

import (
	"log"
	"strconv"

	nsq "github.com/bitly/go-nsq"
)

// Consume ::
func Consume(topic string, chanel string) {

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(topic, chanel, config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		el, err := strconv.Atoi(string(message.Body))
		if err != nil {
			return err
		}
		client := Redis.Get()
		_, err = client.Do("SET", "visitor", el+1)
		if err != nil {
			return err
		}
		return nil
	}))
	err := q.ConnectToNSQD("devel-go.tkpd:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
}
