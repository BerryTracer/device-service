package main

import (
	"context"
	"log"
	"time"

	auth_service "github.com/BerryTracer/auth-service/grpc/proto"
	"github.com/BerryTracer/common-service/adapter/database/mongodb"
	"github.com/BerryTracer/common-service/config"
	"github.com/BerryTracer/device-service/grpc/server"
	"github.com/BerryTracer/device-service/repository"
	"github.com/BerryTracer/device-service/service"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// --- gRPC Client Setup ---
	authAuthServiceURI, err := config.LoadEnv("AUTH_SERVICE_URI")
	if err != nil {
		log.Fatalf("failed to load environment variable: %v", err)
	}

	// Establish a connection to the gRPC server
	conn, err := grpc.Dial(authAuthServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the AuthService
	authServiceClient := auth_service.NewAuthServiceClient(conn)

	// --- MongoDB Setup ---
	// Set a timeout for the MongoDB connection context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Load MongoDB URI from environment variables
	mongodbURI, err := config.LoadEnv("MONGODB_URI")
	if err != nil {
		panic(err)
	}

	// Initialize MongoDB connection
	mongoDB := mongodb.NewMongoDatabase(mongodbURI, "berrytracer", "device")
	if err := mongoDB.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoDB.Disconnect(ctx)

	// Create MongoDB indexes
	indexSpecs := []mongodb.IndexSpec{
		{
			Key:     map[string]interface{}{"serial_number": 1},
			Options: options.Index().SetUnique(true),
		},
	}
	if err := mongoDB.CreateIndexes(ctx, indexSpecs); err != nil {
		log.Fatalf("failed to create indexes: %v", err)
	}

	// --- Repository and Service Initialization ---
	// Initialize the MongoDB adapter for the device repository
	mongoDBAdapter := mongodb.NewMongoAdapter(mongoDB.GetCollection())

	// Set up the device repository with the MongoDB adapter
	deviceRepository := repository.NewDeviceMongoRepository(mongoDBAdapter)

	// Initialize the device service with the device repository
	deviceService := service.NewDeviceService(deviceRepository)

	// --- gRPC Server Initialization ---
	// Start the Device gRPC server
	server.NewDeviceGrpcServer(deviceService, authServiceClient).Run(":50053")
}
