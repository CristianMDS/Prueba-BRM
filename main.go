package main

import (
	"app-api-go/database" // Llamado a las funciones de base de datos
	"github.com/gin-gonic/gin" // Framework de enrutamiento.
	"app-api-go/routes" // Llamado a las rutas.
	"app-api-go/models"
)


func main(){
	config.Conexion() // Conexion a la base de datos
	config.DB.AutoMigrate(&models.Usuario{}) // Migracion de la tabla "Usuarios"
	
	router := gin.Default()	 // Iniciar paquete de enrutamiento
	routes.SetRoutes(router) // Utilizar las rutas creadas en routes.go
	router.Run(":3001") // Habilitar puerto para iniciar la recepcion a través del enrutamiento
}