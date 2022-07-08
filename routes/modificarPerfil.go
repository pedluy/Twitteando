package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedluy/twitteando/bd"
	"github.com/pedluy/twitteando/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Internos"+err.Error(), 400)
	}
	/* Recordemos que IDusuario fue creado en el proceso token como global*/
	/*Curiosidad declaro la variable staus y no la inicializo, porque anteriormente ya he dado un valor a err, y no puedo inicializar err dos veces*/
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro.Reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario"+err.Error(), 400)
	}
	w.WriteHeader(http.StatusCreated)
}
