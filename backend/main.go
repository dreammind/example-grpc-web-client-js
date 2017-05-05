package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/daisuzu/example-grpc-web-client-js/pb"
)

type server struct{}

func (s *server) Call(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println("Call <<<", in)
	return &pb.Response{Value: in.GetValue()}, nil
}

const addr = ":9090"

func main() {
	creds, err := credentials.NewServerTLSFromFile("./certs/cert.pem", "./certs/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	gs := grpc.NewServer(opts...)
	pb.RegisterEchoServer(gs, &server{})

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on", l.Addr())
	log.Println(gs.Serve(l))
}
