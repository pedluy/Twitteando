package middlew

import (
	"net/http"

	"github.com/pedluy/twitteando/bd"
)

/* Importante los middleware lo que reciben y lo que suelta siempre deben ser iguales*/
/* next es una variable que se puede definir como siguiente paso*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
