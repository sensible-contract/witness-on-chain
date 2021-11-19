package main

import (
	"context"
	"log"
	"witness-on-chain/controller"

	"github.com/gin-gonic/gin"
	ginadapter "github.com/linthan/scf-go-api-proxy/gin"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Println("Gin cold start")
	router := gin.Default()

	router.GET("/info", controller.WitnessOnChain)
	router.GET("/timestamp", controller.SignTimestamp)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	cloudfunction.Start(Handler)
}
