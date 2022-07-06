package routes

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pedluy/twitteando/bd"
	"github.com/pedluy/twitteando/models"
)

var Email string
var IDUsuario string

/* Buenas practicas de GO: si una función devuelve varios parámetros y uno es un error, el error al final*/
/* Proceso Token se usa para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("pedluyToken")
	/*Para checkear el JWT necesitamos que sea un puntero. Preguntar a Dani*/
	claims := &models.Claim{}
	/*En el tutorial el token viener con bearer antes del token a mí no. Lo pongo pero debo quitarlo si no funciona*/
	/*Con Split lo divido en dos vectores la cadena. Uno hasta la palabra bearer y otro después (el token)*/
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		/*Para crear un error personalizado ponemos errors.New y el mensaje lo escribimos sin simbolos ni mayusculas*/
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	/*Grabo el token del vector 1 (recordar que el primer vector es 0), quitando espacios y uso trimspace*/
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token no valido")
	}
	return claims, false, string(""), err
}
