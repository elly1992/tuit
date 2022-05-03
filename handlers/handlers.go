package handlers

/*  */

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors" //cors va a dar permisos de acceso de api
)

/*Manejadores seteo mi puerto, el Handler y pongo a ecuchar al servidor*/
func Manejadores() {
	/*cuando se llame a la api va a entrar a ala funcion de
	manejadores*/

	router := mux.NewRouter() // el mux captura el http y le va a dar manejo al responsewriter y al request que viene en el llamado de la api y va a ver si en body del llamado hay informacion, si en el header hay info. y va a poder mandar la respuesta al navegador

	PORT := os.Getenv("PORT") /* Vamos a ver si en el sistema operativo ya hay creado un puerto, si esta creado lo trae*/
	if PORT == "" {           // si port es igual a nada yo se la voy a setear
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //va a poner al servido a escuchar el puerto si hay peticiones, y el handler va a traer la info
}
