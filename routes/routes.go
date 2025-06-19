package routes

import (
	"app-api-go/controller" // Llamado al controlador para actualizar usuarios.
	"app-api-go/database" // Llamado a la base de datos.
	"app-api-go/models"   // Llamado a los modelos.
	"net/http"            // Paquete para manejar los endpoint de las peticiones.
	"github.com/gin-gonic/gin" // Framework para crear el enrutamiento.
	"fmt"
)

func SetRoutes(router *gin.Engine){

	router.GET("/", func(c *gin.Context){
		fmt.Println("Hola mundo")
	})

	// Ruta para solicitar todo el listado de usuarios.
	// 
	// Ruta:    GET http://127.0.0.1:8080/listar-usuarios
	// Parametros:
	//  - JSON 
	//
	// Comportamiento:
	//  - Obtener listado de datos almacenados.
	//  - Si ocurre un error muestra el tipo y codigo.
	//  - En caso de salir todo correctamente envia la respuesta.
	
	router.GET("/listar-usuarios", func (c *gin.Context){
		var usuarios []models.Usuario

		result := config.DB.Debug().Find(&usuarios)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": usuarios})
	})
	
	// Ruta para solicitar un usuario especifico
	// 
	// Ruta:    GET  http://127.0.0.1:8080/usuario/:id
	// Parametros:
	//  - id (URL)
	//  - JSON 
	//
	// Comportamiento:
	//  - Obtener el valor de id.
	//  - Crear la variable usuario con el modelo del usuario.
	//  - Realizar la conexion a base de datos.
	//  - Buscar el usuario utilizando la funcion Where para filtrar el ID.
	//  - En caso de que ocurra un error muestra el tipo y codigo.
	//  - En caso de salir todo correctamente envia la respuesta.

	router.GET("/usuario/:id", func (c *gin.Context){
		id := c.Param("id")
		
		var usuario models.Usuario
		result := config.DB.Debug().Where("id=?", id).Find(&usuario)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
			return
		}

		c.JSON(200, gin.H{"data": usuario})
	})

	// Ruta para crear un usuario.
	// 
	// Ruta:    POST  http://127.0.0.1:8080/crear-usuario
	// Parametros:
	//  - id (URL)
	//  - body (JSON): 
	// {
	//    "Nombre" : "string",
	//    "Apellido" : "string",
	//    "Usuario" : "string",
	//    "Password" : "string",
	//    "Email" : "string",
	//    "Contacto" : "numeric",
	//  }
	//
	// Comportamiento:
	//  - Obtener el modelo del usuario.
	//  - Validar la petición al crear el usuario.
	//  - Encriptar la contraseña.
	//  - Reemplazar la contraseña con la ya encriptada.
	//  - Conectarse a la base de datos y crear el registro1.
	//  - Validar que se haya creado correctamente.
	//  - En caso de salir todo correctamente envia la respuesta.

	router.POST("/crear-usuario", func (c *gin.Context){
		var user models.Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		hashed, err := controller.EncryptPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo encriptar la contraseña"})
			return
		}
		user.Password = hashed
		
		result := config.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario creado", "data": user})
	})

	// Ruta para actualizar al usuario utilizando el id asignado.
	// 
	// Ruta:    PUT  http://127.0.0.1:8080/usuario/:id
	// Parametros:
	//  - id (URL)
	//  - JSON 
	//
	// Comportamiento:
	//  - Crear la variable con el modelo Usuario.
	//  - Validar que la peticion JSON este correcta y no de errores.
	//  - Enviar la informacion a las actualizaciones a la funcion ( UpdateData ) 
	// 		en la carpeta controller.
	//  - Obtener el valor de id.
	//  - Realizar la conexion a base de datos.
	//  - Buscar el usuario utilizando la funcion Where para filtrar por el ID.
	//  - Actualizar los datos de la peticion en el registro de la base de datos.
	//  - En caso de que ocurra un error muestra el tipo y codigo.
	//  - En caso de salir todo correctamente envia la respuesta.

	router.PUT("/actualizar-usuario/:id", func(c *gin.Context) {
		var input models.Usuario
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updates := controller.UpdateData(input)

		id := c.Param("id")
		result := config.DB.Model(&models.Usuario{}).Where("id = ?", id).Updates(updates)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Usuario actualizado", "Info actualizada": updates})
	})


	// Ruta para eliminar el usuario usando su Id.
	//
	// Ruta:    DELETE  http://127.0.0.1:8080/eliminar-usuario/:id
	// Parametros:
	//  - id (URL)
	//  - JSON 
	//
	// Comportamiento:
	//  - Obtener el el valor de id.
	//  - Crear la variable usuario con el modelo del usuario.
	//  - Realizar la conexion a base de datos.
	//  - Eliminar el registro usando el id.
	//  - En caso de salir todo correctamente envia la respuesta.

	router.DELETE("/eliminar-usuario/:id", func (c *gin.Context){
		id := c.Param("id")
		var usuario models.Usuario
		config.DB.Delete(&usuario, id)
		c.JSON(http.StatusOK, gin.H{"message" : "Usuario eliminado"})
	})

	// Ruta para redireccionar en caso de no encontrar una ruta creada.
	router.NoRoute(func(c *gin.Context) {
		c.JSON(406, gin.H{"error": "Ruta no encontrada"})
	})


}