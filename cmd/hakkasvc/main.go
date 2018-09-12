package main

import (
	"fmt"
	"math/rand"
	"time"

	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith"
)

//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
// 10 OMIT
func main() {
	fmt.Println("hakka roundhouse monolith service")
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
	go func() {
		for {
			req := <-requestCh                           // 1 // HL
			ch <- &pb.SumResponse{Result: req.A + req.B} // 2 // HL
		}
	}()
	return ch
}

// 60 OMIT
