# -------- ETAPA 1: Construcción de la aplicación --------
# Utiliza una imagen ligera de Go basada en Alpine como entorno de build
FROM golang:1.24-alpine AS builder

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de dependencias de Go
COPY go.mod go.sum ./
# Descarga las dependencias declaradas
RUN go mod download

# Copia el resto del código fuente al contenedor
COPY . .

# Compila la aplicación Go en modo estático, optimizada para producción
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app

# Copia el script que espera a que MySQL esté disponible
COPY wait-for-it.sh /wait-for-it.sh
# Otorga permisos de ejecución al script
RUN chmod +x /wait-for-it.sh

# -------- ETAPA 2: Imagen final de producción --------
# Usa una imagen base mínima para ejecución
FROM alpine:latest

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/app .

# Copia el script de espera desde la etapa de construcción
COPY --from=builder /wait-for-it.sh /wait-for-it.sh

# Instala bash para ejecutar el script y el cliente MySQL para crear la base
RUN apk add --no-cache bash
RUN apk add --no-cache mysql-client

# Expone el puerto 8080 donde la app escucha
EXPOSE 8080

# Ejecuta el script que:
# 1. Espera a que MySQL esté disponible
# 2. Crea la base de datos si no existe
# 3. Luego ejecuta el binario Go
CMD ["sh", "-c", "/wait-for-it.sh mysql:3306 -- ./app"]



