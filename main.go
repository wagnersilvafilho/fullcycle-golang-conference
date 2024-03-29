package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 2
const iterador = 1

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

func testSite(sites []string) {
	for i := 0; i < iterador; i++ {
		for _, site := range sites {
			resp, _ := http.Get(site)
			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso!")
			} else {
				fmt.Println("Site:", site, "retornou status diferente de 200. Status code:", resp.StatusCode)
			}
		}
		fmt.Println(i+1, "ª validação realizada.")
		time.Sleep(delay * time.Second)
	}
}

func checkPageUp() {
	fmt.Println("Enviando requisição de teste...")
	sites := []string{"https://google.com", "https://g1.globo.com", "https://tray.com.br"}
	testSite(sites)
}

func quit() {
	fmt.Println("Bye...")
	os.Exit(0)
}

func returnToMenu(opcao string) {
	fmt.Println("\nDigite uma opção: ")
	fmt.Println("1- Voltar ao menu anterior")
	fmt.Println("0- Sair")
	fmt.Scan(&opcao)
	fmt.Println("A opção escolhida foi", opcao)
	if opcao == "0" {
		quit()
	}
}

func printMenu() {
	fmt.Println("\n1- Iniciar contador")
	fmt.Println("2- Iniciar servidor web")
	fmt.Println("3- Iniciar validação de sites")
	fmt.Println("0- Sair do programa")
	fmt.Println("Digite uma opção: ")
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
	for {
		printMenu()
		fmt.Scan(&opcao)
		fmt.Println("A opção escolhida foi", opcao)
		switch opcao {
		case "1":
			contador(3)
			returnToMenu(opcao)
			break
		case "2":
			listenWebServer()
			break
		case "3":
			checkPageUp()
			returnToMenu(opcao)
			break
		case "0":
			quit()
		default:
			fmt.Print("\nOpção inválida.\n\n")
		}
	}
}
