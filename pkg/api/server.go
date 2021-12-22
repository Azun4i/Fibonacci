package api

import (
	"context"
	"fmt"
	"gobootstrap/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

//protoc -I gobootstrap  gobootstrap/pkg/api/fibonacci_info.proto --go_out=/home/azunai/Desktop/work/go_pj/gobootstrap/pkg

type Server struct {
	Handl *handler.Data
	UnimplementedFibonacciServer
}

func (s *Server) GetFibonacci(ctx context.Context, req *FibonacciRequest) (*FibonacciResponse, error) {
	res, err := s.Handl.Fib(int(req.GetX()), int(req.GetY()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "")
	}

	return &FibonacciResponse{Res: res.NumFib}, nil
}

func Start(grpcS *grpc.Server, grpcServ FibonacciServer, addr *int) {
	//Используя сгенерированные API, регистрируем реализованный ранее сервис на только что созданном gRPC-сервер grpcS, grpcServ interface UnimplementedFibonacciServer
	RegisterFibonacciServer(grpcS, grpcServ)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *addr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcS.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
