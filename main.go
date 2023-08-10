package main

import (
	"context"
	"os"

	"github.com/Auladeproyectos/TwitterGo/awsgo"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutarLambda())
}

func EjecutarLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !ValidoParametro() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno, debe incluir el SecretName,BucketName, UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

}

func ValidoParametro() bool {
	_, traeparametro := os.LookupEnv("SecretName")
	if !traeparametro {
		return traeparametro
	}

	_, traeparametro = os.LookupEnv("BucketName")
	if !traeparametro {
		return traeparametro
	}

	_, traeparametro = os.LookupEnv("UrlPrefix")
	if !traeparametro {
		return traeparametro
	}

	return traeparametro

}
