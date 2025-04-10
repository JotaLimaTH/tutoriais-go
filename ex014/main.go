package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ //Esse método executa uma função (segundo parâmetro) caso acessemos o endereço especificado (primeiro parâmetro). 
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Users page!")
	   })
	if err := http.ListenAndServe(":8080", nil); err != nil { // Isso vai funcionar na porta 8080, e se houver erros, vai retornar. ListenAndServe() só retorna em caso de erro.
		fmt.Println("Server error:", err)
	}
}