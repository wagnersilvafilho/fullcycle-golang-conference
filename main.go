package main

import (
	"fmt"
	"net/http"
	"time"
)

type Person struct {
	name string
	age  int
}

func contador(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}

func helloWorld() {
	print("Hello World!\n")
}

func (p Person) Walk() {
	fmt.Println(p.name + " is Walking")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func listenWebServer() {
	fmt.Println("Servidor escutando na porta :8000")
	http.ListenAndServe(":8000", nil)
}

func main() {
	helloWorld()
	var a int
	a = 10
	fmt.Println(a)
	var pessoa Person
	pessoa.age = 26
	pessoa.name = "Wagner"
	fmt.Println(pessoa)
	pessoa.Walk()
	go contador(3)
	go contador(3)
	contador(3)
	http.HandleFunc("/", home)
	listenWebServer()
}
