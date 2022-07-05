package models

type RespuestaLogin struct {
	/* Colocamos el omitempty porque en el caso de que haya un error que lo devuelva vacio*/
	Token string `json:"token,omitempty"`
}
