package bd

import (
	"context"

	"github.com/Auladeproyectos/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaexisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCn.Database(DatabaseName)
	col := db.Collection("usuario")
	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)

	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
