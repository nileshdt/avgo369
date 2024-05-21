// Package main is originally generated by trpc-cmdline v1.0.8.
// It is located at `project/cmd/client`.
// Run this file by executing `go run cmd/client/main.go` in the project directory.
package main

import (
	pb "github.com/nileshdt/pb/logs"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func callLogServiceLog() {
	proxy := pb.NewLogServiceClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	// Example usage of unary client.
	reply, err := proxy.Log(ctx, &pb.LogRequest{App: "app", Name: "level", Data: "message123"})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func main() {
	// Load configuration following the logic in trpc.NewServer.
	cfg, err := trpc.LoadConfig(trpc.ServerConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	trpc.SetGlobalConfig(cfg)
	if err := trpc.Setup(cfg); err != nil {
		panic("setup plugin fail: " + err.Error())
	}
	callLogServiceLog()
}
