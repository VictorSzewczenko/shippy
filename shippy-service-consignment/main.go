// shippy-service-consignment/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/VictorSzewczenko/shippy/shippy-service-consignment/proto/consignment"
	userService "github.com/VictorSzewczenko/shippy/shippy-service-user/proto/user"
	vesselProto "github.com/VictorSzewczenko/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthHandler is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
// func AuthHandler() server.HandlerWrapper {
// 	return func(h server.HandlerFunc) server.HandlerFunc {
// 		return func(ctx context.Context, req server.Request, resp interface{}) error {
// 			meta, ok := metadata.FromContext(ctx)
// 			if !ok {
// 				return errors.New("no auth meta-data found in request")
// 			}

// 			service := micro.NewService(
// 				micro.Name("shippy.service.user"),
// 			)

// 			// Note this is now uppercase (not entirely sure why this is...)
// 			token := meta["Token"]
// 			log.Println("Authenticating with token: ", token)

// 			// Auth here
// 			authClient := userService.NewUserService("shippy.service.user", service.Client())
// 			_, err := authClient.ValidateToken(context.Background(), &userService.Token{
// 				Token: token,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 			err = h(ctx, req, resp)
// 			return err
// 		}
// 	}
// }

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Printf("Running authentication wrapper")

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		// Attempt to extract the token directly from the context meta-data, as this is set directly when calling directly from GRPC.
		token := meta["Token"]
		if len(token) == 0 {
			authorization := meta["Authorization"]
			if len(authorization) == 0 {
				return errors.New("token found in auth meta-data")
			}
			token = strings.TrimPrefix(authorization, "Bearer ")

		}
		log.Printf("Authenticating with token: %s", token)

		// Initiating a service here to get the client is a mistake! For some unknown reason, attempting to init the user service like this in order to get teh client after
		// results in this (the consignment service) service's service registry mechanism breaking, and the service becomes un-reachable after the first service registration TTL expires.
		// service := micro.NewService(
		// 	micro.Name("shippy.service.user"),
		// )

		// Auth here
		authClient := userService.NewUserService("shippy.service.user", client.DefaultClient)
		_, err := authClient.ValidateToken(ctx, &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
