package utilidades

import (
	"fmt"
	"math/rand"
	"time"
)

func ObtenerFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return fecha
}

func EnviarCorreo() {

}

func GenerarContrasenia() string {
	rand.Seed(time.Now().UnixNano())
	numeros := "0123456789"
	especiales := "/@#$?+%*"
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		numeros + especiales +
		"abcdefghijklmnopqrstuvwxyz"

	length := 10
	res := make([]byte, length)
	res[0] = numeros[rand.Intn(len(numeros))]
	res[1] = especiales[rand.Intn(len(especiales))]
	for i := 2; i < length; i++ {
		res[i] = a[rand.Intn(len(a))]
	}
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	cadena := string(res)
	return cadena
}
