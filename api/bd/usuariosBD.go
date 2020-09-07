package bd

import (
	"api/modelos"
	"api/utilidades"
	"database/sql"
	"fmt"
)

func GetUsuario(usuario modelos.Usuarios) string {

	var (
		correo      string
		contrasenia string
		res         string
	)

	err := db.QueryRow("SELECT correo, contrasenia FROM UsuarioApp WHERE correo = ? AND contrasenia = ?;", usuario.Correo, usuario.Contrasenia).Scan(&correo, &contrasenia)
	switch err {
	case nil:
		res = "Bienvenido"
	case sql.ErrNoRows:
		res = "Usuario o contraseña incorrectos"
	default:
		fmt.Println(err)
	}

	return res
}

func PostUsuario(usuario modelos.Usuarios) string {

	if datoExiste("telefono", usuario.Telefono) {
		if datoExiste("correo", usuario.Correo) {
			_, err = db.Exec("insert into UsuarioApp (Nombre, ApellidoPaterno, ApellidoMaterno, Telefono, IdPais, Correo, Contrasenia, FechaRegistro, FechaTyC, CorreoVerificado, TelefonoVerificado) values(?,?,?,?,?,?,?,?,?,?,?)",
				usuario.Nombre, usuario.ApellidoPaterno, usuario.ApellidoMaterno, usuario.Telefono,
				usuario.IdPais, usuario.Correo, utilidades.GenerarContrasenia(), utilidades.ObtenerFecha(), utilidades.ObtenerFecha(), nil, nil)
			if err == nil {
				return "Registro exitoso"
			} else {
				return err.Error()
			}
		}
	}

	return "El correo y/o teléfono ya existen"
}

func datoExiste(campo, dato string) bool {
	var (
		res      bool
		obtenido string
	)
	err := db.QueryRow("SELECT "+campo+" FROM UsuarioApp WHERE "+campo+" = ?;", dato).Scan(&obtenido)
	switch err {
	case nil:
		res = false //Existe
	case sql.ErrNoRows:
		res = true //No existe
	default:
		fmt.Println(err)
	}
	return res
}
