package main

import (
	"context"
	"flag"
	"gobootstrap/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	// устанавливае соединение c сервером через port
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("нет соединение c сервером %v", err)
	}
	defer conn.Close()

	//Передаем соединение и создаем экземпляр заглушки, который содержит все удаленные методы, доступные на сервере
	c := api.NewFibonacciClient(conn)

	x := 10
	y := 5

	//Создаем объект Context, Он содержит метаданные,объект будет существовать до тех пор, пока запрос не будет обработан
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetFibonacci(ctx, &api.FibonacciRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetFibonacci is successfully generated%v", r.Res)

}
