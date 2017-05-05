package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/daisuzu/example-grpc-web-client-js/pb"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("./certs/cert.pem", "")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	cc, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
	}
	cli := pb.NewEchoClient(cc)

	out, err := cli.Call(context.Background(), &pb.Request{Value: "Hello gRPC"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
