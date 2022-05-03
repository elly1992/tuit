package middlew

import (
	"net/http"

	"github.com/elly1992/tuit/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	/*
		Los middlewares tienen que recibir algo y enviar el mismo tipo de dato, en este caso recibe una una funcion y devuelve una funcion
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 { //si no hay conexion a la BD devuelve el error
			http.Error(w, "Conexión perdida con la Base de Datos", 500) //El error devuelve el mensaje y un status, el error 500
			return                                                      //Si da error con el return automaticamente lo saca y no continua
		}
		next.ServeHTTP(w, r) //si hay conexion devuelve la funcion con todos los objetos de Responwriter w y request r y pasa la informacion al siguiente paso
	}
}
