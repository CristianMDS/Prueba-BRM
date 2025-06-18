package models

// Definimos la estructura de nuestra tabla "Usuario"
type Usuario struct {
	ID uint `gorm:"primarykey"`
	Nombre string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Usuario string `gorm:"unique" json:"Usuario"`
	Password string `json:"Password"`
	Email string `gorm:"unique" json:"Email"`
	Contacto int `json:"Contacto"`
}