package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	//10 OMIT
	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/encoders/protobuf"
	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith"
	//20 OMIT
)

//30 OMIT
//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
func main() {
	fmt.Println("hakka roundhouse pub/sub service")
	sumService() // 1 // HL
	loadGen()    // 2 // HL
	select {}    // wait forever
}

//40 OMIT
//50 OMIT
func loadGen() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("loadGen conn: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Fatalf("loadGen encoded conn: %v", err)
	}

	go func() {
		for {
			req := pb.SumRequest{
				A: int32(rand.Intn(99)),
				B: int32(rand.Intn(99))}
			c.Publish("my-topic", &req) // HL
			fmt.Printf("Published: %v+%v\n", req.A, req.B)
			time.Sleep(time.Second)
		}
	}()
}

//60 OMIT
//70 OMIT
func sumService() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("sum conn: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Fatalf("sum encoded conn: %v", err)
	}
	c.Subscribe("my-topic", func(r *pb.SumRequest) { // HL
		fmt.Printf("%v+%v = %v\n", r.A, r.B, r.A+r.B)
	})
}

//80 OMIT
