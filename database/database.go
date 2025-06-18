package config

import (
	"gorm.io/driver/mysql" // Driver para utilizar MySQL en el framework de gorm
	"gorm.io/gorm" // ORM para conexion a base de datos.
	"log" // Paquete para registrar los cambios en el programa.
)

var DB *gorm.DB

func Conexion() {
    conexion := "root:@tcp(127.0.0.1:3306)/api_go?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(conexion), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
}