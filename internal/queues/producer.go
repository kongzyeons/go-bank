package queues

import (
	"encoding/json"
	"reflect"

	"github.com/IBM/sarama"
	"github.com/kongzyeons/go-bank/internal/models/events"
)

type EventProducer interface {
	Produce(event events.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return &eventProducer{producer}
}

func (svc *eventProducer) Produce(event events.Event) error {
	topic := reflect.TypeOf(event).Name()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = svc.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil
}
