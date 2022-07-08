package bd

import (
	"context"
	"time"

	"github.com/pedluy/twitteando/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twitteando")
	col := db.Collection("tweet")

	registro := bson.M{
		"usserid": t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	/* El primitive decuelve el ObjectID del último campo insertado*/
	objID, _ := result.InsertedID.(primitive.ObjectID)
	/*Lo convierto en String, se podría hacer con Hex tb*/
	return objID.String(), true, nil

}
