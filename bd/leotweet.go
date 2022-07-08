package bd

import (
	"context"
	"log"
	"time"

	"github.com/pedluy/twitteando/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoTweets devueve un slice, porqu eno devuelvo un único valor sino un array de mensajes*/
func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("Twitteando")
	col := db.Collection("tweet")
	var resultados []*models.DevuelvoTweets
	/* Creo la condición por la que busco en la base de datos*/
	condicion := bson.M{
		"usserid": ID,
	}
	/* Ahora creamos las opciones en mongo para el comportamiento en mi db*/
	opciones := options.Find()
	/*con SetLimit le digo cuantos tweets quiero que me pase*/
	opciones.SetLimit(20)
	/*con SetSort cómo va a ir ordenado, donde Key es el valor por el que voy a ordenar y value -1 es en orden descendente*/
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	/*con SetSkip voy a decir cuantos registro voy a obviar cada vez que realice la consulta*/
	opciones.SetSkip((pagina - 1) * 20)

	/*Vamos a crear un puntero (como si fuera una tabla de db), donde vamos a grabar donde vamosa a ir recorriendo de 1 en uno y procesandolos*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	/* Con el bucle for vamos a buscar recorrer todo con el cursor y no parar hasta que tengamos todo lo que nos interesa
	Para esto le ddebo pasar un contexto, nos vamos a crear uno vacio y distento a ctx */
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true

}
