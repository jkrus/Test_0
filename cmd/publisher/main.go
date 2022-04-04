package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"

	"wb_L0/pkg/models"
)

func main() {
	path := "pkg/models/order.json"
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("ERR open file", err)
	}

	ord := models.Order{}
	json.Unmarshal(d, &ord)
	nc, _ := nats.Connect(nats.DefaultURL)

	// Connect to NATS
	/*	nc, _ := nats.Connect(nats.DefaultURL)
		for i := 0; i < 1000; i++ {
			ord.OrderUid = strconv.Itoa(i)
			bytes, err := json.Marshal(&ord)
			if err != nil {
				return
			}
			nc.Publish("ORDERS", bytes)
		}

	*/

	// Create JetStream Context
	js, _ := nc.JetStream()

	// Create a Stream
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDER"},
	})

	// Simple Stream Publisher
	id := ord.OrderUid
	for i := 0; i < 10000; i++ {
		ord.OrderUid = id + strconv.Itoa(i)
		bytes, err := json.Marshal(&ord)
		if err != nil {
			log.Println(err)
			return
		}
		js.Publish("ORDERS", bytes)
	}

	/*	type badOrder struct {
		}
		o := badOrder{}
		for i := 0; i < 1000; i++ {
			bytes, err := json.Marshal(&o)
			if err != nil {
				log.Println(err)
				return
			}
			js.Publish("ORDERS", bytes)
		}*/

	defer nc.Close()
}
