package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//10 OMIT
	pb "github.com/siuyin/present-hakka_roundhouse/proto/arith" // 1 // HL
)

// 2 go:generate // HL
//go:generate protoc -I ../../proto/arith arith.proto --go_out=plugins=grpc:../../proto/arith
const (
	port = ":50051" // 3 // HL
)

//20 OMIT
//30 OMIT
// server is used to implement ArithServer
type server struct{}

// Sum implements ArithServer.Sum
func (s *server) Sum(ctx context.Context, r *pb.SumRequest) (*pb.SumResponse, error) { // 4 // HL
	return &pb.SumResponse{Result: r.A + r.B}, nil
}

//40 OMIT
//50 OMIT
func main() {
	fmt.Println("GRPC server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArithServer(s, &server{}) // 5 // HL
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//60 OMIT
