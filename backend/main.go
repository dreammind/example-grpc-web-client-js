package main

import (
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
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

	ws := grpcweb.WrapServer(gs)

	mux := http.NewServeMux()
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend"))))
	mux.Handle("/", http.HandlerFunc(ws.ServeHTTP))
	hs := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("Starting server on", hs.Addr)
	log.Println("Please access https://localhost:9090/frontend/")
	log.Println(hs.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem"))
}
