package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Person struct {
	name string
	age  int
}

func contador(count int) {
	fmt.Println("Executando contador... até", count)
	for i := 1; i <= count; i++ {
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
	http.HandleFunc("/", home)
	fmt.Println("Servidor escutando na porta :8000")
	http.ListenAndServe(":8000", nil)
}

func checkPageUp() {
	fmt.Println("Enviando requisição de teste...")
	site := "https://google.com"
	resp, _ := http.Get(site)
	fmt.Println(resp.Status)
}

func main() {
	helloWorld()
	var opcao string
	a := 10
	fmt.Println(a)
	var pessoa Person
	pessoa.age = 26
	pessoa.name = "Wagner"
	fmt.Println(pessoa)
	pessoa.Walk()
	fmt.Println("Olá ", pessoa.name)
	fmt.Println("1- Iniciar contador")
	fmt.Println("2- Iniciar servidor web")
	fmt.Println("0- Sair do programa")
	for {
		fmt.Println("Digite uma opção: ")
		fmt.Scan(&opcao)
		fmt.Println("A opção escolhida foi", opcao)
		switch opcao {
		case "1":
			contador(3)
			break
		case "2":
			listenWebServer()
			break
		case "3":
			checkPageUp()
			os.Exit(0)
		case "0":
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
