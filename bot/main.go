package main

import (
	"context"
	"fmt"
	"os"

	model "github.com/thefarmersfront/doe-bot/core/types"
    tools "github.com/thefarmersfront/doe-bot/bot/cmd"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	FileName := request.QueryStringParameters["FileName"]
	user := request.QueryStringParameters["user"]

	target := model.Upload{FileName: FileName, Title: ""}
	client := model.Client{User: user}

	tools.UploadFile(target, client)

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	// for dev
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file, %v", err))
		}
		target := model.Upload{FileName: "", Title: ""}
		client := model.Client{User: ""}

		tools.UploadFile(target, client)
	} else {
		lambda.Start(handleRequest)
	}

}