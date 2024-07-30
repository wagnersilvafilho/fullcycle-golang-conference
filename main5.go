package main

import (
	"fmt"

	"github.com/wagnersilvafilho/aprendagolang/imersao/users1"
)

func main() {
	u := users1.New("Wagner", 26)
	users1.AddYear(&u)
	fmt.Println(u)
}
