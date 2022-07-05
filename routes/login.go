package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedluy/twitteando/bd"
	"github.com/pedluy/twitteando/jwt"
	"github.com/pedluy/twitteando/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	/*Primero seteamos el header.*/
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El mail es obligatorio"+err.Error(), 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}
	/* Ahora vamos a crear el token para poder comprobar que está en uso*/
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*Vamos a crear una cookie para guardarla en el navegador*/
	/*expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie (w, &http.Cookie{
		Name:"token"
		Value: jwtKey
		Expires: expirationTime
	})*/

}
