package main

import (
	"fmt"
	"go-2/persistencia"
)

func main() {
	asincronia := make(chan int, 1000000) // Canal con buffer de tamaño 50000

	// Enviar valores al canal en una goroutine separada
	go func() {
		for i := 0; i < 10000000; i++ {
			asincronia <- i
			fmt.Printf("Enviado: %d, Buffer Usado/Capacidad: %d/%d\n", i, len(asincronia), cap(asincronia))
		}
		close(asincronia) // Importante cerrar el canal
	}()

	// Recibir valores del canal
	for valor := range asincronia {
		fmt.Println("Recibido:", valor)
	}

	// Dereferencia el puntero para obtener el canal y luego recibe un dato
	dato := <-asincronia

	// Ahora puedes llamar a GuardarDato con el dato recibido
	persistencia.GuardarDato(dato)
	// No es necesario el sleep al final, ya que el rango se bloqueará hasta que el canal se cierre
}
