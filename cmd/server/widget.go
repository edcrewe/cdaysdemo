package main

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"buf.build/go/protovalidate"
	api_v1 "github.com/edcrewe/cdaysdemo/generated/go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const widgetDb = "widgets.json"

var mu sync.Mutex
var v, _ = protovalidate.New()

// Helper to load widgets from file
func loadWidgets() ([]*api_v1.Widget, error) {
	data, err := os.ReadFile(widgetDb)
	if err != nil {
		if os.IsNotExist(err) {
			return []*api_v1.Widget{}, nil
		}
		return nil, err
	}
	var widgets []*api_v1.Widget
	err = json.Unmarshal(data, &widgets)
	return widgets, err
}

// Helper to save widgets to file
func saveWidgets(widgets []*api_v1.Widget) error {
	data, err := json.MarshalIndent(widgets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(widgetDb, data, 0644)
}

func (s *server) CreateWidget(ctx context.Context, req *api_v1.Widget) (*api_v1.WidgetResponse, error) {
	if err := v.Validate(req); err != nil {
		// protovalidate returns a 400 Bad Request to the client
		return nil, status.Errorf(codes.InvalidArgument, "Validation failed: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()

	widgets, _ := loadWidgets()
	widgets = append(widgets, req)

	if err := saveWidgets(widgets); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save: %v", err)
	}

	return &api_v1.WidgetResponse{
		Message: "Thanks for the update!",
		Widget:  req,
	}, nil
}

func (s *server) ListWidgets(ctx context.Context, _ *emptypb.Empty) (*api_v1.WidgetList, error) {
	widgets, err := loadWidgets()
	if err != nil {
		return nil, err
	}
	return &api_v1.WidgetList{Widgets: widgets}, nil
}

func (s *server) GetWidget(ctx context.Context, req *api_v1.Widget) (*api_v1.Widget, error) {
	widgets, _ := loadWidgets()
	for _, w := range widgets {
		if w.Id == req.Id {
			return w, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "widget %d not found", req.Id)
}

func (s *server) DeleteWidget(ctx context.Context, req *api_v1.Widget) (*emptypb.Empty, error) {
	mu.Lock()
	defer mu.Unlock()

	widgets, _ := loadWidgets()
	newWidgets := []*api_v1.Widget{}
	for _, w := range widgets {
		if w.Id != req.Id {
			newWidgets = append(newWidgets, w)
		}
	}
	saveWidgets(newWidgets)
	return &emptypb.Empty{}, nil
}
