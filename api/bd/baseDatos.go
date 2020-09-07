package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func NuevaConexionBD() {
	//Se configura la conexion a la base de datos
	db, err = sql.Open("mysql", "bienhechor:Bienhechor_1234;@tcp(74.208.31.248:3306)/bienhechorApp")
	revisarError(err)
	//Se comprueba que la conexion siga activa
	err = db.Ping()
	revisarError(err)
	fmt.Println("Conectado a la BD!!")
}

func TerminarConexionBD() {
	db.Close()
}

func revisarError(err error) {
	if err != nil {
		panic(err)
	}
}
