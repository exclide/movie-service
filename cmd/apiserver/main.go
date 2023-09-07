package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/exclide/movie-service/internal/app/apiserver"
	"github.com/exclide/movie-service/internal/app/grpcserver"
	"github.com/exclide/movie-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServ := &grpcserver.GrpcServer{}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterReverserServer(s, grpcServ)

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	serv := apiserver.NewServer(config)

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}

}
