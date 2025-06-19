package config

import (
	"gorm.io/driver/mysql" // Driver para utilizar MySQL en el framework de gorm
	"gorm.io/gorm" // ORM para conexion a base de datos.
	"log" // Paquete para registrar los cambios en el programa.
)

var DB *gorm.DB // Inicializamos la variable DB apuntando a la funcion de gorm

// Creamos la funcion para la conexion a base de datos.
// 
// Funcionamiento:
//  - Inicializamos e instanciamos la variable "conexion" con la informacion de 
//      nuestro servidor.
//  - Generamos la conexion a la base de datos
//  - En caso tener errores los mostramos en nuestro sistema de log.
//  - Crear la base de datos en caso de que no llegara a existir.
//  - Avisar si es posible que la base de datos ya exista.

func Conexion() {
    conexion := "root:@tcp(mysql:3306)/api_go?charset=utf8mb4&parseTime=True"
    var err error
    DB, err = gorm.Open(mysql.Open(conexion), &gorm.Config{}) 
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
}