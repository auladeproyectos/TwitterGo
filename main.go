package main

import (
	"context"
	"os"
	"strings"

	"github.com/Auladeproyectos/TwitterGo/awsgo"
	"github.com/Auladeproyectos/TwitterGo/bd"
	"github.com/Auladeproyectos/TwitterGo/handlers"
	"github.com/Auladeproyectos/TwitterGo/models"
	"github.com/Auladeproyectos/TwitterGo/secretmanager"
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
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la Lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSing"), SecretModel.JWTsing)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("boday"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Chequeo la conexion a la base de datos
	err = bd.ConectarBd(awsgo.Ctx)

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando en la base de datos " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respApi := handlers.Manejadores(awsgo.Ctx, request)

	if respApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {

		return respApi.CustomResp, nil

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
