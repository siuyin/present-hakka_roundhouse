package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith"
	"google.golang.org/grpc"
)

//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
// 10 OMIT
func main() {
	fmt.Println("hakka roundhouse monolith service 2")
	requestCh := make(chan *pb.SumRequest)
	loadGen(requestCh)                  // 1 // HL
	responseCh := sumService(requestCh) // 2 // HL
	for {
		select {
		case res := <-responseCh:
			fmt.Printf("%v\n", res.Result)
		}
	}
}

// 20 OMIT
// 30 OMIT
func loadGen(requestCh chan *pb.SumRequest) {
	go func() {
		for {
			req := pb.SumRequest{ // 1 // HL
				A: int32(rand.Intn(99)),
				B: int32(rand.Intn(99))}
			fmt.Printf("Requesting %v+%v\n", req.A, req.B)
			requestCh <- &req // 2 // HL
			time.Sleep(time.Second)
		}
	}()
}

// 40 OMIT
// 50 OMIT
func sumService(requestCh chan *pb.SumRequest) chan *pb.SumResponse {
	ch := make(chan *pb.SumResponse)
	// 52 OMIT
	const address = "localhost:50051" // 0 // HL
	go func() {
		for {

			req := <-requestCh // 1 // HL
			// Set up a connection to the server.
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Printf("did not connect: %v", err)
				continue
			}
			defer conn.Close()

			c := pb.NewArithClient(conn) // 2  // HL
			// Contact the server and print out its response.
			res, err := c.Sum(context.Background(), req) // HL
			if err != nil {
				log.Printf("could not compute sum: %v", err)
				continue
			}

			ch <- res // 3 // HL
		}
	}()
	// 53 OMIT
	return ch
}

// 60 OMIT
