package bd

import (
	"context"

	"github.com/Auladeproyectos/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCn.Database(DatabaseName)
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
