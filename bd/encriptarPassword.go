package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPasswod(pass string) (string, error) {
	/*Costo es un algoritmo basado en 2 elevado al costo, es el número de pasadas antes de guardar*/
	/*Recomendación: el costo de un superuser meter 8 (min6) en usuario normal poner 6*/
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
