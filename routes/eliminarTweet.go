package routes

import (
	"net/http"

	"github.com/pedluy/twitteando/bd"
)

/* Es importante que las llamadas a DB y los routes se llamen diferente para no confundir*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe mandar el parámetro ID", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar borrar", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
