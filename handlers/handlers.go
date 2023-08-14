package handlers

import (
	"context"
	"fmt"

	"github.com/Auladeproyectos/TwitterGo/jwt"
	"github.com/Auladeproyectos/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi

	r.Status = 400

	isOK, statusCode, msg, claim := validoAuthorization(ctx, request)

	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {

	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//

	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	}
	r.Message = "Method Ivalid"
	return r

}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obteneravatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}
	claim, todoOK, msg, err := jwt.ProcesosToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el Token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token Ok")
	return true, 200, msg, *claim

}