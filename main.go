package main

import (
	"context"
	"log"
	"time"

	"github.com/BerryTracer/common-service/adapter/database/mongodb"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoDB := mongodb.NewMongoDatabase("mongodb://localhost:27017", "berrytracer", "device")
	if err := mongoDB.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoDB.Disconnect(ctx)

	// indexSpecs := []mongodb.IndexSpec{
	// 	{
	// 		Key:     map[string]interface{}{"fieldname": 1}, // Example index
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// 	// Add more index specs as needed
	// }

}
