package main

import (
	"log"
	"os"

	"github.com/m0rk0vka/hex/internal/adapters/app/api"
	"github.com/m0rk0vka/hex/internal/adapters/core/arithmetic"
	"github.com/m0rk0vka/hex/internal/adapters/framework/right/db"

	gRPC "github.com/m0rk0vka/hex/internal/adapters/framework/left/grpc"
	"github.com/m0rk0vka/hex/internal/ports"
)

func main() {
	var err error
	// ports
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var dbaseAdapter ports.DbPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
