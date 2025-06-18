package controller

import (
	"errors" // Paquete utilizado para el manejo y control de errores.
	"golang.org/x/crypto/bcrypt" // Libreria utilizada para encriptar informacioón.
)

// Funcion para encriptar la contraseña del usuario
// 
// Parametro: 
// - pass: Recibe la contraseña.
// 
// Retorna: 
// - string: El hash o la contraseña ya encriptada.
// - error: En caso de fallar retorna el error.
func EncryptPassword(pass string) (string, error){
	// Valida si la variable pass esta vacia
	if pass == "" {
		return "", errors.New("contraseña vacia")
	}

	// Genera el Hash usando bcrypt en caso de tener un error lo retorna
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Retornamos el hash de nuestra contraseña
	return string(hash), nil
}