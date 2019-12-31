package main

import (
	"context"
	"fmt"

	"github.com/lukasjarosch/genki"
	"github.com/lukasjarosch/genki/cli"
	"github.com/lukasjarosch/genki/config"
	"github.com/lukasjarosch/genki/gateway"
	"github.com/lukasjarosch/genki/logger"
	"github.com/lukasjarosch/genki/server/grpc"
	"github.com/lukasjarosch/genki/server/http"
	"github.com/spf13/pflag"

	pb "github.com/lukasjarosch/secret-santa/api/v1"
	"github.com/lukasjarosch/secret-santa/internal/proto"
	"github.com/lukasjarosch/secret-santa/internal/santa"
)

const (
	Service                   = "secret-santa"
	SecretSantaServiceAddress = "grpc-secret-santa-address"
)

func main() {
	logger.EnsureLoggerFromConfig()
	app := genki.NewApplication(genki.Name(Service))

	// setup mysql datastore

	// setup usecase and middleware
	useCase := santa.NewUseCase()

	// setup gRPC server
	grpcServer := grpc.NewServer(grpc.Name(Service))
	pb.RegisterSecretSantaAPIServer(grpcServer.Server(), proto.NewSecretSantaService(useCase))
	app.AddServer(grpcServer)

	// setup http api gateway server
	apiGateway, closeGatewayContext := setupGateway()
	defer closeGatewayContext()

	// setup http server
	apiServer := http.NewServer(
		http.Name("secret-santa-api"),
		http.LoggingSkipEndpoints("/health"),
	)
	apiServer.Handle("/", apiGateway.HttpMux())
	app.AddServer(apiServer)

	// fire!
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}

func setupGateway() (gateway.Gateway, context.CancelFunc) {
	// the gateway requires a context. once the context cancels, the client connections are also terminated
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	gw := gateway.NewGateway(ctx)

	if err := registerGateway(gw); err != nil {
		logger.Fatal("failed to register gateway: %s", err)
	}

	return gw, cancel
}

func registerGateway(gw gateway.Gateway) error {
	endpoint := config.GetString(SecretSantaServiceAddress)
	err := pb.RegisterSecretSantaAPIHandlerFromEndpoint(gw.Context(), gw.HttpMux(), endpoint, gw.GrpcDialOpts())
	if err != nil {
		return fmt.Errorf("failed to register secret-santa service gateway: %s", err)
	}
	logger.Info("registered http gateway for secret-santa")
	return nil
}

// serviceFlags defines the API specific CLI flags/env-vars
func serviceFlags() *pflag.FlagSet {
	fs := pflag.NewFlagSet(Service, pflag.ContinueOnError)
	fs.String(SecretSantaServiceAddress, "localhost:50051", "gRPC host:port of the secret-santa service")
	return fs
}

// init takes care of setting up the CLI flags, parsing and ultimately binding
// them to the configuration. After they are bound, they are globally accessible via the config package.
func init() {
	flags := cli.NewFlagSet(Service)
	flags.Add(logger.Flags, http.Flags, serviceFlags)
	flags.Parse()
	config.BindFlagSet(flags.Set())
}
