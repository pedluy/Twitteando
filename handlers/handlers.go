package handlers

import (
	"log"
	"net/http"
	"os"

	/*Tengo que llamar a la carpeta del middlew para que lo reconozca*/
	"github.com/gorilla/mux"
	"github.com/pedluy/twitteando/middlew"
	"github.com/pedluy/twitteando/routes"
	"github.com/rs/cors"
)

/* manejadores configuro el puerto y el handler para escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	/*Enlazo el handler para apuntar al perfil pero llamo a dos middleware encadenados para checkear el JWT*/
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.VerPerfil))).Methods("POST")

	/* Abrimos el puerto*/
	PORT := os.Getenv("PORT")
	/*Si el puerto es vacio en el sistema, la forzamos con este if*/
	if PORT == "" {
		PORT = "8080"
	}
	/* Los cors son los permisos de accesos que tenemos que hacer para comu
	nicarnos con un API online, aquí le generamos permisos a todo el mundo*/
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
