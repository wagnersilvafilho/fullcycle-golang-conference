package main

import (
	"fmt"

	"github.com/wagnersilvafilho/aprendagolang/imersao/users"
)

func main() {
	u := users.New("Wagner", 26)
	users.AddYear(&u)
	fmt.Println(u)
}
