package main

/*Esto es necesario ponerlo asi porque necesita llegar a la carpeta*/
import (
	"log"

	"github.com/pedluy/twitteando/bd"
	"github.com/pedluy/twitteando/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
