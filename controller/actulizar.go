package controller

import (
	"reflect" // Reflect es un paquete para inspeccionar y manipular la informacion y los tipos.
	"strings" // Paquete para manipular los strings
)

// UpdateData convierte una estructura en un mapa dinámico listo para actualizar registros.
//
// Solo se incluirán los campos cuyo valor sea diferente al valor cero de su tipo,
// y cuya etiqueta JSON esté definida.
//
// Parámetro:
// - input: estructura con los datos a procesar.
//
// Retorna:
// - map[string]interface{} con los campos y valores válidos para la actualización.
func UpdateData(input interface{}) map[string] interface{}{
	resultado := make(map[string] interface{}) // Contenedor de los campos no vacios
	tag := reflect.ValueOf(input) // Valores actuales de los campos
	tipo := reflect.TypeOf(input) // Tipos y etiquetas de cada campo

	// Recorremos el json
	for i := 0; i < tag.NumField(); i++ {
		nombre := tag.Field(i) // Es el "NameTag" de cada fila
		valor := tipo.Field(i) // Es el valor de cada "NameTag"
		jsonTag := strings.Split(valor.Tag.Get("json"), ",")[0]

		if jsonTag == "-" || jsonTag == ""{
			continue
		}

		// Validamos que las filas no esten vacias y que cuenten con un valor.
		zero := reflect.Zero(nombre.Type()).Interface()
		if !reflect.DeepEqual(nombre.Interface(), zero) {
			resultado[jsonTag] = nombre.Interface()
		}
	}

	// Regresamos un arreglo con la informacion para que el update lo entienda
	return resultado 

}