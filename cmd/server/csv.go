package main

import (
	"bufio"
	"context"
	"io"
	"os"

	"log"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_v1 "github.com/edcrewe/cdaysdemo/generated/go/v1"
)

// GetCSVFile returns a CSV file via gRPC as HttpBody
func (s *server) GetCSVFile(ctx context.Context, req *api_v1.StringMessage) (*httpbody.HttpBody, error) {
	csvData, err := os.ReadFile("cmd/server/" + req.FileName)
	if err != nil {
		log.Printf("Error reading file %s: %v", req.FileName, err)
		if os.IsNotExist(err) {
			return nil, status.Errorf(codes.NotFound, "file not found: %s", req.FileName)
		}
		return nil, status.Errorf(codes.Internal, "failed to read file: %v", err)
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
