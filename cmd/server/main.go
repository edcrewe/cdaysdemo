package main

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"log"
	"net"
	"os"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"embed"

	api_v1 "github.com/edcrewe/cdaysdemo/generated/go/v1"
)

//go:embed *.csv
var EmbedFS embed.FS

type server struct {
	api_v1.DemoServiceServer
}

// GetCSVFile to return a CSV file via gRPC as HttpBody
func (s *server) GetCSVFile(ctx context.Context, req *api_v1.StringMessage) (*httpbody.HttpBody, error) {
	csvData, err := EmbedFS.ReadFile(req.FileName)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{
		ContentType: "text/csv",
		Data:        csvData,
	}, nil
}

// StreamCSVFile to stream large CSV files via HttpBody
func (s *server) StreamCSVFile(req *api_v1.StringMessage, responseStream api_v1.DemoService_StreamCSVFileServer) error {
	f, err := os.Open("cmd/server/" + req.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	buf := make([]byte, 4*1024*1024) // Use 4 MB buffer

	for {
		n, err := r.Read(buf)
		if n > 0 {
			resp := &httpbody.HttpBody{
				ContentType: "text/csv",
				Data:        buf[:n],
			}
			if err := responseStream.Send(resp); err != nil {
				return nil
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
	}
	return nil
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
