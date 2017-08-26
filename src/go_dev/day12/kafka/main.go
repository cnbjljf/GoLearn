// main
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("192.168.14.4:9092", ","), nil)
	if err != nil {
		fmt.Println("failed to start consumer:%s", err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_lo")
	if err != nil {
		fmt.Println("Failed to get the list of partition:", err)
		return
	}

	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer %d: %s \n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, offset:%d key:%s value:%s", msg.Partition,
					msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	time.Sleep(time.Hour)
	consumer.Close()
}
