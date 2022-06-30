package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Esta primera variable vamos a exportarla a toda las partes del código empieza por mayuscula*/
var MongoCN = conectarBD()

/*esta variable se trae la base de datos*/
var clientOptions = options.Client().ApplyURI("mongodb+srv://pedluy:618G7eubAIQWiuZ5@clusterv1.gsp2u.mongodb.net/?retryWrites=true&w=majority")

func conectarBD() *mongo.Client {
	/*Go usa context en memoria para poder trabajar las llamadas, para comunicar y establecer unos valores de setup
	Cuando se mete un "TODO" significa que no vamos a usar ninguna instrucción*/
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		/* Debemos poner .Error la función para que el error se convierta en un string*/
		log.Fatal(err.Error())
		return client
	}
	/*Le hacemos un ping para saber que está la base de datos arriba*/
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexión OK con DB")
	return client
}

func CheckConnection() int {
	/* Es importante que el ping lo hagamos sobre mongo CN que es nuestra variable de conexión*/
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
