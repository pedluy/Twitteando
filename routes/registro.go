package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedluy/twitteando/bd"
	"github.com/pedluy/twitteando/models"
)

/*Registro es la funcion para crear el registro en la DB*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	/*El body de un http decoder es un string que se lee una vez y se destruye*/
	/*Primero creo el modelo de usuario y lo decodifico en el t*/
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	/* Vamos a hacer unas validaciones previas en el API, si el largo del mail es 0*/
	if len(t.Email) == 0 {
		http.Error(w, "El mail de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password necesita al menos 6 caracteres", 400)
		return
	}

	/*Ahora vamos a hacer una validación de datos*/
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese mail", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Existe un problema en el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No ha podido lograr insertar el registro", 405)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
