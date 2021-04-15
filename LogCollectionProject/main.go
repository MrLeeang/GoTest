package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	elasticv7 "github.com/olivere/elastic/v7"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Sender() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.1.7:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func Consumer() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}

var sample_key string = "123"

func etcdClient() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.197.99:2379"},
		DialTimeout: 25 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// watch
	go func() {
		for {
			fmt.Println("watch sample_key...")
			watchCh := cli.Watch(context.Background(), "sample_key", clientv3.WithPrefix())
			for wresp := range watchCh {
				for _, evt := range wresp.Events {
					fmt.Println(evt.Type.String())
					fmt.Println(string(evt.Kv.Key))
					fmt.Println(string(evt.Kv.Value))
					sample_key = string(evt.Kv.Value)
				}
			}
		}
	}()

	go func() {
		for {
			fmt.Println(sample_key)
			time.Sleep(5 * time.Second)
		}
	}()

	// _, err = cli.Put(ctx, "sample_key", "123")
	// if err != nil {
	// 	panic(err.Error())
	// }

	getresp, err := cli.Get(ctx, "sample_key")
	if err != nil {
		panic(err.Error())
	}
	for _, ev := range getresp.Kvs {
		fmt.Println(string(ev.Key))
		fmt.Println(string(ev.Value))
	}

	getresp, err = cli.Get(ctx, "sample_key", clientv3.WithPrefix())

	for _, ev := range getresp.Kvs {
		fmt.Println(string(ev.Key))
		fmt.Println(string(ev.Value))
	}
}

func esTest() {
	esClient, err := elasticv7.NewClient(elasticv7.SetURL("http://192.168.197.1:9200"))
	if err != nil {
		panic(err.Error())
	}
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Married bool   `json:"married"`
	}

	p := &Person{}
	p.Name = "LIHW"
	p.Age = 18
	p.Married = true

	// 插入数据
	/*
			curl -H "ContentType:application/json" -X POST 127.0.0.1:9200/user/person -d '
		{
			"name": "dsb",
			"age": 9000,
			"married": true
		}'
	*/
	put1, err := esClient.Index().Index("users").Type("person").BodyJson(p).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(put1.Id)
	fmt.Println(put1.Index)
	fmt.Println(put1.Type)
}

func main() {
	// t, err := tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true})
	// for line := range t.Lines {
	// 	fmt.Println(line.Text)
	// }
	esTest()

}
