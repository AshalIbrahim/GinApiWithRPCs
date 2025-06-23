package main

import (
	"log"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"google.golang.org/grpc"
	userpb "github.com/AshalIbrahim/ginApi/proto/userpb" // adjust path as needed
)

var grpcClient userpb.UserServiceClient

func initGRPCClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	grpcClient = userpb.NewUserServiceClient(conn)
	log.Println("Connected to gRPC server.")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initGRPCClient()

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", AllUsers)
			v1.POST("/users", createUser)
			v1.PUT("/users/:id", updateUser)
			v1.DELETE("/users/:id", deleteUser)
		}
		v2 := api.Group("/v2")
		{
			v2.GET("/", V2api)
		}
	}

	r.StaticFile("/swagger-ui", "./swagger-ui.html")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}