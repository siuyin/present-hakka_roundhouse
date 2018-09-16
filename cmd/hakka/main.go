package main

import (
	"fmt"

	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith" // 1 // HL
)

//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
func main() {
	fmt.Println("hakka roundhouse monolith")
	res := sum(&pb.SumRequest{A: 2, B: 3}) // 2 // HL
	fmt.Printf("2+3 = %v\n", res.Result)
}

func sum(r *pb.SumRequest) *pb.SumResponse {
	return &pb.SumResponse{Result: r.A + r.B} // 3 // HL
}
