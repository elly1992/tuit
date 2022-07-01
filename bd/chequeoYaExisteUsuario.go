package bd

import (
	"context"
	"time"

	"github.com/elly1992/tuit/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya esta en la BD*/

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tuiter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	//la M va a ser como un Map string, va a devolver todo en json
	//el texto lo va a transformar en formato bson, la bd trabaja con ese tipo de archivo

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex() //en el tercer parametro se devolvera este valor de ID
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
