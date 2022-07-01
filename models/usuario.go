/*Endpoint para el registro de usuarios*/

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB*/
type Usuario struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	//objectID es el id de mongo, es un objeto binario
	//se utiliza la comilla invertida por ser codigo bson
	//los objectID se guardan en _id esa es la convencion de todos los id en mongo
	//el omitempty sirve para que si este campo en un intercambio de datos viene vacio, que lo omita, que no lo tome en cuenta para conformar ningun json que hagamos. Omite si viene una clave valor vacia. y esto es hacia la base antes de conformar el json
	//luego va a devolver el formato json el id en minuscula
	Nombre string `bson:"nombre" json:"nombre,omitempty"`
	// el omitempty se coloco aqui porque si no hubiera nombre en el json se omite el mostrarse en el navegador
	Apellidos       string    `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	//la base de datos no reconoce un tipo de dato DATE entonces lo trae
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password,omitempty"` //siempre va a ser omitempty para el navegador (segundo)porque en el navegador no se debe de mostrar una contraseña
	Avatar    string `bson:"avatar" json:"avatar,omitempty"`
	Banner    string `bson:"banner" json:"banner,omitempty"`
	Biografia string `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb  string `bson:"sitioweb" json:"sitioweb,omitempty"`
}
