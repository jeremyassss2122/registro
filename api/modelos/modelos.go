package modelos

type Usuarios struct {
	IdUsuarioApp       int    `json:"idUsuarioApp"`
	Nombre             string `json:"nombre"`
	ApellidoPaterno    string `json:"apellidoPaterno"`
	ApellidoMaterno    string `json:"apellidoMaterno"`
	Telefono           string `json:"telefono"`
	IdPais             string `json:"idPais"`
	Correo             string `json:"correo"`
	Contrasenia        string `json:"contrasenia"`
	FechaRegistro      string `json:"fechaRegistro"`
	FechaTyC           string `json:"fechaTyC"`
	CorreoVerificado   string `json:"correoVerificado"`
	TelefonoVerificado string `json:"telefonoVerificado"`
}
