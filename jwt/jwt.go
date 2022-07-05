package jwt

import (
	"time"
	/*Con el formato de texto que ponemos aquí, podemos crear un alias para un paquete*/
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pedluy/twitteando/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	/* JWT trabaja con slice de Byte no con string*/
	miClave := []byte("pedluytoken")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		/*El formato Unix devuelve la decha como un número Long*/
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	/*Creamos la variable token con parámetros primero un metodo de firma*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
