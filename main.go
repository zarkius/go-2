package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Usuarios struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

var usuarios = make(map[string]Usuarios)

func manejadorUsuarios(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		obtenerUsuario(w, r)
	case http.MethodPost:
		crearUsuario(w, r)
	case http.MethodPut:
		actualizarUsuario(w, r)
	case http.MethodDelete:
		eliminarUsuario(w, r)
	default:
		// Manejar otros métodos o devolver un error
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}

func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/usuarios"):]
	usuario, existe := usuarios[id]
	if !existe {
		fmt.Printf("No existe el usuario")
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func crearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuarios
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usuarios[usuario.ID] = usuario
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/usuarios"):]
	_, existe := usuarios[id]
	if !existe {
		http.NotFound(w, r)
		return
	}
	var usuario Usuarios
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usuarios[id] = usuario
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/usuarios"):]
	_, existe := usuarios[id]
	if !existe {
		http.NotFound(w, r)
		return
	}
	delete(usuarios, id)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/usuarios", manejadorUsuarios) // Registra el manejador de usuarios una sola vez

	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
