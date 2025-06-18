package routes

// 
import (
	"net/http" // Paquete para manejar los endpoint de las peticiones.
	// "strconv" // Paquete para hacer cast a variables.
	"github.com/gin-gonic/gin" // Framework para crear el enrutamiento.
	"app-api-go/models" // Llamado a los modelos
	"app-api-go/database" // Llamado a la base de datos
	"log"
)

func SetRoutes(router *gin.Engine){
	// Ruta para solicitar todo el listado de usuarios.
	router.GET("/listar-usuarios", func (c *gin.Context){
		var usuarios []models.Usuario

		result := config.DB.Debug().Find(&usuarios)
		log.Println("Antes de")
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
			return
		}

		log.Println("Usuarios obtenidos: ", usuarios)
		c.JSON(http.StatusOK, gin.H{"data": usuarios})
	})
	
	// Ruta para solicitar un usuario especifico
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
	router.POST("/crear-usuario", func (c *gin.Context){
		var user models.Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		result := config.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario creado", "data": user})
	})

	// Ruta para actualizar al usuario utilizando el id asignado.
	// router.PUT("/actualizar-usuario/:id", func(c *gin.Context){
	// 	id := c.Param("id")
	// 	var usuario models.Usuario
		
	// 	config.DB.Update(&usuario, id)

	// 	c.JSON(http.StatusOK, gin.H{"message" : "Toca ver como actualizamos", "data": user.Id})
	// })

	// Ruta para eliminar el usuario usando su Id.
	router.DELETE("/eliminar-usuario/:id", func (c *gin.Context){
		id := c.Param("id")
		var usuario models.Usuario
		message := "Usuario "+usuario.Nombre+" eliminado correctamente"
		config.DB.Delete(&usuario, id)
		c.JSON(http.StatusOK, gin.H{"message" : message})
	})

	// Mensaje para redireccionar en caso de algun error en una ruta.
	router.NoRoute(func(c *gin.Context) {
		c.JSON(406, gin.H{"error": "Ruta no encontrada"})
	})


}