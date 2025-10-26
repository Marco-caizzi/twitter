package models

import "github.com/aws/aws-lambda-go/events"

type RespAPI struct {
	StatusCode              int                            `json:"statusCode"`
	Headers                 map[string]string              `json:"headers"`
	Body                    string                         `json:"body"`
	APIGatewayProxyResponse events.APIGatewayProxyResponse `json:"-"`
}
