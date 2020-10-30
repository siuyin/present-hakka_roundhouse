package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	//10 OMIT
	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/encoders/protobuf"
	"github.com/siuyin/dflt"
	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith"
	//20 OMIT
)

var natsURL string

func init() {
	natsURL = dflt.EnvString("NATS_URL", "nats://localhost:4222")
}

//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
//30 OMIT
func main() {
	fmt.Println("Hakka Roundhouse with NATS pub/sub service")
	sumService() // 1 // HL
	loadGen()    // 2 // HL
	select {}    // wait forever
}

//40 OMIT
func loadGen() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("loadGen conn: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Fatalf("loadGen encoded conn: %v", err)
	}

	//50 OMIT
	go func() {
		for {
			req := pb.SumRequest{
				A: int32(rand.Intn(99)),
				B: int32(rand.Intn(99))}
			if err := c.Publish("my-topic", &req); err != nil { // HL
				log.Printf("could not reach server\n")
				continue
			}
			fmt.Printf("Published: %v+%v\n", req.A, req.B)
			time.Sleep(3 * time.Second)
		}
	}()
	//60 OMIT
}

//70 OMIT
func sumService() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("sum conn: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Fatalf("sum encoded conn: %v", err)
	}
	c.Subscribe("my-topic", func(r *pb.SumRequest) { // HL
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("     added: %v+%v = %v\n", r.A, r.B, r.A+r.B)
	})
}

//80 OMIT
