package config

import (
	"gorm.io/driver/mysql" // Driver para utilizar MySQL en el framework de gorm
	"gorm.io/gorm" // ORM para conexion a base de datos.
	"log" // Paquete para registrar los cambios en el programa.
)

var DB *gorm.DB // Inicializamos la variable DB apuntando a la funcion de gorm

// Creamos la funcion para la conexion a base de datos.
func Conexion() {
    // Inicializamos e instanciamos la variable "conexion" con la informacion de 
    // nuestro servidor local
    conexion := "root:@tcp(127.0.0.1:3306)/api_go?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(conexion), &gorm.Config{}) // Generamos la conexion a la base de datos
    // En caso tener errores los mostramos en nuestro sistema de log.
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
}