/**
 * Created with IntelliJ IDEA.
 * User: chris
 * Date: 08/09/2013
 * Time: 16:45
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"time"
	"github.com/coreos/go-etcd/etcd"
	"github.com/coreos/etcd/store"
	"strconv"
)

func main() {

	nodename := "chris"

	c := etcd.NewClient()


	ch := make(chan *store.Response)
	//stop := make(chan bool, 100)

	go setLoop("bar", c, nodename)
	go receiver(ch, c)

	c.Watch(nodename, 0, ch, nil)
}


func setLoop(value string, c *etcd.Client, nodename string) {
	time.Sleep(time.Second)
	for i := 0; true; i++ {
		newValue := fmt.Sprintf("%s_%v", value, i)
		c.Set(nodename+"/foo"+strconv.Itoa(i), newValue, 100)
		//time.Sleep(time.Second / 10000)
	}
}

func receiver(channel chan *store.Response, client *etcd.Client) {
	for true {
	    test := <-channel
		//fmt.Printf("%+v\n", test)
		results, err := client.Get("chris")

		if err != nil || results[0].Key != "/foo" || results[0].Value != "bar" {
			if err != nil {
				fmt.Print(err)
			}
			fmt.Printf("hello %s %s %v\n", results[0].Key, results[0].Value, results[0].TTL)
		}

		fmt.Printf("%+v\n\n", test)


	}
}
