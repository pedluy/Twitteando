package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pedluy/twitteando/bd"
)

func LeoTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe envial el parámetro página con un valor mayor a 0", http.StatusBadRequest)
	}
	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweet(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "apllication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
