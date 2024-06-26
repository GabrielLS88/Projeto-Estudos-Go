package main

import (
	"fmt"
	"net/http"
)


func ola(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Ola mundo")
}

func cabecalho(w http.ResponseWriter, req *http.Request){
	for nome, cabecalho := reange cabecalho{
		for _, c := range cabecalho{
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main(){
	http.HandleFunc("/ola", ola)
    http.HandleFunc("/cabecalho", cabecalho)

	
    http.ListenAndServe("localhost:8080", nil)
    fmt.Println("Servidor rodando na porta 8080")
    fmt.Scanln()
    fmt.Println("Servidor parado")
    fmt.Scanln()
    fmt.Println("Servidor rodando na porta 8080")
    fmt.Scanln()
    fmt.Println("Servidor parado")
    fmt.Scanln()
    fmt.Println("Servidor rodando na porta 8080")
    fmt.Scanln()
    fmt.Println("Servidor parado")
    fmt.Scanln()
    fmt.Println("Servidor rodando na porta 8080")
}
