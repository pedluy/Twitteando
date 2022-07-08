package bd

import (
	"context"
	"time"

	"github.com/pedluy/twitteando/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("Twitteando")
	col := db.Collection("usuarios")
	registro := make(map[string]interface{})
	/*Asumimos que cuando mandemos desde el front, se mandan sólo los campos que han sufrido modificacionesconst
	pero claro no podemos borrar los valores que ya existen. Por este motivo metemos la siguiente condición*/
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["Avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["Banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updString := bson.M{
		/* Para actualizar un registro en Mongo, tenemos que hacer a la funcion de mongo $set*/
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	/*Nuestro filtro con nuestra condición*/
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	/*Ahora toca ejecutar la instrucción de mongo*/
	_, err := col.UpdateOne(ctx, filtro, updString)
	if err != nil {
		return false, err
	}
	return true, nil

}
