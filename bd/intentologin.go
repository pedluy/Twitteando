package bd

import (
	"github.com/pedluy/twitteando/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin devuelve un bool de si es ok o no y si es Ok te trae el modelo del usuario*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	/* Creamos una variable tipo model. Como vemos no pongo BD antes del paquete porque estamos dentro del mismo paquete.*/
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}
	/* Ahora compruebo que la clave es igual a la de la DB encriptada*/
	/* Creo una variable tipo slice byte*/
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	/*ahora uso una funcion de bcrypt que permita comparar dos password primero mando la encryptada (hash)*/
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
