package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/nileshdt/pb/logs"
	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/nileshdt/pb/logs/logs_mock.go -package=logs -self_package=github.com/nileshdt/pb/logs --source=stub/github.com/nileshdt/pb/logs/logs.trpc.go

func Test_logServiceImpl_Log(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logServiceService := pb.NewMockLogServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := logServiceService.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
		s := &logServiceImpl{}
		return s.Log(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.LogRequest
		rsp *pb.LogResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.LogResponse
			var err error
			if rsp, err = logServiceService.Log(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("logServiceImpl.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("logServiceImpl.Log() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
