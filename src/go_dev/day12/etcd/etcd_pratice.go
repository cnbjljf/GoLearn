// etcd_pratice
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed,err", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Get(ctx, "/logger/conf/")
	cancel()
	if err != nil {
		fmt.Println("put failed,err", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logger/conf/")
	cancel()
	if err != nil {
		fmt.Println("get Failed,err", err)
		return
	}
	fmt.Printf("get the value:%T", resp.Kvs)
}
