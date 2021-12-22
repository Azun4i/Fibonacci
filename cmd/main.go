package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"gobootstrap/internal/app/apiserver"
	api "gobootstrap/pkg/api"
	"google.golang.org/grpc"
	"log"
)

var (
	configPath string
	port       = flag.Int("port", 50051, "The server port")
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfir()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("failed to parse the config file %v", err)
	}

	// grpc
	grpcS := grpc.NewServer()
	grpcServ := &api.Server{}

	go api.Start(grpcS, grpcServ, port)
	defer grpcS.Stop()

	//rest
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

//http POST http://localhost:8090/fibonacci? x=10 y=0
