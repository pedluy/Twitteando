package bd

import (
	"context"
	"time"

	"github.com/pedluy/twitteando/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	/* ctx es una variable contexto*/
	/* Al contexto le meto que me lo muestre en background y que tarde en dar error 15 segundos*/
	/* defer se aplica como última instrucción y hago que cuando llega el cancel pare de meter datos en background*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitteando")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPasswod(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err == nil {
		return "", false, err
	}
	/* ¿Cómo obtengo el iD?*/
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
