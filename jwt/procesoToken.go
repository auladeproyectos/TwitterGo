package jwt

import (
	"errors"
	"strings"

	"github.com/Auladeproyectos/TwitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesosToken(tk string, JWTSing string) (*models.Claim, bool, string, error) {
	miclave := []byte(JWTSing)

	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de Token Invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miclave, nil
	})

	if err == nil {
		//Rutina que chequea contra la base de datos
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Invalido")
	}

	return &claims, false, string(""), err

}
