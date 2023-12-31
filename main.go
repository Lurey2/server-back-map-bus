package main

import (
	"context"
	"rutasMap/v2/conf"
	"rutasMap/v2/routers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	conf.InitializeConfig()
	r := setupRouter()
	ginLambda = ginadapter.New(r)

	lambda.Start(Handler)

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	routers.SetRouter(r)
	return r
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
