package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type MqInterface interface {
	NewClient() error
	Subscribe(msgSub *MsgSub) error
	Publish(msgPub *MsgPub) error
}

type MsgSub struct {
	Topic      string
	GroupId    string
	ClientName string
	MsgHandle  func(msgBytes []byte)
}

type MsgPub struct {
	Topic string
	Data  []byte
}

type Client struct {
	Name        string
	Server      []string
	ClientID    string
	User        string
	Password    string
	Config      *sarama.Config
	KafkaClient sarama.SyncProducer
}

func (c *Client) NewClient() error {
	config := sarama.NewConfig()
	//指定 Kafka 版本，选择和购买的 CKafka 相对应的版本，如果不指定，sarama 会使用最低支持的版本
	config.Version = sarama.V1_1_0_0
	config.Net.SASL.Enable = true
	config.Net.SASL.User = c.User
	config.Net.SASL.Password = c.Password
	c.Config = config

	config.Producer.Return.Successes = true
	kafkaClient, err := sarama.NewSyncProducer(c.Server, config)
	if err != nil {
		log.Println("create kafka async producer failed: ", err)
		return err
	}
	c.KafkaClient = kafkaClient
	return nil
}

func (c *Client) Subscribe(msgSub *MsgSub) error {
	//consumer group
	consumer := Consumer{
		MsgHandle: msgSub.MsgHandle,
		ready:     make(chan bool),
	}
	ctx, _ := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(c.Server, msgSub.GroupId, c.Config)
	if err != nil {
		fmt.Printf("%s Error creating consumer group client: %v", msgSub.ClientName, err)
	}
	//defer func() {
	//	cancel()
	//	if err = client.Close(); err != nil {
	//		loge.Error("Error closing client: %v", err)
	//	}
	//}()

	go func() {
		for {
			//Consume 需要在一个无限循环中调用，当重平衡发生的时候，需要重新创建 consumer session 来获得新 ConsumeClaim
			if err := client.Consume(ctx, []string{msgSub.Topic}, &consumer); err != nil {
				fmt.Printf(fmt.Sprintf("Error from consumer:%q err:%s", msgSub.Topic, err))
				time.Sleep(100 * time.Millisecond)
			}
			//如果 context 设置为取消，则直接退出
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	fmt.Printf("Sarama consumer up and running!...")
	return nil
}

func (c *Client) Publish(msgPub *MsgPub) error {

	producerMessage := &sarama.ProducerMessage{
		Topic:     msgPub.Topic,
		Key:       sarama.StringEncoder(c.ClientID),
		Value:     sarama.StringEncoder(msgPub.Data),
		Timestamp: time.Now(),
	}

	if _, _, err := c.KafkaClient.SendMessage(producerMessage); err != nil {
		return err
	}
	return nil
}

//Consumer 消费者结构体
type Consumer struct {
	MsgHandle func(msgBytes []byte)
	ready     chan bool
}

//Setup 函数会在创建新的 consumer session 的时候调用，调用时期发生在 ConsumeClaim 调用前
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

//Cleanup 函数会在所有的 ConsumeClaim 协程退出后被调用
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim 是实际处理消息的函数
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 注意:
	// 不要使用协程启动以下代码.
	// ConsumeClaim 会自己拉起协程，具体行为见源码:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		consumer.MsgHandle(message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}
