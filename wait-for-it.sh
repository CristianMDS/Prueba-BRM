#!/bin/sh

TARGET="$1"
shift

echo "Esperando a que MySQL esté disponible en $TARGET..."

# Espera a que el puerto esté abierto
until nc -z $(echo "$TARGET" | cut -d: -f1) $(echo "$TARGET" | cut -d: -f2); do
  sleep 1
done

echo "MySQL está listo."

# Crear la base de datos si no existe
echo "Verificando existencia de la base de datos 'api_go'..."

mysql -h $(echo "$TARGET" | cut -d: -f1) -P $(echo "$TARGET" | cut -d: -f2) -u root -e "CREATE DATABASE IF NOT EXISTS api_go;"

echo "Base de datos 'api_go' verificada o creada."

# Ejecuta la aplicación
exec "$@"
