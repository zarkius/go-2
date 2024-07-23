// persistencia.go
package persistencia

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct to represent the data
type Data struct {
	Dato int `json:"dato"`
}

// GuardarDato guarda un dato en un archivo JSON
func GuardarDato(dato int) {
	// Create a struct instance with the data
	data := Data{Dato: dato}

	// Convert the struct to JSON
	datoBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir el dato a JSON:", err)
		return
	}

	// Crear o abrir el archivo datos.json
	archivo, err := os.OpenFile("datos.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	// Escribir el JSON al archivo
	_, err = archivo.Write(datoBytes)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// Escribir una nueva línea después del JSON para separar los datos
	_, err = archivo.WriteString("\n")
	if err != nil {
		fmt.Println("Error al escribir nueva línea en el archivo:", err)
		return
	}

	fmt.Printf("Dato %d guardado con éxito en JSON\n", dato)
}
