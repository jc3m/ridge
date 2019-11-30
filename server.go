package main

import (
  "context"
  "fmt"
  "log"
  "net"
  "time"

  pb "github.com/jc3m/ridge/generated"
  "google.golang.org/grpc"
)

// TODO: Read from flag
var port string = "8000"

type server struct {}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
  log.Printf("Received: %v", in.GetName())
  return &pb.HelloResponse{Greeting: "Hello " + in.GetName()}, nil
}

func main() {
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
  if err != nil {
    log.Fatalf("Error listening: %v", err)
  }

  s := grpc.NewServer(
    grpc.ConnectionTimeout(30 * time.Second), // TODO: Make this a global
  )
  pb.RegisterHelloServer(s, &server{})

  log.Printf("Listening...")
  err = s.Serve(lis)
  if err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
