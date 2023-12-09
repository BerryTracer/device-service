package main

import (
	"context"
	"log"
	"time"

	auth_service "github.com/BerryTracer/auth-service/grpc/proto"
	"github.com/BerryTracer/common-service/adapter"
	"github.com/BerryTracer/common-service/adapter/database/mongodb"
	"github.com/BerryTracer/device-service/grpc/server"
	"github.com/BerryTracer/device-service/repository"
	"github.com/BerryTracer/device-service/service"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn.Close()

	authServiceClient := auth_service.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoDB := mongodb.NewMongoDatabase("mongodb://localhost:27017", "berrytracer", "device")
	if err := mongoDB.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoDB.Disconnect(ctx)

	indexSpecs := []mongodb.IndexSpec{
		{
			Key:     map[string]interface{}{"serial_number": 1}, // Example index
			Options: options.Index().SetUnique(true),
		},
		// Add more index specs as needed
	}

	if err := mongoDB.CreateIndexes(ctx, indexSpecs); err != nil {
		log.Fatalf("failed to create indexes: %v", err)
	}

	mongoDBAdapter := adapter.NewMongoAdapter(mongoDB.GetCollection())
	deviceRepository := repository.NewDeviceMongoRepository(mongoDBAdapter)
	deviceService := service.NewDeviceService(deviceRepository)

	server.NewDeviceGrpcServer(deviceService, authServiceClient).Run(":50053")
}
