package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar las contraseñas*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 // costo, es la el numero que coloquemos nosotros, basado en 2 elevado al costo
	//usuario normal 6 y si es mucha la seguraidad necesaria 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	// [] slice de algun tipo es un vector sin cantidad de elementos

	return string(bytes), err
}
