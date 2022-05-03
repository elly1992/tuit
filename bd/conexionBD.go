package bd //todos los archivos dentro de la carpeta bd pertenecen al paquete bd. El nombre la carpeta raiz es el nombre del paquete.
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
) //Para que una funcion se vea en otros paquetes, el nombre tiene que iniciar con mayuscula.

//grabar nombre, texto dentro del log de ejecucion

var MongoCN = ConectarBD()                                                                                                                         //Se esta ejecutando al conexion a la bd y va a devolver la conexion, esta variable va a ser exportada a todos los archivos que estan dentro de la carpeta BD
var clientOptions = options.Client().ApplyURI("mongodb+srv://ellym:3420bcee@tuiter.ljcmz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority") // esta funcion es de un paquete de mongo y y esta variable es de tipo clientOptions y recibe un parametro que es la url de donde hay que conectarse

func ConectarBD() *mongo.Client {
	//ConectarBD creara la conexion a la BD
	client, err := mongo.Connect(context.TODO(), clientOptions) //dentro del llamado de una api go acreado el contexto y es un espacio en memoria donde se puede compartir cosas, setear un contexto de ejecucion. (dentro del contexto se puede limitar el tiempo de ejecucion de una peticion).
	//El contexto sirve  para comunicar informacion, entre ejecucion y ejecucion y ademas nos permite setear una serie de valores como un timeout

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) //verifica que haya conexion
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

func ChequeoConnection() int { //ChequeoConnection es el ping a la BD
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
