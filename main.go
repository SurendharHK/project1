package main

import (
	"context"
	"fmt"
	"log"
	"project1/config"
	"project1/constants"
	"project1/controllers"
	"project1/routes"
	"project1/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	customerCollection := mongoClient.Database(constants.DatabaseName).Collection("payments")
	
	

	transactionCollection := mongoClient.Database(constants.DatabaseName).Collection("transactions")
	transactionService := services.NewTransactionServiceInit(mongoClient, customerCollection, transactionCollection, ctx)
	transactionController := controllers.InitTransactionController(transactionService)
	routes.TransactionRoutes(server, transactionController)
	
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
