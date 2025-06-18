package routes

// 
import (
	"net/http" // Paquete para manejar los endpoint de las peticiones.
	"strconv" // Paquete para hacer cast a variables.
	"github.com/gin-gonic/gin" // Framework para crear el enrutamiento.
)


// Estructura del usuario para JSON
type Usuario struct {
	Id int `json:"Id"`
	Nombre string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Usuario string `json:"Usuario"`
	Password string `json:"Password"`
	Email string `json:"Email"`
	Contacto int `json:"Contacto"`
}

func SetRoutes(router *gin.Engine){
	// Ruta para solicitar todo el listado de usuarios.
	router.GET("/listar-usuarios", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Aqui deberia esta la lista de usuarios"})
	})
	
	// Ruta para solicitar un usuario especifico
	router.GET("/usuario/:id", func (c *gin.Context){
		id := c.Param("id")
		Idx, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"user_id": Idx})
	})

	// Ruta para crear un usuario.
	router.POST("/crear-usuario", func (c *gin.Context){
		var user Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario recibido", "data": user})
	})

	// Ruta para actualizar al usuario utilizando el id asignado.
	router.PUT("/actualizar-usuario/:id", func(c *gin.Context){
		id := c.Param("id")
		var user Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		userId, err := strconv.Atoi(id) // Convertir id a int
		if err != nil {
			c.JSON(400, gin.H{"error": "ID debe ser un número válido"})
			return
		}
		if id == ""{
			c.JSON(403, gin.H{"message": "puede que sea por que no hay info"})
		}
		user.Id = userId
		c.JSON(http.StatusOK, gin.H{"message" : "Toca ver como actualizamos", "data": user.Id})
	})

	// Ruta para eliminar el usuario usando su Id.
	router.DELETE("/eliminar-usuario/:id", func (c *gin.Context){
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message" : "Toca ver como Borramos", "data": id})
	})

	// Mensaje para redireccionar en caso de algun error en una ruta.
	router.NoRoute(func(c *gin.Context) {
		c.JSON(406, gin.H{"error": "Ruta no encontrada"})
	})


}