package controladores

import (
	"api/bd"
	"api/modelos"
	"encoding/json"
	"net/http"
)

func GetLoginEndPoint(w http.ResponseWriter, r *http.Request) {
	var usuario modelos.Usuarios
	json.NewDecoder(r.Body).Decode(&usuario)
	json.NewEncoder(w).Encode(bd.GetUsuario(usuario))
}
func PostUsuarioEndPoint(w http.ResponseWriter, r *http.Request) {
	var usuario modelos.Usuarios
	json.NewDecoder(r.Body).Decode(&usuario)
	json.NewEncoder(w).Encode(bd.PostUsuario(usuario))
}
