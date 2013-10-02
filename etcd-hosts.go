package main

import (
	"fmt"
	"github.com/coreos/etcd/store"
	"github.com/coreos/go-etcd/etcd"
	"time"
)

type Host struct {
	Hostname  string
	Ipaddress string
	Ttl       int64
}

func main() {

	nodename := "hosts"

	etcdClient := etcd.NewClient()

	ch := make(chan *store.Response)
	hostChannel := make(chan *Host, 5)
	//stop := make(chan bool, 100)

	go receiver(ch, etcdClient, hostChannel)
	go loop(ch, etcdClient, nodename)
	go hostthing(hostChannel)
	etcdClient.Watch(nodename, 0, ch, nil)
}

func loop(channel chan *store.Response, etcdClient *etcd.Client, nodename string) {
	for {
		responses, err := etcdClient.Get(nodename)
		if err != nil {
			//print an error
		} else {
			for _, response := range responses {
				channel <- response
			}
		}
		time.Sleep(time.Second * 2)

	}
}

func receiver(channel chan *store.Response, etcdClient *etcd.Client, hostChannel chan *Host) {
	for true {
		response := <-channel
		fmt.Printf("%+v\n", response)

		host := Host{response.Key, response.Value, response.TTL}
		fmt.Println(host)
		hostChannel <- &host
	}
}

func hostthing(hostChannel chan *Host) {
	hosts := make(map[string]*Host)
	for {
		host := <-hostChannel
		hosts[host.Hostname] = host
		fmt.Println(host)
	}
}
