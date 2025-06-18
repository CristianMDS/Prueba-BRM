# 🏗️ Proyecto API REST en Go (Golang)

Programado por: Cristian David Mora Saenz

E-mail: cr.mora00@gmail.com

---

## 📖 Descripción
Esta API REST fue desarrollada como parte de la prueba técnica para **BRM S.A.** Su objetivo principal es permitir la gestión de usuarios a través de operaciones como **crear**, **consultar**, **listar**, **actualizar** y **eliminar** registros.

El servicio está construido utilizando el lenguaje **Golang**, siguiendo una arquitectura limpia y modular. Como motor de persistencia se utiliza **MySQL**, lo cual garantiza un manejo eficiente y escalable de los datos.

La API está diseñada para facilitar la integración mediante herramientas como **Postman** y puede adaptarse fácilmente a distintos entornos de desarrollo o producción.

---
## 🔩 Patrón de diseño (MVC)

Me gusta trabajar con el patrón de diseño **MVC** porque considero que ofrece una estructura robusta y bien definida. Es ampliamente adoptado por diversos frameworks para organizar proyectos, lo que lo convierte en una opción práctica y adaptable a diferentes lenguajes de programación y herramientas de desarrollo.

---

## 🎖️ ¿Qué aprendi?

Durante el desarrollo de este proyecto, tuve la oportunidad de sumergirme por primera vez en el lenguaje **Golang**. Elegí este stack en lugar de optar por algo más familiar como **Node.js**, con el objetivo de ampliar mis conocimientos y explorar nuevas herramientas en el ecosistema backend.

Al principio, comprender la sintaxis de Go y su enfoque estructurado fue un desafío. Sin embargo, a medida que fui avanzando, descubrí una gran satisfacción en su claridad, su rendimiento y, especialmente, en su **flexibilidad para organizar el código de forma modular**. Me pareció muy interesante cómo se pueden integrar librerías para construir funcionalidades de manera progresiva, permitiendo escalar la solución sin perder orden ni control.

Esta experiencia no solo reforzó mi habilidad para aprender tecnologías desde cero, sino que también despertó en mí un interés genuino por seguir explorando todo lo que **Golang** puede ofrecer en el desarrollo de APIs robustas y eficientes.

---

## 🛠️ ¿Qué mejoraria?

Considero fundamental incorporar un módulo de **autenticación** y un sistema de **roles y permisos**, que permita controlar el acceso a ciertos recursos y operaciones. Esta funcionalidad no solo reforzaría la seguridad del sistema, sino que también lo haría más completo y preparado para aplicaciones reales en producción.

---

## 🏁 ¿Qué complementaría?

Me gustaría implementar la creación de un **frontend interactivo**, posiblemente utilizando **React**, una de mis tecnologías favoritas. Esto permitiría mostrar el funcionamiento de la API de manera más visual e intuitiva, facilitando tanto las pruebas como la comprensión por parte de los usuarios.

---

## ⚙️ Tecnologías
| Herramienta | Versión mínima | Uso principal |
|-------------|----------------|---------------|
| Go          | 1.24           | Lógica del servidor |
| Gin         | 1.10.1         | Router HTTP |
| GORM        | 1.30           | Capa de persistencia |

---

## 📝 Requisitos previos
```bash
# Go
go 1.24.4
# Postman
Descargar https://www.postman.com/downloads/
# MySQL
Descargar https://www.apachefriends.org/

```

## 📦 Instalacion
```bash
# 1) Clona el repositorio
git clone https://github.com/CristianMDS/Prueba-BRM.git
cd Prueba-BRM

# 2) Descarga dependencias
go mod tidy

# 3) Ejecuta el servicio
go run main.go
```
---

## 📂 Estructura del proyecto
```bash
.
├── controller/
│   └── actualizar.go       # Controlador para actualizar usuario
│   └── hashPassword.go     # Controlador para encriptar contraseñas
├── database/
|   └── database.go         # Conexion con la base de datos
├── models/
|   └── usuario.go          # Modelo de los Usuarios
├── routes/
|   └── routes.go           # Archivo de enrutamiento
└── main.go                 # Archivo principal
```

---

## ⚙️ Uso de la API
| Método | Endpoint | Descripcion | Body JSON Ejemplo |
|-------------|----------------|---------------|--------------|
| GET          | /listar-usuarios | Listar todos los usuarios | - |
| POST         | /crear-usuario | Crear un usuario | { "Nombre": "string", "Apellido": "string", "Usuario": "string", "Password": "string", "Email": "string", "Contacto": numeric} |
| GET        | /usuario/{id} | Listar un solo usuario usando su {id} | - |
| PUT      | /actualizar-usuario/{id} | Actualizar un usuario con su {id} | { "Nombre": "string", "Apellido": "string", "Usuario": "string", "Password": "string", "Email": "string", "Contacto": numeric} |
| DELETE      | /eliminar-usuario/{id} | Eliminar un usuario con su {id} | - |

---

## 📷 Peticion POST
![Peticion post en postman](img/post.jpg)

## 📷 Peticiones GET
![Peticion post en postman](img/get.jpg)
![Peticion post en postman](img/get_id.jpg)

## 📷 Peticiones PUT
![Peticion post en postman](img/put.jpg)

## 📷 Peticiones DELETE
![Peticion post en postman](img/delete.jpg)
