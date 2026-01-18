package main

import (
	"context"
	"fmt"

	"log"
	"net"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api_v1 "github.com/edcrewe/cdaysdemo/generated/go/v1"
)

type server struct {
	api_v1.UnimplementedDemoServiceServer // fixes unimplemented errors
}

// GetWebPage returns a simple HTML page via gRPC as HttpBody
func (s *server) GetWebPage(ctx context.Context, req *api_v1.GetPageRequest) (*httpbody.HttpBody, error) {
	htmlContent := "<html><body><h1>Hello from gRPC!</h1><p>Demo index page transcoded via Envoy for <a href=\"https://github.com/edcrewe/cdaysdemo\">Container Days transcoding Demo repo</a></p></body></html>"

	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte(htmlContent),
	}, nil
}

func main() {
	// Initialize server
	srv := &server{}

	fmt.Println("Starting gRPC server on :9090")
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api_v1.RegisterDemoServiceServer(s, srv)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("gRPC server starting on :9090")

	// RUN THIS IN THE MAIN THREAD (Blocking)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
