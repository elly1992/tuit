package main

/*Backend de una replica de twitter*/

import (
	"log"

	"github.com/elly1992/tuit/bd"
	"github.com/elly1992/tuit/handlers" //Esto es para identificar nuestro paquetes, ya go sigue una estructura de carpetas, previamente establecida.
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
