package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/stan.go"
)

type SumRequest struct {
	A, B int
}

const clusterID = "test-cluster"

// s OMIT
func main() {
	fmt.Println("Hakka Roundhouse with NATS Streaming service")
	sumService()
	loadGen()
	select {}
}

// e OMIT

func loadGen() {
	go func() {
		clientID := "loadGen"
		sc, err := stan.Connect(clusterID, clientID)
		if err != nil {
			log.Fatal(err)
		}
		defer sc.Close()
		// 10 OMIT
		for {
			req := SumRequest{
				A: rand.Intn(99),
				B: rand.Intn(99),
			}
			b, err := json.Marshal(req)
			if err != nil {
				log.Fatal(err)
			}
			sc.Publish("my-topic", b) // HL
			fmt.Printf("Published: %v+%v\n", req.A, req.B)
			time.Sleep(3 * time.Second)

		}
		// 20 OMIT

	}()

}

func sumService() {
	go func() {
		clientID := "sumService"
		sc, err := stan.Connect(clusterID, clientID)
		if err != nil {
			log.Fatal(err)
		}
		defer sc.Close()

		// 30 OMIT
		sub, err := sc.Subscribe("my-topic", func(m *stan.Msg) {
			var r SumRequest
			err := json.Unmarshal(m.Data, &r)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("     added: %v+%v = %v\n", r.A, r.B, r.A+r.B)
		},
			stan.DurableName("sumService"), // Please remember what I have already received. // HL
			//stan.DeliverAllAvailable(), // HL
		)
		// 40 OMIT
		if err != nil {
			log.Fatal(err)
		}
		defer sub.Close()

		select {}
	}()
}
