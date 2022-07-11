package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("Twitteando")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id":     objID,
		"usserid": UserID,
	}
	/* Sólo devuelvo el error, porque o se hace o no se hace el borrado de DB, no necesito más*/
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
